package network

import (
	"fmt"
	"path"
)

const (
	BeaconClientPrysm = "prysm"

	DockerComposeDefaultName = "docker-compose.yml"
	DockerGitLocation        = "https://raw.githubusercontent.com/lukso-network/network-configs/main/%s/docker/docker-compose.%s.yml"
	ConfigGitLocation        = "https://raw.githubusercontent.com/lukso-network/network-configs/main/%s/configs"
)

type ResourceDownloader struct {
	ConfigLocation string
	DockerLocation string
	Chain          Chain
}

func NewResourceDownloader(chain Chain, clientVersion string) ResourceDownloader {
	downloader := ResourceDownloader{}
	downloader.DockerLocation = fmt.Sprintf(DockerGitLocation, chain.String(), clientVersion)
	downloader.ConfigLocation = fmt.Sprintf(ConfigGitLocation, chain.String())
	downloader.Chain = chain
	return downloader
}

func (d ResourceDownloader) DownloadAll(nodeName string) error {
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

func (d ResourceDownloader) downloadDocker() error {
	fmt.Printf("Fetching docker-compose from %s..... ", d.DockerLocation)
	if err := downloadFile(d.DockerLocation, DockerComposeDefaultName); err != nil {
		return err
	}
	fmt.Printf("success\n")
	return nil
}

func (d ResourceDownloader) downloadConfig() error {
	for _, fileName := range ConfigFiles {
		fileUrl, err := getConfigFileUrl(d.Chain, fileName)
		if err != nil {
			return err
		}
		fmt.Printf("Fetching %s from %s..... ", fileName, fileUrl)
		destLocation := path.Join(ConfigDirectory, fileName)
		if err = downloadFile(fileUrl.String(), destLocation); err != nil {
			return err
		}
		fmt.Printf("success\n")
	}
	return nil
}
