package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/creativeprojects/clog"
)

func generateDocs() error {
	prepareTheme()

	versions, err := getVersions(versionsPathPrefix)
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
	err := os.Symlink(filepath.Join("../..", versionsPathPrefix, version, "content"), "./source/docs/content")
	if err != nil {
		return fmt.Errorf("cannot create symlink: %w", err)
	}
	defer os.Remove("./source/docs/content")

	err = os.Symlink(filepath.Join("../..", versionsPathPrefix, version, "jsonschema"), "./source/docs/static/jsonschema")
	if err != nil {
		return fmt.Errorf("cannot create symlink: %w", err)
	}
	defer os.Remove("./source/docs/static/jsonschema")

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
