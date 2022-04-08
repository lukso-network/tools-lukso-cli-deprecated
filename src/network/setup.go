package network

import (
	"context"
	"errors"
	"fmt"
	"github.com/hashicorp/go-getter"
	"github.com/joho/godotenv"
	"github.com/lukso-network/lukso-cli/src"
	"gopkg.in/yaml.v3"
	"os"
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
	configs, err := GetLoadedNodeConfigs()
	if err != nil {
		return err
	}
	dstConfigPath, err := configs.getConfigPath()
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

func GenerateDefaultNodeConfigs(chainId string) error {
	var nodeConfig *NodeConfigs
	switch chainId {
	case src.L16Network:
		nodeConfig = DefaultL16NodeConfigs
	default:
		return errors.New("invalid chainId selected")
	}
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	luksoConfigHomePath := path.Join(userHomeDir, fmt.Sprintf(".lukso_%s", chainId))
	err = os.MkdirAll(luksoConfigHomePath, 0777)
	if err != nil {
		return err
	}
	luksoConfigFile := path.Join(luksoConfigHomePath, "node_configs.yaml")
	yamlData, err := yaml.Marshal(nodeConfig)
	if err != nil {
		return err
	}
	return os.WriteFile(luksoConfigFile, yamlData, os.ModePerm)
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
