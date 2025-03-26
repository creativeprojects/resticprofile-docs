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
	pflag.StringVar(&baseURL, "baseURL", defaultBaseURL, "base URL without version")
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
		clog.Info("please specify any of the commands: [snapshot, cleanup, pageversions, theme, generate, serve or changelog]")
	}
	if err != nil {
		clog.Error(err)
		os.Exit(1)
	}
}
