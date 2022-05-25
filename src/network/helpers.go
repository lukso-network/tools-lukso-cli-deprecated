package network

import (
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
)

func getNetworkDataLocation(chain Chain) (*url.URL, error) {
	u, err := url.Parse(GitUrl)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, ConfigRepoName, ConfigBranchName, chain.String())
	return u, nil
}

func getNetworkConfigUrl(chain Chain) (*url.URL, error) {
	urlWithNetworkName, err := getNetworkDataLocation(chain)
	if err != nil {
		return nil, err
	}
	urlWithNetworkName.Path = path.Join(urlWithNetworkName.Path, ConfigDirectory)
	return urlWithNetworkName, nil
}

func getBootnodeUrl(chain Chain) (*url.URL, error) {
	urlWithNetworkName, err := getNetworkDataLocation(chain)
	if err != nil {
		return nil, err
	}
	urlWithNetworkName.Path = path.Join(urlWithNetworkName.Path, BootnodesDirectory, BootnodeJSONName)
	return urlWithNetworkName, nil
}

func getConfigFileUrl(chain Chain, fileName string) (*url.URL, error) {
	urlWithNetworkName, err := getNetworkConfigUrl(chain)
	if err != nil {
		return nil, err
	}
	urlWithNetworkName.Path = path.Join(urlWithNetworkName.Path, fileName)
	return urlWithNetworkName, nil
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

func GetLoadedNodeConfigs() (*NodeConfigs, error) {
	var nodeConfig NodeConfigs
	err := viper.Unmarshal(&nodeConfig)
	if err != nil {
		return nil, err
	}
	return &nodeConfig, nil
}
