package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/creativeprojects/clog"
	"github.com/go-git/go-git/v5"
)

func openRepo(path string) (*git.Repository, error) {
	_ = os.MkdirAll(path, 0o700)
	repo, err := git.PlainOpen(path)
	if err != nil {
		clog.Info("source empty, cloning repository...")
		repo, err = git.PlainClone(path, false, &git.CloneOptions{URL: "https://github.com/creativeprojects/resticprofile.git"})
		if err != nil {
			return nil, fmt.Errorf("cannot clone git repository: %w", err)
		}
	}

	worktree, err := repo.Worktree()
	if err != nil {
		return nil, fmt.Errorf("cannot load worktree: %w", err)
	}
	status, err := worktree.Status()
	if err != nil {
		return nil, fmt.Errorf("cannot get status: %w", err)
	}
	if !status.IsClean() {
		return nil, errors.New("repository is not clean")
	}
	err = worktree.Pull(&git.PullOptions{RemoteName: "origin"})
	if err != nil && !errors.Is(err, git.NoErrAlreadyUpToDate) {
		return nil, fmt.Errorf("cannot pull from remote: %w", err)
	}
	return repo, nil
}
