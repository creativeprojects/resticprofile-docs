package main

import (
	"flag"
	"os"

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
		err = cleanupDocs(versionsPathPrefix)

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
