package version

import (
	"fmt"
	"runtime"
)

const ReleaseURL = "https://github.com/lukso-network/lukso-cli/releases/download/%s/lukso-cli-%s-%s"

func Install(version string) error {
	downloadURL := fmt.Sprintf(ReleaseURL, version, runtime.GOOS, runtime.GOARCH)
	fmt.Println(downloadURL)
	return nil
}
