package network

import (
	"errors"
	"github.com/spf13/cobra"
	"os"
	"path"
)

func (config *NodeConfigs) getNodeName(proposedNodeName string) string {
	if proposedNodeName != "" {
		return proposedNodeName
	}
	nodeInfo := config.getNode()
	if nodeInfo != nil {
		configNodeName := nodeInfo.getName()
		if configNodeName != "" {
			return configNodeName
		}
	}
	hostName, _ := os.Hostname()
	return hostName
}

func (config *NodeConfigs) getConfigPath() (string, error) {
	if config.Configs == nil || config.getConfigs().getVolume() == "" {
		homedir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		defaultConfigPath := path.Join(homedir, ".lukso_configs")
		return defaultConfigPath, err
	}
	return config.Configs.Volume, nil
}

func (config *NodeConfigs) getKeyStorePath() (string, error) {
	if config.Configs == nil || config.getConfigs().getVolume() == "" {
		homedir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		defaultConfigPath := path.Join(homedir, ".lukso_keystore")
		return defaultConfigPath, err
	}
	return config.Keystore.Volume, nil
}

func (config *NodeConfigs) getExecutionData() *ClientDetails {
	return config.Execution
}

func (config *NodeConfigs) getConsensusData() *ClientDetails {
	return config.Consensus
}

func (config *NodeConfigs) getValidatorData() *ClientDetails {
	return config.Validator
}

func getDataFromContainer(dataContainer *ClientDetails, dataKey string) string {
	if dataContainer == nil {
		return ""
	}
	switch dataKey {
	case keyStatsAddress:
		return dataContainer.getStatAddress()
	case keyVerbosity:
		return dataContainer.getVerbosity()
	case keyEtherBase:
		return dataContainer.getEtherbase()
	case keyDataVolume:
		return dataContainer.getDataVolume()
	case keyNetworkId:
		return dataContainer.getNetworkID()
	case keyBootNode:
		return dataContainer.getBootnode()
	case keyVersion:
		return dataContainer.getVersion()
	default:
		return ""
	}
}

func (config *NodeConfigs) getGethHttpPort() (string, error) {
	gethPorts := config.getPort("geth")
	if gethPorts != nil {
		return gethPorts.getHttpPort(), nil
	}
	return "", errors.New("gethPorts are not available in config file")
}

func (config *NodeConfigs) gethGethPeerPort() (string, error) {
	gethPorts := config.getPort("geth")
	if gethPorts != nil {
		return gethPorts.getPeerPort(), nil
	}
	return "", errors.New("gethPorts are not available in config   file")
}

func getEnvironmentConfig(nodeName string) map[string]string {

	nodeConfig, err := GetLoadedNodeConfigs()
	if err != nil {
		cobra.CompError(err.Error())
		return nil
	}
	newEnvData := make(map[string]string)
	newEnvData["CONFIGS_VOLUME"], err = nodeConfig.getConfigPath()
	if err != nil {
		cobra.CompError(err.Error())
		return nil
	}
	newEnvData["KEYSTORE_VOLUME"], err = nodeConfig.getKeyStorePath()
	if err != nil {
		cobra.CompError(err.Error())
		return nil
	}
	newEnvData["EXTERNAL_IP"], err = getPublicIP()
	if err != nil {
		cobra.CompError(err.Error())
		return nil
	}
	newEnvData["NODE_NAME"] = nodeConfig.getNodeName(nodeName)
	executionContainer := nodeConfig.getExecutionData()
	consensusContainer := nodeConfig.getConsensusData()
	validatorContainer := nodeConfig.getValidatorData()

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

	gethHttpPort, err := nodeConfig.getGethHttpPort()
	if err != nil {
		cobra.CompError(err.Error())
		return nil
	}
	gethPeerPort, err := nodeConfig.gethGethPeerPort()
	if err != nil {
		cobra.CompError(err.Error())
		return nil
	}

	newEnvData["GETH_HTTP_PORT"] = gethHttpPort
	newEnvData["GETH_PEER_PORT"] = gethPeerPort
	return newEnvData
}
