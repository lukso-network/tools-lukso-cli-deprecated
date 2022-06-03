package network

import (
	"errors"
	"github.com/spf13/cobra"
)

func (config *NodeConfigs) getGethHttpPort() (string, error) {
	gethPorts := config.getPort("geth")
	if gethPorts != nil {
		return gethPorts.HttpPort, nil
	}
	return "", errors.New("gethPorts are not available in config file")
}

func (config *NodeConfigs) gethGethPeerPort() (string, error) {
	gethPorts := config.getPort("geth")
	if gethPorts != nil {
		return gethPorts.PeerPort, nil
	}
	return "", errors.New("gethPorts are not available in config   file")
}

func GetEnvironmentConfig() map[string]string {
	nodeConfig := MustGetNodeConfig()
	var err error
	newEnvData := make(map[string]string)
	newEnvData["CONFIGS_VOLUME"] = nodeConfig.Configs.Volume
	newEnvData["KEYSTORE_VOLUME"] = nodeConfig.Keystore.Volume
	newEnvData["EXTERNAL_IP"], err = getPublicIP()
	if err != nil {
		cobra.CompError(err.Error())
		return nil
	}
	newEnvData["NODE_NAME"] = nodeConfig.Node.Name
	e := nodeConfig.Execution
	c := nodeConfig.Consensus
	v := nodeConfig.Validator

	newEnvData["EXECUTION_DATA_VOLUME"] = e.DataVolume
	newEnvData["CONSENSUS_DATA_VOLUME"] = c.DataVolume
	newEnvData["VALIDATOR_DATA_VOLUME"] = v.DataVolume

	newEnvData["ETH_STATS_ADDRESS"] = e.StatsAddress
	newEnvData["ETH_2_STATS_ADDRESS"] = c.StatsAddress

	newEnvData["GETH_VERBOSITY"] = e.Verbosity
	newEnvData["PRYSM_VERBOSITY"] = c.Verbosity

	newEnvData["PRYSM_BEACON_VERSION"] = c.Version
	newEnvData["GETH_VERSION"] = e.Version
	newEnvData["GETH_ETHERBASE"] = e.Etherbase

	newEnvData["GETH_NETWORK_ID"] = e.NetworkId
	newEnvData["PRYSM_BOOTSTRAP_NODE"] = c.Bootnode
	newEnvData["GETH_BOOTSTRAP_NODE"] = e.Bootnode

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
