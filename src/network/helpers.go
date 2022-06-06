package network

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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
