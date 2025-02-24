package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/creativeprojects/clog"
	"github.com/go-git/go-git/v5"
)

const (
	themeSubmoduleName = "docs/themes/hugo-theme-relearn"
)

func generateDocs() error {
	source := "./source"
	repo, err := openRepo(source)
	if err != nil {
		return fmt.Errorf("cannot load repository: %w", err)
	}
	worktree, err := repo.Worktree()
	if err != nil {
		return fmt.Errorf("cannot load worktree: %w", err)
	}
	theme, err := worktree.Submodule(themeSubmoduleName)
	if err != nil {
		return fmt.Errorf("cannot find submodule: %w", err)
	}
	status, err := theme.Status()
	if err != nil {
		return fmt.Errorf("cannot load submodule status: %w", err)
	}
	clog.Infof("updating theme submodule to %s", status.Expected.String())
	err = theme.Update(&git.SubmoduleUpdateOptions{
		Init:  true,
		Depth: 1,
	})
	if err != nil {
		return fmt.Errorf("cannot update git submodule: %w", err)
	}

	versions, err := getVersions("./")
	if err != nil {
		return fmt.Errorf("cannot load versions: %w", err)
	}

	err = os.Rename("./source/docs/content", "./source/docs/content___")
	if err != nil {
		return fmt.Errorf("cannot rename content: %w", err)
	}
	for _, version := range versions {
		clog.Debugf("generating version %s", version)
		err = generateDocVersion(version)
		if err != nil {
			return err
		}
	}
	err = os.Rename("./source/docs/content___", "./source/docs/content")
	if err != nil {
		return fmt.Errorf("cannot rename content: %w", err)
	}
	return nil
}

func generateDocVersion(version string) error {
	err := os.Symlink(fmt.Sprintf("../../%s", version), "./source/docs/content")
	if err != nil {
		return fmt.Errorf("cannot create symlink: %w", err)
	}
	defer os.Remove("./source/docs/content")

	cmd := exec.Command("hugo", "build", "--minify", "--cleanDestinationDir", "--destination", fmt.Sprintf("../public/%s", version))
	cmd.Dir = "./source/docs"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
