package main

import (
	"flag"
	"os"
	"strings"

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
		version := flag.Arg(1)
		if version == "" || !strings.HasPrefix(version, "v") {
			clog.Error("please specify which version (of file path) to cleanup: mkdocs cleanup v0.18.0")
			os.Exit(1)
		}
		err = cleanupDocs(version)

	case "generate":
		err = generateDocs()

	default:
		clog.Info("please specify any of the commands: [snapshot, cleanup or generate]")
	}
	if err != nil {
		clog.Error(err)
		os.Exit(1)
	}
}
