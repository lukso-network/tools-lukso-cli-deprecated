package network

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"

	"github.com/lukso-network/lukso-cli/src"
)

func getRepoUrl() (*url.URL, error) {
	u, err := url.Parse(src.GitUrl)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, src.ConfigRepoName)
	return u, nil
}

func getBranchUrl() (*url.URL, error) {
	repoUrl, err := getRepoUrl()
	if err != nil {
		return nil, err
	}
	repoUrl.Path = path.Join(repoUrl.Path, src.ConfigBranchName)
	return repoUrl, nil
}

func getUrlWithNetworkName() (*url.URL, error) {
	branchUrl, err := getBranchUrl()
	if err != nil {
		return nil, err
	}
	branchUrl.Path = path.Join(branchUrl.Path, viper.GetString(src.ViperKeyNetworkName))
	return branchUrl, nil
}

func getNetworkConfigUrl() (*url.URL, error) {
	urlWithNetworkName, err := getUrlWithNetworkName()
	if err != nil {
		return nil, err
	}
	urlWithNetworkName.Path = path.Join(urlWithNetworkName.Path, "dev", src.NetworkVersion)
	return urlWithNetworkName, nil
}

func getNetSetupFileUrl(fileName string) (*url.URL, error) {
	urlWithNetworkName, err := getUrlWithNetworkName()
	if err != nil {
		return nil, err
	}
	urlWithNetworkName.Path = path.Join(urlWithNetworkName.Path, "network_setup_kit/"+fileName)
	return urlWithNetworkName, nil
}

func getConfigFileUrl(fileName string) (*url.URL, error) {
	urlWithNetworkName, err := getNetworkConfigUrl()
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

func getConfigPath() (string, error) {
	envVariables, err := godotenv.Read(".env")
	if err != nil {
		return "", err
	}

	return envVariables["CONFIGS_VOLUME"], nil
}

func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
