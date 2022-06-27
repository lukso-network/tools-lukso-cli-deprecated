package version

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-getter"
	"github.com/lukso-network/lukso-cli/src/utils"
	"log"
	"os"
	"runtime"
)

const (
	ReleaseURL = "https://github.com/lukso-network/lukso-cli/releases/download/%s/lukso-cli-%s-%s"
	UserPath   = "/usr/local/bin"
)

func Install(version string) error {
	downloadURL := fmt.Sprintf(ReleaseURL, version, runtime.GOOS, runtime.GOARCH)
	err := os.MkdirAll(UserPath, os.ModePerm)
	if err != nil {
		return err
	}
	client := &getter.Client{
		Ctx:  context.Background(),
		Src:  downloadURL,
		Dst:  UserPath + "/lukso",
		Dir:  true,
		Mode: getter.ClientModeFile,
	}
	if err = client.Get(); err != nil {
		return err
	}
	err = os.Chmod(UserPath+"/lukso", 0755)
	if err != nil {
		log.Fatal(err)
	}
	utils.Coloredln(fmt.Sprintf("%s installed!", version))
	return nil
}
