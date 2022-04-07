package network

import (
	"context"
	"github.com/hashicorp/go-getter"
	"github.com/joho/godotenv"
	"github.com/lukso-network/lukso-cli/src"
	"path"
)

func downloadFile(src, dest string) error {
	client := &getter.Client{
		Ctx:  context.Background(),
		Src:  src,
		Dst:  dest,
		Dir:  true,
		Mode: getter.ClientModeFile,
	}
	if err := client.Get(); err != nil {
		return err
	}
	return nil
}

func downloadNetworkSetupFiles() error {
	for _, fileName := range src.NetworkSetupFiles {
		fileUrl, err := getNetSetupFileUrl(fileName)
		if err != nil {
			return err
		}
		if err = downloadFile(fileUrl.String(), fileName); err != nil {
			return err
		}
	}
	return nil
}

func downloadConfigFiles() error {
	dstConfigPath, err := getConfigPath()
	if err != nil {
		return err
	}
	for _, fileName := range src.ConfigFiles {
		fileUrl, err := getConfigFileUrl(fileName)
		if err != nil {
			return err
		}
		destLocation := path.Join(dstConfigPath, fileName)
		if err = downloadFile(fileUrl.String(), destLocation); err != nil {
			return err
		}
	}
	return nil
}

func generateEnvFile(hostName string) error {

	return godotenv.Write(getEnvironmentConfig(hostName), ".env")
}

func SetupNetwork(nodeName string) error {
	err := downloadNetworkSetupFiles()
	if err != nil {
		return err
	}

	err = generateEnvFile(nodeName)
	if err != nil {
		return err
	}

	err = downloadConfigFiles()
	if err != nil {
		return err
	}
	return nil
}
