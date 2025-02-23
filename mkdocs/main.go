package main

import (
	"flag"
	"fmt"
	"io/fs"
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
	default:
		err = createSnapshots()
	}
	if err != nil {
		clog.Error(err)
		os.Exit(1)
	}
}

func cleanupDocs(source string) error {
	return filepath.WalkDir(source, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".md" {
			fmt.Println(path)
		}
		return nil
	})
}
