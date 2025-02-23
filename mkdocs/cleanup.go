package main

import (
	"bytes"
	"errors"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/creativeprojects/clog"
	"github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v3"
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
	var headerTOML, headerYAML bool
	var bufferTOML, bufferYAML = &bytes.Buffer{}, &bytes.Buffer{}
	var lineNum int
	lines := bytes.Split(content, []byte{'\n'})
	for _, line := range lines {
		lineNum++
		line = bytes.TrimSpace(line)
		lineStr := string(line)
		if lineStr == "+++" {
			clog.Tracef("TOML marker on line %d", lineNum)
			headerTOML = !headerTOML
			if !headerTOML {
				break
			}
			continue
		}
		if lineStr == "---" {
			clog.Tracef("YAML marker on line %d", lineNum)
			headerYAML = !headerYAML
			if !headerYAML {
				break
			}
			continue
		}
		if headerTOML {
			bufferTOML.Write(line)
			bufferTOML.WriteByte('\n')
		} else if headerYAML {
			bufferYAML.Write(line)
			bufferYAML.WriteByte('\n')
		}
	}
	rewrite := false
	header := make(map[string]any)
	if bufferTOML.Len() > 0 {
		rewrite = true
		decoder := toml.NewDecoder(bufferTOML)
		err = decoder.Decode(&header)
	} else if bufferYAML.Len() > 0 {
		decoder := yaml.NewDecoder(bufferYAML)
		err = decoder.Decode(&header)
	}
	if err != nil {
		return err
	}
	if _, found := header["date"]; found {
		rewrite = true
		delete(header, "date")
	}
	if rewrite {
		clog.Debugf("rewrite needed: %+v\n", header)
	}
	return nil
}
