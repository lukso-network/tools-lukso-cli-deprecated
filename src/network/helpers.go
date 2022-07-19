package network

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-getter"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func GetIPAndHostName(nodeName string) (*NodeDetails, error) {
	ip, err := getPublicIP()
	if err != nil {
		return nil, err
	}

	if nodeName == "" {
		hostName, err := os.Hostname()
		if err != nil {
			return nil, err
		}
		nodeName = hostName
	}

	return &NodeDetails{
		IP:   ip,
		Name: nodeName,
	}, nil
}

func getPublicIP() (string, error) {
	url := "https://api.ipify.org?format=text"
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", ip), nil
}

func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

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
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("couldn't download file %d", resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}

func removeContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}
