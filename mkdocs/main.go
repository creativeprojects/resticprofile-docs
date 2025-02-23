package main

import (
	"bytes"
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/creativeprojects/clog"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func main() {
	clog.SetDefaultLogger(clog.NewFilteredConsoleLogger(clog.LevelDebug))
	source := "./source"
	_ = os.MkdirAll(source, 0o700)
	repo, err := git.PlainOpen(source)
	if err != nil {
		clog.Info("source empty, cloning repository...")
		repo, err = git.PlainClone(source, false, &git.CloneOptions{URL: "https://github.com/creativeprojects/resticprofile.git"})
		if err != nil {
			clog.Errorf("cannot clone git repository: %s", err)
			os.Exit(1)
		}
	}

	worktree, err := repo.Worktree()
	if err != nil {
		clog.Errorf("cannot load worktree: %s", err)
		os.Exit(1)
	}
	status, err := worktree.Status()
	if err != nil {
		clog.Errorf("cannot get status: %s", err)
		os.Exit(1)
	}
	if !status.IsClean() {
		clog.Error("repository is not clean, cannot generate documentation")
		os.Exit(1)
	}
	err = worktree.Pull(&git.PullOptions{RemoteName: "origin"})
	if err != nil && !errors.Is(err, git.NoErrAlreadyUpToDate) {
		clog.Errorf("cannot pull from remote: %s", err)
		os.Exit(1)
	}
	head, err := repo.Head()
	if err != nil {
		clog.Errorf("cannot retrieve HEAD: %s", err)
		os.Exit(1)
	}
	clog.Debugf("HEAD is at %s on %s", head.Hash(), head.Name().Short())
	tagrefs, err := repo.Tags()
	if err != nil {
		clog.Errorf("cannot load tags: %s", err)
		os.Exit(1)
	}
	err = tagrefs.ForEach(func(t *plumbing.Reference) error {
		clog.Debugf("--> tag %s", t.Name().Short())
		err := worktree.Checkout(&git.CheckoutOptions{Branch: t.Name()})
		if err != nil {
			clog.Warningf("cannot checkout %q: %s", t.Name(), err)
		}
		found, reference := detectDocumentation(t.Name().Short(), worktree)
		if found && reference == "none" {
			from := filepath.Join(source, "docs/content")
			to := filepath.Join("./", t.Name().Short())
			destInfo, err := os.Stat(to)
			// don't recreate the snapshot if it already exists
			if err == fs.ErrNotExist && destInfo == nil {
				_ = os.MkdirAll(to, 0o777)
				clog.Infof("copying %q to %q", from, to)
				err = copyDocs(from, to)
				if err != nil {
					clog.Warningf("cannot copy files: %s", err)
				}
			}
		}

		err = worktree.Reset(&git.ResetOptions{Mode: git.HardReset, Commit: head.Hash()})
		if err != nil {
			clog.Warningf("cannot reset %q: %s", t.Name(), err)
		}
		return nil
	})
	if err != nil {
		clog.Errorf("error iterating over tags: %s", err)
		os.Exit(1)
	}
}

func detectDocumentation(tagRef string, worktree *git.Worktree) (bool, string) {
	reference := "none"
	finfo, err := worktree.Filesystem.Stat("/docs")
	if err != nil || finfo == nil || !finfo.IsDir() {
		return false, ""
	}
	finfo, err = worktree.Filesystem.Stat("/docs/content/reference")
	if err == nil && finfo != nil && finfo.IsDir() {
		reference = "static"
	} else {
		file, err := worktree.Filesystem.Open("/Makefile")
		if err != nil {
			clog.Errorf("cannot open Makefile: %s", err)
			return true, reference
		}
		defer file.Close()
		makefile, err := io.ReadAll(file)
		if err != nil {
			clog.Errorf("cannot read Makefile: %s", err)
			return true, reference
		}
		if bytes.Contains(makefile, []byte("reference")) {
			reference = "generated"
		}
	}
	clog.Infof("documentation found in tag %s with reference: %s", tagRef, reference)
	return true, reference
}

func copyDocs(source, dest string) error {
	return filepath.WalkDir(source, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		newPath := filepath.Join(dest, strings.TrimPrefix(path, source))
		if d.IsDir() {
			return os.MkdirAll(newPath, 0777)
		}

		if !d.Type().IsRegular() {
			return &fs.PathError{Op: "CopyFS", Path: path, Err: fs.ErrInvalid}
		}

		r, err := os.Open(path)
		if err != nil {
			return err
		}
		defer r.Close()
		info, err := r.Stat()
		if err != nil {
			return err
		}
		w, err := os.OpenFile(newPath, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0666|info.Mode()&0777)
		if err != nil {
			return err
		}

		if _, err := io.Copy(w, r); err != nil {
			w.Close()
			return &fs.PathError{Op: "Copy", Path: newPath, Err: err}
		}
		return w.Close()
	})
}
