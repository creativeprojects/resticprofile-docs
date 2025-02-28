package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/creativeprojects/clog"
)

const (
	pageVersionsOpeningShortcode = "{{< pageversions "
	pageVersionsClosingShortcode = " >}}"
)

func createPageOtherVersions() error {
	pagesPerVersion := make(map[string][]string)
	versions, err := getVersions(versionsPathPrefix)
	if err != nil {
		return err
	}
	// build a list of pages per version
	for _, version := range versions {
		clog.Infof("processing version %s", version)
		pages, err := parseVersion(filepath.Join(versionsPathPrefix, version, contentDirectory))
		if err != nil {
			return fmt.Errorf("cannot parse version %s: %w", version, err)
		}
		clog.Infof("version %s: found %d pages", version, len(pages))
		pagesPerVersion[version] = pages
	}
	// add tag on each page for each version
	for version, pages := range pagesPerVersion {
		for _, page := range pages {
			path := findPagePath(version, page)
			clog.Debugf("adding tag to %s", path)
			err := addOtherVersionsTag(path, otherVersions(pagesPerVersion, page, version))
			if err != nil {
				return fmt.Errorf("cannot add tag to %s: %w", path, err)
			}
		}
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

func parseVersion(root string) ([]string, error) {
	pages := make([]string, 0)
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		relpath := strings.TrimPrefix(path, root)
		if d.IsDir() {
			return nil
		}
		if filepath.Ext(relpath) != fileExt {
			return nil
		}
		if filepath.Base(relpath) == "index.md" || filepath.Base(relpath) == "_index.md" {
			pages = append(pages, filepath.Dir(relpath))
			return nil
		}
		pages = append(pages, strings.TrimSuffix(relpath, fileExt))
		return nil
	})
	if err != nil {
		return nil, err
	}
	return pages, nil
}

func findPagePath(version, page string) string {
	path := ""
	if page != "/" {
		// try .md file directly
		path = filepath.Join(versionsPathPrefix, version, contentDirectory, page+".md")
		finfo, err := os.Stat(path)
		if err == nil && finfo != nil && !finfo.IsDir() {
			return path
		}
	}
	// try _index.md file
	path = filepath.Join(versionsPathPrefix, version, contentDirectory, page, "_index.md")
	finfo, err := os.Stat(path)
	if err == nil && finfo != nil && !finfo.IsDir() {
		return path
	}
	// try index.md file
	path = filepath.Join(versionsPathPrefix, version, contentDirectory, page, "index.md")
	finfo, err = os.Stat(path)
	if err == nil && finfo != nil && !finfo.IsDir() {
		return path
	}
	panic(fmt.Sprintf("cannot find page %q in version %s", page, version))
}

func addOtherVersionsTag(path string, versions []string) error {
	input, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	output, err := os.Create(path)
	if err != nil {
		return err
	}
	defer output.Close()

	lines := strings.Split(string(input), "\n")
	found := false
	for i, line := range lines {
		if strings.Contains(line, pageVersionsOpeningShortcode) {
			found = true
			line = buildTag(versions)
		}
		// don't add newline at the end of the file
		if line != "" || i < len(lines)-1 {
			output.WriteString(line)
			output.WriteString("\n")
		}
	}
	if !found {
		output.WriteString("\n")
		output.WriteString(buildTag(versions))
		output.WriteString("\n")
	}
	return nil
}

func buildTag(versions []string) string {
	if len(versions) == 0 {
		return ""
	}
	return fmt.Sprintf(`%s"%s"%s`, pageVersionsOpeningShortcode, strings.Join(versions, `" "`), pageVersionsClosingShortcode)
}

func otherVersions(pagesPerVersion map[string][]string, page, currentVersion string) []string {
	other := make([]string, 0, len(pagesPerVersion[currentVersion])-1)
	for version, pages := range pagesPerVersion {
		if currentVersion == version {
			continue
		}
		if slices.Contains(pages, page) {
			other = append(other, version)
		}
	}
	// keys coming from map are not sorted
	slices.Sort(other)
	return other
}
