package main

import (
	"fmt"

	"github.com/creativeprojects/clog"
	"github.com/go-git/go-git/v5"
)

const (
	themeSubmoduleName = "docs/themes/hugo-theme-relearn"
)

func generateDocs(path string) error {
	source := "./source"
	repo, err := openRepo(source)
	if err != nil {
		return err
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
	return nil
}
