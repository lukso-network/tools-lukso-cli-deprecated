package network

import (
	"fmt"
	"os"
	"path"
)

const (
	BeaconClientPrysm = "prysm"

	DockerComposeDefaultName = "docker-compose.yml"
	DockerGitLocation        = "https://raw.githubusercontent.com/lukso-network/network-configs/main/devnets/%s/docker/docker-compose.%s.yml"
	ConfigGitLocation        = "https://raw.githubusercontent.com/lukso-network/network-configs/main/devnets/%s/configs"
	DevDockerGitLocation     = "https://raw.githubusercontent.com/lukso-network/network-configs/main/devnets/%s/docker/docker-compose.%s.yml"
	DevConfigGitLocation     = "https://raw.githubusercontent.com/lukso-network/network-configs/main/devnets/%s/configs"
)

type ResourceDownloader struct {
	ConfigLocation string
	DockerLocation string
}

func NewResourceDownloader(chain Chain, clientVersion string) ResourceDownloader {
	downloader := ResourceDownloader{}
	downloader.DockerLocation = fmt.Sprintf(DockerGitLocation, chain.String(), clientVersion)
	downloader.ConfigLocation = fmt.Sprintf(ConfigGitLocation, chain.String())
	return downloader
}

func NewDevResourceDownloader(devLocation, clientVersion string) ResourceDownloader {
	downloader := ResourceDownloader{}
	downloader.DockerLocation = fmt.Sprintf(DevDockerGitLocation, devLocation, clientVersion)
	downloader.ConfigLocation = fmt.Sprintf(DevConfigGitLocation, devLocation)
	return downloader
}

func (d ResourceDownloader) DownloadAll() error {
	err := d.downloadDocker()
	if err != nil {
		return err
	}

	err = d.downloadConfig()
	if err != nil {
		return err
	}

	err = GenerateEnvFile()
	if err != nil {
		return err
	}

	return nil
}

func (d ResourceDownloader) downloadDocker() error {
	fmt.Printf("Fetching docker-compose from %s..... ", d.DockerLocation)
	os.Remove(DockerComposeDefaultName)
	if err := downloadFile(d.DockerLocation, DockerComposeDefaultName); err != nil {
		return err
	}
	fmt.Printf("success\n")
	return nil
}

func (d ResourceDownloader) downloadConfig() error {
	for _, fileName := range []string{"config.yaml", "genesis.json", "genesis.ssz"} {
		fileUrl := fmt.Sprintf("%s/%s", d.ConfigLocation, fileName)
		fmt.Printf("Fetching %s from %s..... ", fileName, fileUrl)
		destLocation := path.Join(ConfigDirectory, fileName)
		if err := downloadFile(fileUrl, destLocation); err != nil {
			fmt.Printf("unsuccessful\n")
			// we can ignore the gensis.ssz as a chain can be setup in a way where it is not needed
			if fileName != "genesis.ssz" {
				return err
			} else {
				continue
			}
		}
		fmt.Printf("success\n")
	}
	return nil
}
