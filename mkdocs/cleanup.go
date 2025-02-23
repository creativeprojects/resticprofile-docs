package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/creativeprojects/clog"
)

func cleanupDocs(version string) error {
	if version == "" {
		return errors.New("invalid version")
	}
	return filepath.WalkDir(version, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if filepath.Ext(path) != ".md" {
			return nil
		}
		clog.Debugf("cleaning up %q", path)
		err = cleanupMD(path)
		if err != nil {
			clog.Warning("cleaning up %q: %s", path, err)
		}
		return nil
	})
}

func cleanupMD(path string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	var toml, yaml bool
	var bufferTOML, bufferYAML = &bytes.Buffer{}, &bytes.Buffer{}
	lines := bytes.Split(content, []byte{'\n'})
	for lineNum, line := range lines {
		line = bytes.TrimSpace(line)
		lineStr := string(line)
		if lineStr == "+++" {
			clog.Tracef("TOML marker on line %d", lineNum+1)
			toml = !toml
			if !toml {
				break
			}
			continue
		}
		if lineStr == "---" {
			clog.Tracef("YAML marker on line %d", lineNum+1)
			yaml = !yaml
			if !yaml {
				break
			}
			continue
		}
		if toml {
			bufferTOML.Write(line)
			bufferTOML.WriteByte('\n')
		} else if yaml {
			bufferYAML.Write(line)
			bufferYAML.WriteByte('\n')
		}
	}
	if bufferTOML.Len() > 0 {
		fmt.Println(bufferTOML.String())
	} else if bufferYAML.Len() > 0 {
		fmt.Println(bufferYAML.String())
	}
	return nil
}
