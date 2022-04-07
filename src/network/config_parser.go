package network

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path"
)

func getNodeInfo() map[string]interface{} {
	return viper.GetStringMap(keyNode)
}

func getNodeName(proposedNodeName string) string {
	if proposedNodeName != "" {
		return proposedNodeName
	}
	nodeInfo := getNodeInfo()
	if nodeInfo[keyName] != nil {
		configNodeName := nodeInfo[keyName].(string)
		if configNodeName != "" {
			return configNodeName
		}
	}
	hostName, _ := os.Hostname()
	return hostName
}

func getVolume(key, defaultDir string) (string, error) {
	configs := viper.GetStringMapString(key)
	if configs == nil || configs[keyVolume] == "" {
		homedir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		defaultConfigPath := path.Join(homedir, defaultDir)
		return defaultConfigPath, err
	}
	return configs[keyVolume], nil
}

func GetConfigPath() (string, error) {
	return getVolume(keyConfigs, ".lukso_config")
}

func GetKeyStorePath() (string, error) {
	return getVolume(keyKeystore, ".lukso_keystore")
}

func getExecutionData() map[string]string {
	return viper.GetStringMapString(keyExecution)
}

func getConsensusData() map[string]string {
	return viper.GetStringMapString(keyConsensus)
}

func getValidatorData() map[string]string {
	return viper.GetStringMapString(keyValidator)
}

func getDataFromContainer(dataContainer map[string]string, dataKey string) string {
	if dataContainer == nil {
		return ""
	}
	return dataContainer[dataKey]
}

func getPorts() map[string]interface{} {
	return viper.GetStringMap(keyPorts)
}

func getGethPorts() map[string]interface{} {
	allPorts := getPorts()
	if allPorts == nil || allPorts[keyGeth] == nil {
		return nil
	}
	return allPorts[keyGeth].(map[string]interface{})
}

func getGethHttpPort() (int, error) {
	gethPorts := getGethPorts()
	if gethPorts == nil || gethPorts[keyhttpPort] == nil {
		return 0, errors.New("gethPorts are not available in config file")
	}
	return gethPorts[keyhttpPort].(int), nil
}

func gethGethPeerPort() (int, error) {
	gethPorts := getGethPorts()
	if gethPorts == nil || gethPorts[keypeerPort] == nil {
		return 0, errors.New("gethPorts are not available in config file")
	}
	return gethPorts[keypeerPort].(int), nil
}

func getEnvironmentConfig(nodeName string) map[string]string {
	newEnvData := make(map[string]string)
	var err error
	newEnvData["CONFIGS_VOLUME"], err = GetConfigPath()
	if err != nil {
		cobra.CompError(err.Error())
		return nil
	}
	newEnvData["KEYSTORE_VOLUME"], err = GetKeyStorePath()
	if err != nil {
		cobra.CompError(err.Error())
		return nil
	}
	newEnvData["EXTERNAL_IP"], err = getPublicIP()
	if err != nil {
		cobra.CompError(err.Error())
		return nil
	}
	newEnvData["NODE_NAME"] = getNodeName(nodeName)
	executionContainer := getExecutionData()
	consensusContainer := getConsensusData()
	validatorContainer := getValidatorData()

	newEnvData["EXECUTION_DATA_VOLUME"] = getDataFromContainer(executionContainer, keyDataVolume)
	newEnvData["CONSENSUS_DATA_VOLUME"] = getDataFromContainer(consensusContainer, keyDataVolume)
	newEnvData["VALIDATOR_DATA_VOLUME"] = getDataFromContainer(validatorContainer, keyDataVolume)

	newEnvData["ETH_STATS_ADDRESS"] = getDataFromContainer(executionContainer, keyStatsAddress)
	newEnvData["ETH_2_STATS_ADDRESS"] = getDataFromContainer(consensusContainer, keyStatsAddress)

	newEnvData["GETH_VERBOSITY"] = getDataFromContainer(executionContainer, keyVerbosity)
	newEnvData["PRYSM_VERBOSITY"] = getDataFromContainer(consensusContainer, keyVerbosity)

	newEnvData["PRYSM_BEACON_VERSION"] = getDataFromContainer(consensusContainer, keyVersion)
	newEnvData["GETH_ETHERBASE"] = getDataFromContainer(executionContainer, keyEtherBase)

	newEnvData["GETH_NETWORK_ID"] = getDataFromContainer(executionContainer, keyNetworkId)
	newEnvData["PRYSM_BOOTSTRAP_NODE"] = getDataFromContainer(consensusContainer, keyBootNode)
	newEnvData["GETH_BOOTSTRAP_NODE"] = getDataFromContainer(executionContainer, keyBootNode)

	gethHttpPort, err := getGethHttpPort()
	if err != nil {
		cobra.CompError(err.Error())
		return nil
	}
	gethPeerPort, err := gethGethPeerPort()
	if err != nil {
		cobra.CompError(err.Error())
		return nil
	}

	newEnvData["GETH_HTTP_PORT"] = fmt.Sprintf("%d", gethHttpPort)
	newEnvData["GETH_PEER_PORT"] = fmt.Sprintf("%d", gethPeerPort)
	return newEnvData
}
