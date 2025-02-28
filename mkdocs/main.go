package main

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/creativeprojects/clog"
)

func main() {
	var verbose bool
	flag.BoolVar(&verbose, "v", false, "display more information")
	flag.Parse()

	level := clog.LevelInfo
	if verbose {
		level = clog.LevelDebug
	}
	clog.SetDefaultLogger(clog.NewFilteredConsoleLogger(level))

	var err error
	switch flag.Arg(0) {
	case "snapshot":
		err = createSnapshots()

	case "cleanup":
		path := versionsPathPrefix
		version := flag.Arg(1)
		if version != "" {
			path = filepath.Join(versionsPathPrefix, version)
		}
		err = cleanupDocs(path)

	case "otherversions":
		err = createOtherVersions()

	case "theme":
		err = prepareTheme()

	case "generate":
		err = generateDocs()

	case "serve":
		version := flag.Arg(1)
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
