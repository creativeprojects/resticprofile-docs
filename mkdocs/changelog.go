package main

import (
	"context"
	"os"
	"regexp"
	"strings"

	"github.com/creativeprojects/clog"
	"github.com/google/go-github/v69/github"
	"gopkg.in/yaml.v3"
)

var (
	changelogFrontmatter = map[string]interface{}{
		"archetype": "chapter",
		"pre":       "<b>8. </b>",
		"title":     "Release Notes",
		"weight":    8,
	}
)

func createReleaseNotes() error {
	ctx := context.Background()
	client := github.NewClient(nil)
	if token := os.Getenv("GITHUB_TOKEN"); token != "" {
		client = client.WithAuthToken(token)
	}

	changelog, err := os.Create("changelog.md")
	if err != nil {
		return err
	}
	defer changelog.Close()

	changelog.WriteString("---\n")
	header, err := yaml.Marshal(changelogFrontmatter)
	if err != nil {
		return err
	}
	changelog.Write(header)
	changelog.WriteString("---\n\n")

	page := 1
	for {
		clog.Debugf("fetching page %d", page)
		releases, resp, err := client.Repositories.ListReleases(ctx, "creativeprojects", "resticprofile", &github.ListOptions{PerPage: 10, Page: page})
		if err != nil {
			return err
		}

		for _, release := range releases {
			body := release.GetBody()
			if body == "" {
				continue
			}
			body = cleanBody(body)

			changelog.WriteString("# ")
			changelog.WriteString(release.GetName())
			changelog.WriteString(" (")
			changelog.WriteString(release.GetPublishedAt().Format("2006-01-02"))
			changelog.WriteString(")\n\n")
			changelog.WriteString(body)
			changelog.WriteString("\n")
		}

		page = resp.NextPage
		if page == 0 {
			// last page
			break
		}
	}
	return nil
}

func cleanBody(input string) string {
	changelogPattern := regexp.MustCompile(`^(\* ){0,1}([0-9a-f]+) `)
	output := &strings.Builder{}
	lines := strings.Split(input, "\n")
	changelog := false
	for _, line := range lines {
		line = strings.TrimSuffix(line, "\r")
		if strings.HasPrefix(line, "## Changelog") {
			changelog = true
		}
		if changelog {
			line = changelogPattern.ReplaceAllString(line, "* [$2](https://github.com/creativeprojects/resticprofile/commit/$2) ")
		} else {
			// remove links to old documentation
			line = regexp.MustCompile(`\[(.*?)\]\(https://github\.com/creativeprojects/resticprofile/tree.*?\)`).ReplaceAllString(line, "$1")
		}
		_, _ = output.WriteString(line)
		_ = output.WriteByte('\n')
	}
	return output.String()
}
