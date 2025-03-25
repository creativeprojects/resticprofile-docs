package main

import (
	"os"
	"path/filepath"

	"github.com/creativeprojects/clog"
	"github.com/spf13/pflag"
)

func main() {
	var verbose bool
	var baseURL string
	pflag.BoolVarP(&verbose, "verbose", "v", false, "display more information")
	pflag.StringVar(&baseURL, "baseURL", "", "base URL (without version)")
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
		err = generateDocs()

	case "serve":
		version := pflag.Arg(1)
		if version == "" {
			clog.Errorf("missing version: mkdocs serve v0.21.0")
			os.Exit(1)
		}
		err = serveDocVersion(version)

	case "changelog":
		err = createReleaseNotes()

	default:
		clog.Info("please specify any of the commands: [snapshot, cleanup or generate]")
	}
	if err != nil {
		clog.Error(err)
		os.Exit(1)
	}
}
