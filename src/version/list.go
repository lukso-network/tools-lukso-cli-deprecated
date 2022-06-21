package version

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
)

func List() error {
	client := github.NewClient(nil)
	releases, rsp, err := client.Repositories.ListReleases(context.Background(), "lukso-network", "lukso-cli", nil)
	if err != nil {
		return err
	}

	fmt.Printf("\n%+v\n", releases)
	fmt.Printf("\n%+v\n", rsp)
	return nil
}

func ListRemote() error {
	return nil
}

func parse() error {
	return nil
}
