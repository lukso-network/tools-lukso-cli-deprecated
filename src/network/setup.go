package network

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-getter"
	"github.com/joho/godotenv"
	"io/ioutil"
	"net/http"
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

func downloadFileOverHttp(url string) ([]byte, error) {
	fmt.Println("Fetching bootnodes from", url)
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func GenerateEnvFile(hostName string) error {
	return godotenv.Write(GetEnvironmentConfig(hostName), ".env")
}

func GenerateDefaultNodeConfigs(chain Chain) error {
	var nodeConfig = GetDefaultNodeConfig(chain)
	return nodeConfig.WriteOrUpdateNodeConfig()
}

func SetupNetwork(chain Chain, nodeName string) error {
	fmt.Printf("Setting up network for chain %s.\n", chain.String())
	clientVersion := BeaconClientPrysm

	if !IsChainSupported(chain) {
		return fmt.Errorf("the network %s does not exist or is not supported\n", chain.String())
	}

	return NewResourceDownloader(chain, clientVersion).DownloadAll(nodeName)
}
