package network

import (
	"fmt"
	"github.com/lukso-network/lukso-cli/src"
	"path"
)

const (
	BeaconClientPrysm = "prysm"

	DockerComposeDefaultName = "docker-compose.yml"
	DockerGitLocation        = "https://raw.githubusercontent.com/lukso-network/network-configs/main/%s/docker/docker-compose.%s.yml"
	ConfigGitLocation        = "https://raw.githubusercontent.com/lukso-network/network-configs/main/%s/configs"
	ConfigDirectory          = "configs"
)

type NetworkResourceConfig struct {
	ConfigLocation string
	DockerLocation string
}

func NewResourceConfig(networkVersion string, clientVersion string) (NetworkResourceConfig, error) {
	c := NetworkResourceConfig{}
	if networkVersion != ChainL16Beta && networkVersion != ChainL16 {
		return c, fmt.Errorf("unknown network version %s", networkVersion)
	}
	c.DockerLocation = fmt.Sprintf(DockerGitLocation, networkVersion, clientVersion)
	c.ConfigLocation = fmt.Sprintf(ConfigGitLocation, networkVersion)
	return c, nil
}

type NetworkResourceDownloader struct {
	NetworkResourceConfig NetworkResourceConfig
}

func NewNetworkResourceDownloader(config NetworkResourceConfig) NetworkResourceDownloader {
	return NetworkResourceDownloader{
		config,
	}
}

func (d NetworkResourceDownloader) DownloadAll(nodeName string) error {
	err := d.downloadDocker()
	if err != nil {
		return err
	}

	err = d.downloadConfig()
	if err != nil {
		return err
	}

	err = GenerateEnvFile(nodeName)
	if err != nil {
		return err
	}

	return nil
}

func (d NetworkResourceDownloader) downloadDocker() error {
	if err := downloadFile(d.NetworkResourceConfig.DockerLocation, DockerComposeDefaultName); err != nil {
		return err
	}
	return nil
}

func (d NetworkResourceDownloader) downloadConfig() error {
	for _, fileName := range src.ConfigFiles {
		fileUrl, err := getConfigFileUrl(fileName)
		if err != nil {
			return err
		}
		destLocation := path.Join(ConfigDirectory, fileName)
		if err = downloadFile(fileUrl.String(), destLocation); err != nil {
			return err
		}
	}
	return nil
}
