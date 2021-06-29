package main

import (
	"context"
	"github.com/google/go-github/v36/github"
	"github.com/hashicorp/go-version"
	"io/ioutil"
	"math"
	"os"
	"sort"
	"strings"
)

func main() {
	client := github.NewClient(nil)
	releases, _, err := client.Repositories.ListReleases(context.Background(), os.Args[1], os.Args[2], &github.ListOptions{PerPage: math.MaxInt32})
	if err != nil {
		panic(err)
	}

	changelog := "# Changelog\n\n"

	sort.Slice(releases, func(i, j int) bool {
		versionI, _ := version.NewVersion(releases[i].GetTagName())
		versionJ, _ := version.NewVersion(releases[j].GetTagName())
		return versionI.GreaterThan(versionJ)
	})

	for _, release := range releases {
		changelog += strings.ReplaceAll(release.GetBody(), "# ", "## ") + "\n\n"
	}

	err = ioutil.WriteFile("./CHANGELOG.md", []byte(changelog), 0777)
	if err != nil {
		panic(err)
	}
}
