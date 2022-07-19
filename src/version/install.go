package version

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-getter"
	"github.com/lukso-network/lukso-cli/src/utils"
	"io"
	"os"
	"runtime"
	"time"
)

const (
	ReleaseURL = "https://github.com/lukso-network/lukso-cli/releases/download/%s/lukso-cli-%s-%s"
	UserPath   = "/usr/local/bin"
	TempPath   = "/tmp/"
)

func Install(version string) error {
	downloadURL := fmt.Sprintf(ReleaseURL, version, runtime.GOOS, runtime.GOARCH)
	err := os.MkdirAll(UserPath, os.ModePerm)
	if err != nil {
		return err
	}
	err = os.MkdirAll(TempPath, os.ModePerm)
	if err != nil {
		return err
	}
	tmpFile := TempPath + "/lukso_" + time.Now().String()
	client := &getter.Client{
		Ctx:  context.Background(),
		Src:  downloadURL,
		Dst:  tmpFile,
		Dir:  false,
		Mode: getter.ClientModeFile,
	}
	if err = client.Get(); err != nil {
		return err
	}
	err = os.Remove(UserPath + "/lukso")
	if err != nil {
		return err
	}
	err = copyFile(tmpFile, UserPath+"/lukso")
	if err != nil {
		return err
	}
	err = os.Remove(tmpFile)
	if err != nil {
		return err
	}
	err = os.Chmod(UserPath+"/lukso", 0755)
	if err != nil {
		return err
	}
	utils.Coloredln(fmt.Sprintf("%s installed!", version))
	return nil
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}
