package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/creativeprojects/clog"
)

func createOtherVersions() error {
	versions, err := getVersions(versionsPathPrefix)
	if err != nil {
		return err
	}
	for _, version := range versions {
		clog.Infof("processing version %s", version)
	}
	return nil
}

func getVersions(from string) ([]string, error) {
	entries, err := os.ReadDir(from)
	if err != nil {
		return nil, fmt.Errorf("cannot read source: %w", err)
	}
	versions := make([]string, 0)
	for _, entry := range entries {
		if entry.IsDir() && strings.HasPrefix(entry.Name(), "v") {
			versions = append(versions, entry.Name())
		}
	}
	return versions, nil
}

func parseVersion(path, version string) error {
	return nil
}
