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

func (config *NodeConfigs) GetKeyStorePath() (string, error) {
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

func GetExecutionDataVolume() (string, error) {
	config := MustGetNodeConfig()
	executionContainer := config.getExecutionData()
	return getDataFromContainer(executionContainer, keyDataVolume), nil
}

func GetConsensusDataVolume() (string, error) {
	config := MustGetNodeConfig()
	dataContainer := config.getConsensusData()
	return getDataFromContainer(dataContainer, keyDataVolume), nil
}

func GetValidatorDataVolume() (string, error) {
	config := MustGetNodeConfig()
	dataContainer := config.getValidatorData()
	return getDataFromContainer(dataContainer, keyDataVolume), nil
}

func GetEnvironmentConfig() map[string]string {
	nodeConfig := MustGetNodeConfig()
	var err error
	newEnvData := make(map[string]string)
	newEnvData["CONFIGS_VOLUME"], err = nodeConfig.getConfigPath()
	if err != nil {
		cobra.CompError(err.Error())
		return nil
	}
	newEnvData["KEYSTORE_VOLUME"], err = nodeConfig.GetKeyStorePath()
	if err != nil {
		cobra.CompError(err.Error())
		return nil
	}
	newEnvData["EXTERNAL_IP"], err = getPublicIP()
	if err != nil {
		cobra.CompError(err.Error())
		return nil
	}
	newEnvData["NODE_NAME"] = nodeConfig.Node.Name
	executionContainer := nodeConfig.getExecutionData()
	consensusContainer := nodeConfig.getConsensusData()
	validatorContainer := nodeConfig.getValidatorData()

	newEnvData["EXECUTION_DATA_VOLUME"] = executionContainer.DataVolume
	newEnvData["CONSENSUS_DATA_VOLUME"] = consensusContainer.DataVolume
	newEnvData["VALIDATOR_DATA_VOLUME"] = validatorContainer.DataVolume

	newEnvData["ETH_STATS_ADDRESS"] = executionContainer.StatsAddress
	newEnvData["ETH_2_STATS_ADDRESS"] = consensusContainer.StatsAddress

	newEnvData["GETH_VERBOSITY"] = executionContainer.Verbosity
	newEnvData["PRYSM_VERBOSITY"] = consensusContainer.Verbosity

	newEnvData["PRYSM_BEACON_VERSION"] = consensusContainer.Version
	newEnvData["GETH_VERSION"] = executionContainer.Version
	newEnvData["GETH_ETHERBASE"] = executionContainer.Etherbase

	newEnvData["GETH_NETWORK_ID"] = executionContainer.NetworkId
	newEnvData["PRYSM_BOOTSTRAP_NODE"] = consensusContainer.Bootnode
	newEnvData["GETH_BOOTSTRAP_NODE"] = executionContainer.Bootnode

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
