package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"

	"github.com/creativeprojects/clog"
	"github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v3"
)

const (
	tagTOML = "+++"
	tagYAML = "---"
	fileExt = ".md"
)

var (
	removeHeaders     = []string{"date", "tags"}
	domainReplacement = []string{
		"https://creativeprojects.github.io/resticprofile",
		"https://dev.resticprofile.pages.dev/%s",
	}
	simpleReplacements = [][]string{
		{"tabs groupId=", "tabs groupid="},
		{"tab name=", "tab title="},
		{" vebose ", " verbose "},
	}
	regexpReplacements = [][]string{
		{`{{<\s*ref\s+"([^"]*)"\s*>}}`, `{{% ref "$1" %}}`},
		{`{{%(\s*)attachments\s+`, `{{%${1}resources `},
	}
)

func cleanupDocs(root, version string) error {
	if root == "" {
		return errors.New("please specify a path of file(s) to cleanup")
	}
	// single file?
	finfo, err := os.Stat(root)
	if err != nil {
		return err
	}
	if !finfo.IsDir() {
		if filepath.Ext(root) == fileExt {
			return cleanupMD(root, version)
		}
		return fmt.Errorf("not a directory: %q", root)
	}
	return filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if filepath.Ext(path) != fileExt {
			return nil
		}
		clog.Debugf("cleaning up %q", path)
		err = cleanupMD(path, version)
		if err != nil {
			clog.Warning("cleaning up %q: %s", path, err)
		}
		return nil
	})
}

func cleanupMD(path, version string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	var headerTOML, headerYAML, hasHeader bool
	var bufferTOML, bufferYAML = &bytes.Buffer{}, &bytes.Buffer{}
	var lineNum int
	lines := bytes.Split(content, []byte{'\n'})
	for _, line := range lines {
		lineNum++
		line = bytes.TrimSpace(line)
		lineStr := string(line)
		if lineStr == tagTOML {
			clog.Tracef("TOML marker on line %d", lineNum)
			hasHeader = true
			headerTOML = !headerTOML
			if !headerTOML {
				break
			}
			continue
		}
		if lineStr == tagYAML {
			clog.Tracef("YAML marker on line %d", lineNum)
			hasHeader = true
			headerYAML = !headerYAML
			if !headerYAML {
				break
			}
			continue
		}
		if headerTOML {
			_, _ = bufferTOML.Write(line)
			_ = bufferTOML.WriteByte('\n')
		} else if headerYAML {
			_, _ = bufferYAML.Write(line)
			_ = bufferYAML.WriteByte('\n')
		}
	}
	headerChanged := false
	header := make(map[string]any)
	if bufferTOML.Len() > 0 {
		headerChanged = true
		decoder := toml.NewDecoder(bufferTOML)
		err = decoder.Decode(&header)
	} else if bufferYAML.Len() > 0 {
		decoder := yaml.NewDecoder(bufferYAML)
		err = decoder.Decode(&header)
	}
	if err != nil {
		return err
	}
	for _, removeHeader := range removeHeaders {
		if _, found := header[removeHeader]; found {
			headerChanged = true
			delete(header, removeHeader)
		}
	}

	if !hasHeader {
		// the remaining is the whole content
		lineNum = 0
	}

	remainingLines, contentChanged := cleanContent(lines[lineNum:], version)
	if headerChanged || contentChanged {
		clog.Debugf("rewrite needed: %+v\n", header)
		return rewriteMD(path, header, remainingLines)
	}
	return nil
}

func rewriteMD(filename string, header map[string]any, lines [][]byte) error {
	finfo, err := os.Stat(filename)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(filename, os.O_TRUNC|os.O_WRONLY, finfo.Mode())
	if err != nil {
		return err
	}
	defer file.Close()

	if len(header) > 0 {
		err = writeLine(file, []byte(tagYAML))
		if err != nil {
			return err
		}
		encoder := yaml.NewEncoder(file)
		err = encoder.Encode(header)
		if err != nil {
			return err
		}
		err = writeLine(file, []byte(tagYAML))
		if err != nil {
			return err
		}
	}

	for _, line := range lines {
		err = writeLine(file, line)
		if err != nil {
			return err
		}
	}
	return nil
}

func writeLine(w io.Writer, content []byte) error {
	_, err := w.Write(content)
	if err != nil {
		return err
	}
	_, err = w.Write([]byte{'\n'})
	return err
}

func cleanContent(lines [][]byte, version string) ([][]byte, bool) {
	changed := false
	output := make([][]byte, len(lines))
	for i, line := range lines {

		if bytes.Contains(line, []byte(domainReplacement[0])) {
			changed = true
			line = bytes.ReplaceAll(
				line,
				[]byte(domainReplacement[0]),
				[]byte(fmt.Sprintf(domainReplacement[1], version)),
			)
		}

		for _, replacement := range simpleReplacements {
			if bytes.Contains(line, []byte(replacement[0])) {
				changed = true
				line = bytes.ReplaceAll(
					line,
					[]byte(replacement[0]),
					[]byte(replacement[1]),
				)
			}
		}
		for _, replacement := range regexpReplacements {
			pattern := regexp.MustCompile(replacement[0])
			if pattern.Find(line) != nil {
				changed = true
				line = pattern.ReplaceAll(
					line,
					[]byte(replacement[1]))
			}
		}
		output[i] = line
	}
	return output, changed
}
