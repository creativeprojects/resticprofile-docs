package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"text/template"

	"github.com/creativeprojects/clog"
)

func generateDocs(baseURL string) error {
	err := prepareTheme()
	if err != nil {
		return fmt.Errorf("cannot load theme: %w", err)
	}

	allVersions, err := getVersions(versionsPathPrefix)
	if err != nil {
		return fmt.Errorf("cannot load versions: %w", err)
	}

	_ = os.RemoveAll("./public")

	for _, version := range allVersions {
		clog.Debugf("generating version %s", version)
		err = generateDocVersion(baseURL, version, allVersions)
		if err != nil {
			return err
		}
	}

	return nil
}

func generateDocVersion(baseURL, version string, allVersions []string) error {
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

	err = generateHugoConfig(baseURL, version, allVersions)
	if err != nil {
		return err
	}

	latestVersion := allVersions[len(allVersions)-1]
	publishDir := publishDirectory
	if latestVersion != version {
		publishDir = filepath.Join(publishDir, version)
	}

	cmd := exec.Command(
		"hugo",
		"build",
		"--minify",
		"--destination", publishDir,
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

func serveDocVersion(baseURL, version string) error {
	// catch CRTL-C but do nothing with the context as the hugo command will end
	_, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer cancel()

	allVersions, err := getVersions(versionsPathPrefix)
	if err != nil {
		return fmt.Errorf("cannot load versions: %w", err)
	}

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

	err = generateHugoConfig(baseURL, version, allVersions)
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

func generateHugoConfig(baseURL, currentVersion string, allVersions []string) error {
	if !strings.HasSuffix(baseURL, "/") {
		baseURL += "/"
	}
	latestVersion := allVersions[len(allVersions)-1]
	templ, err := template.ParseFiles(hugoConfigTemplate)
	if err != nil {
		return err
	}
	file, err := os.Create(hugoConfigFile)
	if err != nil {
		return err
	}
	defer file.Close()

	versionURL := baseURL
	if currentVersion != latestVersion {
		versionURL = fmt.Sprintf("%s%s/", baseURL, currentVersion)
	}
	data := TemplateContext{
		Current: TemplateVersion{
			BaseURL:  versionURL,
			Version:  currentVersion,
			Title:    currentVersion,
			IsLatest: currentVersion == latestVersion,
		},
		Versions: make([]TemplateVersion, len(allVersions)),
	}
	for i, otherVersion := range allVersions {
		versionURL := baseURL
		if otherVersion != latestVersion {
			versionURL = fmt.Sprintf("%s%s/", baseURL, otherVersion)
		}
		data.Versions[i] = TemplateVersion{
			BaseURL:  versionURL,
			Version:  otherVersion,
			Title:    otherVersion,
			IsLatest: otherVersion == latestVersion,
		}
	}
	err = templ.Execute(file, data)
	if err != nil {
		return err
	}
	return nil
}
