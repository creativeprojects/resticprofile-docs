package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/creativeprojects/clog"
)

func generateDocs() error {
	err := prepareTheme()
	if err != nil {
		return fmt.Errorf("cannot load theme: %w", err)
	}

	versions, err := getVersions(versionsPathPrefix)
	if err != nil {
		return fmt.Errorf("cannot load versions: %w", err)
	}

	err = linkContent()
	if err != nil {
		return err
	}

	for _, version := range versions {
		clog.Debugf("generating version %s", version)
		err = generateDocVersion(version)
		if err != nil {
			return err
		}
	}

	return unlinkContent()
}

func generateDocVersion(version string) error {
	err := os.Symlink(filepath.Join("../..", versionsPathPrefix, version, "content"), "./source/docs/content")
	if err != nil {
		return fmt.Errorf("cannot create symlink: %w", err)
	}
	defer os.Remove("./source/docs/content")

	// this cannot be a symbolic link
	jsonschemaSource := filepath.Join(versionsPathPrefix, version, "jsonschema")
	if isDir(jsonschemaSource) {
		err = copyFiles(jsonschemaSource, "./source/docs/static/jsonschema")
		if err != nil {
			return fmt.Errorf("cannot copy jsonschema files: %w", err)
		}
		defer os.RemoveAll("./source/docs/static/jsonschema")
	}

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

func serveDocVersion(version string) error {
	// catch CRTL-C but do nothing with the context as the hugo command will end
	_, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer cancel()

	err := linkContent()
	if err != nil {
		return err
	}
	defer func() {
		err := unlinkContent()
		if err != nil {
			clog.Error(err)
		}
	}()

	err = os.Symlink(filepath.Join("../..", versionsPathPrefix, version, "content"), "./source/docs/content")
	if err != nil {
		return fmt.Errorf("cannot create symlink: %w", err)
	}
	defer os.Remove("./source/docs/content")

	// this cannot be a symbolic link
	jsonschemaSource := filepath.Join(versionsPathPrefix, version, "jsonschema")
	if isDir(jsonschemaSource) {
		err = copyFiles(jsonschemaSource, "./source/docs/static/jsonschema")
		if err != nil {
			return fmt.Errorf("cannot copy jsonschema files: %w", err)
		}
		defer os.RemoveAll("./source/docs/static/jsonschema")
	}

	cmd := exec.Command(
		"hugo",
		"serve",
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

func linkContent() error {
	if isDir("./source/docs/content") {
		err := os.Rename("./source/docs/content", "./source/docs/content___")
		if err != nil {
			return fmt.Errorf("cannot rename content: %w", err)
		}
	}
	return nil
}

func unlinkContent() error {
	err := os.Rename("./source/docs/content___", "./source/docs/content")
	if err != nil {
		return fmt.Errorf("cannot rename content: %w", err)
	}
	return nil
}
