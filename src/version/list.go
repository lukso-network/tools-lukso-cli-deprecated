package version

import (
	"context"
	"github.com/google/go-github/github"
	"github.com/lukso-network/lukso-cli/src/utils"
	"runtime"
	"strings"
)

func List(currentVersion string) error {

	releases, err := getReleases()
	if err != nil {
		return err
	}
	for _, release := range releases {
		releaseTag := *release.TagName
		if runtime.GOOS == "windows" {
			releaseTag = strings.TrimRight(releaseTag, "\r\n")
		} else {
			releaseTag = strings.TrimRight(releaseTag, "\n")
		}
		if strings.Compare(releaseTag, currentVersion) == 0 {
			utils.SelectedColorPrintln("Currently installed", releaseTag, utils.ConsoleColorGreen, utils.ConsoleColorBlue)
		} else {
			utils.Coloredln(releaseTag)
		}
	}
	return nil
}

func GetLatestVersion() (string, error) {
	releases, err := getReleases()
	if err != nil {
		return "", err
	}
	latestRelease := releases[0]
	return *latestRelease.TagName, nil
}

func getReleases() ([]*github.RepositoryRelease, error) {
	client := github.NewClient(nil)
	releases, _, err := client.Repositories.ListReleases(context.Background(), "lukso-network", "lukso-cli", nil)
	if err != nil {
		return nil, err
	}
	return releases, nil
}
