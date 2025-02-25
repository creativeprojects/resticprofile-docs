package main

import (
	"fmt"
	"os"
	"os/exec"
)

const (
	themeSubmoduleName = "docs/themes/hugo-theme-relearn"
)

func generateDocs() error {
	// sourceRepo, err := openSourceRepo(sourceRepositoryPath)
	// if err != nil {
	// 	return fmt.Errorf("cannot load repository: %w", err)
	// }
	// worktree, err := sourceRepo.Worktree()
	// if err != nil {
	// 	return fmt.Errorf("cannot load worktree: %w", err)
	// }
	// stop using shitty submodules
	// theme, err := worktree.Submodule(themeSubmoduleName)
	// if err != nil {
	// 	return fmt.Errorf("cannot find submodule: %w", err)
	// }
	// status, err := theme.Status()
	// if err != nil {
	// 	return fmt.Errorf("cannot load submodule status: %w", err)
	// }
	// themeCommit := status.Expected.String()
	// clog.Infof("cloning or updating theme to %s", themeCommit)
	themeRepo, err := openThemeRepo(themeRepositoryPath, themeVersionTag)
	if err != nil {
		return fmt.Errorf("cannot clone or update theme repository: %w", err)
	}
	themeRepo.Head()

	// versions, err := getVersions(versionsPathPrefix)
	// if err != nil {
	// 	return fmt.Errorf("cannot load versions: %w", err)
	// }

	// err = os.Rename("./source/docs/content", "./source/docs/content___")
	// if err != nil {
	// 	return fmt.Errorf("cannot rename content: %w", err)
	// }
	// for _, version := range versions {
	// 	clog.Debugf("generating version %s", version)
	// 	err = generateDocVersion(version)
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	// err = os.Rename("./source/docs/content___", "./source/docs/content")
	// if err != nil {
	// 	return fmt.Errorf("cannot rename content: %w", err)
	// }
	return nil
}

func generateDocVersion(version string) error {
	err := os.Symlink(fmt.Sprintf("../../%s", version), "./source/docs/content")
	if err != nil {
		return fmt.Errorf("cannot create symlink: %w", err)
	}
	defer os.Remove("./source/docs/content")

	cmd := exec.Command(
		"hugo",
		"build",
		"--minify",
		"--cleanDestinationDir",
		"--destination", fmt.Sprintf("../public/%s", version),
		"--baseURL", fmt.Sprintf("https://dev.resticprofile.pages.dev/%s", version),
		"--themesDir", "../../themes",
	)
	cmd.Dir = "./source/docs"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
