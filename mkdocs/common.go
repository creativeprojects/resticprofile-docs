package main

import (
	"io/fs"
	"os"
)

// isDir returns true only if the path is a real directory.
// It returns false if the path is a symbolic link to a directory
func isDir(path string) bool {
	finfo, err := os.Lstat(path)
	if err != nil || finfo == nil {
		return false
	}
	return finfo.IsDir()
}

func isSymlink(path string) bool {
	finfo, err := os.Lstat(path)
	if err != nil || finfo == nil {
		return false
	}
	return finfo.Mode().Type()&fs.ModeSymlink != 0
}
