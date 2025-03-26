package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/creativeprojects/clog"
	"github.com/spf13/pflag"
)

func main() {
	var verbose bool
	var baseURL string
	pflag.BoolVarP(&verbose, "verbose", "v", false, "display debugging information")
	pflag.StringVar(&baseURL, "baseURL", defaultBaseURL, "base URL without version")

	pflag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\nCommands:\n")
		fmt.Fprintf(os.Stderr, "\tsnapshots\tcreates version snapshots from source git tags, also calls 'cleanup'\n")
		fmt.Fprintf(os.Stderr, "\tcleanup\t\tcleans all the .md files in a version: cleanup v0.22.0\n")
		fmt.Fprintf(os.Stderr, "\tpageversions\tgenerates the 'other versions' at the bottom of each page\n")
		fmt.Fprintf(os.Stderr, "\ttheme\t\tdownloads the hugo theme at the specified version (%s)\n", themeVersionTag)
		fmt.Fprintf(os.Stderr, "\tgenerate\tbuilds the documentation website into the 'public' directory\n")
		fmt.Fprintf(os.Stderr, "\tserve\t\tserves local version of the website (from 'public'). If a version is specified it will serve directly from hugo\n")
		fmt.Fprintf(os.Stderr, "\tchangelog\tcreates the changelog file from release notes on GitHub\n")
		fmt.Fprintf(os.Stderr, "\nGlobal flags:\n")
		pflag.PrintDefaults()
	}
	pflag.Parse()

	level := clog.LevelInfo
	if verbose {
		level = clog.LevelDebug
	}
	clog.SetDefaultLogger(clog.NewFilteredConsoleLogger(level))

	var err error
	switch pflag.Arg(0) {
	case "snapshot":
		err = createSnapshots()

	case "cleanup":
		path := versionsPathPrefix
		version := pflag.Arg(1)
		if version != "" {
			path = filepath.Join(versionsPathPrefix, version)
		}
		err = cleanupDocs(path)

	case "pageversions":
		err = createPageOtherVersions()

	case "theme":
		err = prepareTheme()

	case "generate":
		err = generateDocs(baseURL)

	case "serve":
		version := pflag.Arg(1)
		if version == "" {
			err = serveDirectory("./public")
		} else {
			err = serveDocVersion(baseURL, version)
		}

	case "changelog":
		err = createReleaseNotes()

	default:
		pflag.Usage()
	}
	if err != nil {
		clog.Error(err)
		os.Exit(1)
	}
}
