package network

import (
	"context"
	"errors"
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
	return godotenv.Write(getEnvironmentConfig(hostName), ".env")
}

func GenerateDefaultNodeConfigs(chainId string) error {
	var nodeConfig *NodeConfigs
	switch chainId {
	case src.L16Network:
		nodeConfig = DefaultL16BetaNodeConfigs
	default:
		return errors.New("invalid chainId selected")
	}
	return nodeConfig.WriteOrUpdateNodeConfig()
}

func SetupNetwork(nodeName string) error {
	// TODO When a second network arrives needs to modifable
	networkVersion := ChainL16Beta
	clientVersion := BeaconClientPrysm

	config, err := NewResourceConfig(networkVersion, clientVersion)
	if err != nil {
		return err
	}

	return NewNetworkResourceDownloader(config).DownloadAll(nodeName)
}
