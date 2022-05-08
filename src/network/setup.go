package network

import (
	"context"
	"errors"
	"fmt"
	"os/exec"
	"runtime"

	"github.com/hashicorp/go-getter"
	"github.com/joho/godotenv"
	"github.com/lukso-network/lukso-cli/src"
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

func GenerateEnvFile(hostName string) error {

	err := godotenv.Write(getEnvironmentConfig(hostName), ".env")
	if err != nil {
		return err
	}
	return removeQuoteFromFile(".env")
}

func removeQuoteFromFile(fileName string) error {
	cmd := exec.Command("sed", "-i", "s/\\\"//g", fileName)
	if runtime.GOOS == "darwin" {
		cmd = exec.Command("sed", "-i", "", "s/\\\"//g", fileName)
	}
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error code %s. %s", err, string(output))
	}
	return nil
}

func GenerateDefaultNodeConfigs(chainId string) error {
	var nodeConfig *NodeConfigs
	switch chainId {
	case src.L16Network:
		nodeConfig = DefaultL16NodeConfigs
	default:
		return errors.New("invalid chainId selected")
	}
	return nodeConfig.WriteOrUpdateNodeConfig()
}

func SetupNetwork(nodeName string) error {
	// TODO When a second network arrives needs to modifable
	networkVersion := L16Beta
	clientVersion := BeaconClientPrysm

	config, err := NewResourceConfig(networkVersion, clientVersion)
	if err != nil {
		return err
	}

	return NewNetworkResourceDownloader(config).DownloadAll(nodeName)
}
