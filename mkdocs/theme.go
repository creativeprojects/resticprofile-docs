package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/creativeprojects/clog"
	"github.com/google/go-github/v69/github"
)

func prepareTheme() error {
	expected, err := isExpectedVersion(themeVersionTag)
	if err == nil && expected {
		clog.Infof("theme directory is already at version %s", themeVersionTag)
		return nil
	}

	err = os.MkdirAll(themeRepositoryPath, 0o755)
	if err != nil {
		return err
	}

	ctx := context.Background()
	client := github.NewClient(nil)
	release, _, err := client.Repositories.GetReleaseByTag(ctx, themeRepositoryOwner, themeRepositoryName, themeVersionTag)
	if err != nil {
		return err
	}
	tarball := release.GetTarballURL()
	if tarball == "" {
		return fmt.Errorf("tarball not found in release %s", release.GetName())
	}

	clog.Debugf("downloading %q", tarball)
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, tarball, http.NoBody)
	if err != nil {
		return err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	stripFirstDir := ""

	gz, err := gzip.NewReader(response.Body)
	if err != nil {
		return err
	}
	tr := tar.NewReader(gz)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			return err
		}

		switch hdr.Typeflag {
		case tar.TypeXGlobalHeader:
			// we do nothing with this type

		case tar.TypeDir:
			if stripFirstDir == "" {
				clog.Debugf("stripping first directory: %s", hdr.Name)
				stripFirstDir = hdr.Name
				break
			}
			name := strings.TrimPrefix(hdr.Name, stripFirstDir)
			name = filepath.Join(themeRepositoryPath, name)
			err = os.MkdirAll(name, hdr.FileInfo().Mode().Perm())
			if err != nil {
				return fmt.Errorf("cannot create directory %q: %w", name, err)
			}

		case tar.TypeReg:
			name := strings.TrimPrefix(hdr.Name, stripFirstDir)
			name = filepath.Join(themeRepositoryPath, name)

			output, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, hdr.FileInfo().Mode().Perm())
			if err != nil {
				return fmt.Errorf("cannot open file for writing %q: %w", name, err)
			}
			if _, err := io.Copy(output, tr); err != nil {
				_ = output.Close()
				return fmt.Errorf("cannot extract file %q: %w", name, err)
			}
			_ = output.Close()

		default:
			clog.Warningf("file type not implemented: type '%s', size %d, name %q", string(hdr.Typeflag), hdr.Size, hdr.Name)

		}
	}
	return nil
}

func isExpectedVersion(expected string) (bool, error) {
	changelog, err := os.ReadFile(filepath.Join(themeRepositoryPath, "CHANGELOG.md"))
	if err != nil {
		return false, err
	}
	lines := bytes.Split(changelog, []byte{'\n'})
	for _, line := range lines {
		line = bytes.TrimSpace(line)
		if bytes.HasPrefix(line, []byte("## ")) {
			return bytes.Contains(line, []byte(expected)), nil
		}
	}
	return false, nil
}
