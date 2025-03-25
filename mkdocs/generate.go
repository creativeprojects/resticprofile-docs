package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"syscall"
	"text/template"

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

	for _, version := range versions {
		clog.Debugf("generating version %s", version)
		err = generateDocVersion(version)
		if err != nil {
			return err
		}
	}

	return nil
}

func generateDocVersion(version string) error {
	unlink, err := linkContent(version)
	if err != nil {
		return err
	}
	defer unlink()

	unlinkJsonSchema, err := linkJsonSchema(version)
	if err != nil {
		return err
	}
	defer unlinkJsonSchema()

	cmd := exec.Command(
		"hugo",
		"build",
		"--minify",
		"--cleanDestinationDir",
		"--destination", fmt.Sprintf("../public/%s", version),
		"--baseURL", fmt.Sprintf("https://dev.resticprofile.pages.dev/%s", version),
	)
	cmd.Dir = docsRootPath
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

	unlinkContent, err := linkContent(version)
	if err != nil {
		return err
	}
	defer unlinkContent()

	unlinkJsonSchema, err := linkJsonSchema(version)
	if err != nil {
		return err
	}
	defer unlinkJsonSchema()

	err = generateHugoConfig(version, versions)
	if err != nil {
		return err
	}

	cmd := exec.Command(
		"hugo",
		"serve",
	)
	cmd.Dir = docsRootPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func linkContent(version string) (func() error, error) {
	err := os.Symlink(filepath.Join("..", versionsPathPrefix, version, "content"), filepath.Join(docsRootPath, contentDirectory))
	if err != nil {
		return nil, fmt.Errorf("cannot create symlink: %w", err)
	}
	return func() error {
		return os.Remove(filepath.Join(docsRootPath, contentDirectory))
	}, nil
}

func linkJsonSchema(version string) (func() error, error) {
	// this cannot be a symbolic link
	jsonschemaSource := filepath.Join(versionsPathPrefix, version, "jsonschema")
	if isDir(jsonschemaSource) {
		err := copyFiles(jsonschemaSource, filepath.Join(docsRootPath, "static/jsonschema"))
		if err != nil {
			return nil, fmt.Errorf("cannot copy jsonschema files: %w", err)
		}
		return func() error {
			return os.RemoveAll(filepath.Join(docsRootPath, "static/jsonschema"))
		}, nil
	}
	return func() error {
		return nil
	}, nil
}

func generateHugoConfig(currentVersion string, otherVersions []string) error {
	templ, err := template.ParseFiles(hugoConfigTemplate)
	if err != nil {
		return err
	}
	file, err := os.Create(hugoConfigFile)
	if err != nil {
		return err
	}
	defer file.Close()

	data := TemplateContext{
		Current: TemplateVersion{
			Version: currentVersion,
		},
		Versions: make([]TemplateVersion, len(versions)),
	}
	for i, otherVersion := range versions {
		data.Versions[i] = TemplateVersion{
			Version: otherVersion,
		}
	}
	err = templ.Execute(file, data)
	if err != nil {
		return err
	}
	return nil
}
