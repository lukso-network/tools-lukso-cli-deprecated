package version

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
)

func List() error {
	client := github.NewClient(nil)
	releases, _, err := client.Repositories.ListReleases(context.Background(), "lukso-network", "lukso-cli", nil)
	if err != nil {
		return err
	}

	for _, release := range releases {
		fmt.Println(*release.TagName)
	}
	return nil
}
