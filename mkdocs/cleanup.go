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
	"strings"

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
	removeHeaders = []string{"date", "tags"}
	replacements  = []Replacement{
		{Simple, "tabs groupId=", "tabs groupid="},
		{Simple, "tab name=", "tab title="},
		{Simple, " vebose ", " verbose "},
		{Regexp, `{{<\s*ref\s+"([^"]*)"\s*>}}`, `{{% ref "$1" %}}`},
		{Regexp, `{{%(\s*)attachments\s+`, `{{%${1}resources `},
		{Simple, "% ref \"", "% relref \""},
		{Regexp, `\(https:\/\/creativeprojects\.github\.io\/resticprofile([^\)]*)\)`, `({{% relref "$1" %}})`},
	}
)

func cleanupDocs(root string) error {
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
			return cleanupMD(root)
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
	var headerTOML, headerYAML, hasHeader, fixEndOfYAML bool
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
		} else if strings.HasPrefix(lineStr, tagYAML) && hasHeader {
			// this is a little template bug sticking text right behind the "---" marker
			fixEndOfYAML = true
			break
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
	} else if fixEndOfYAML {
		lineNum--
		lines[lineNum] = bytes.TrimPrefix(lines[lineNum], []byte(tagYAML))
	}

	remainingLines, contentChanged := cleanContent(lines[lineNum:], filepath.Base(path) == "_index.md")
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

	for i, line := range lines {
		// don't write a newline at the end of the file
		if len(line) > 0 || i < len(lines)-1 {
			err = writeLine(file, line)
			if err != nil {
				return err
			}
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

func cleanContent(lines [][]byte, removeFirstTitle bool) ([][]byte, bool) {
	changed := false
	otherLine := false // set to true when we're passed the first title line
	output := make([][]byte, len(lines))
	for i, line := range lines {
		if len(bytes.TrimSpace(line)) > 0 {
			if !bytes.HasPrefix(line, []byte("# ")) {
				otherLine = true
			} else if !otherLine && removeFirstTitle {
				// this is the first title to remove
				changed = true
				continue
			}
			for _, replacement := range replacements {
				switch replacement.Type {
				case Simple:
					if bytes.Contains(line, []byte(replacement.From)) {
						changed = true
						line = bytes.ReplaceAll(
							line,
							[]byte(replacement.From),
							[]byte(replacement.To),
						)
					}

				case Regexp:
					pattern := regexp.MustCompile(replacement.From)
					if pattern.Find(line) != nil {
						changed = true
						line = pattern.ReplaceAll(
							line,
							[]byte(replacement.To))
					}
				}
			}
		}
		output[i] = line
	}
	return output, changed
}
