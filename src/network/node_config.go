package network

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"

	"gopkg.in/yaml.v3"
)

func GetDefaultNodeConfigByOptionParam(chain string) *NodeConfigs {
	return GetDefaultNodeConfig(GetChainByString(chain))
}

func GetDefaultNodeConfig(chain Chain) *NodeConfigs {
	switch chain {
	case L16Beta:
		return DefaultL16BetaNodeConfigs
	case Local:
		return DefaultLocalNodeConfigs
	default:
		return DefaultL16BetaNodeConfigs
	}
}

type DataVolume struct {
	Volume string `yaml:",omitempty"`
}

type NodeDetails struct {
	IP   string `yaml:",omitempty"`
	Name string `yaml:",omitempty"`
}

type ClientDetails struct {
	StatsAddress string `yaml:",omitempty"`
	Verbosity    string `yaml:",omitempty"`
	Etherbase    string `yaml:",omitempty"`
	DataVolume   string `yaml:",omitempty"`
	NetworkId    string `yaml:",omitempty"`
	Bootnode     string `yaml:",omitempty"`
	Version      string `yaml:",omitempty"`
}

type PortDescription struct {
	HttpPort string `yaml:",omitempty"`
	PeerPort string `yaml:",omitempty"`
}

type NodeApi struct {
	ConsensusApi string `yaml:",omitempty"`
	ExecutionApi string `yaml:",omitempty"`
}

type ChainConfig struct {
	Name string `yaml:",omitempty"`
	ID   string `yaml:",omitempty"`
}

type NodeConfigs struct {
	Chain                *ChainConfig      `yaml:",omitempty"`
	Configs              *DataVolume       `yaml:",omitempty"`
	Keystore             *DataVolume       `yaml:",omitempty"`
	Node                 *NodeDetails      `yaml:",omitempty"`
	Execution            *ClientDetails    `yaml:",omitempty"`
	Consensus            *ClientDetails    `yaml:",omitempty"`
	Validator            *ClientDetails    `yaml:",omitempty"`
	ValidatorCredentials *ValidatorSecrets `yaml:",omitempty"`
	ApiEndpoints         *NodeApi          `yaml:",omitempty"`

	Ports map[string]PortDescription `yaml:",omitempty"`
}

func (d *DataVolume) getVolume() string {
	return d.Volume
}

func (node *NodeDetails) getIP() string {
	return node.IP
}

func (node *NodeDetails) getName() string {
	return node.Name
}

func (cd *ClientDetails) getStatAddress() string {
	return cd.StatsAddress
}

func (cd *ClientDetails) getVerbosity() string {
	return cd.Verbosity
}

func (cd *ClientDetails) getEtherbase() string {
	return cd.Etherbase
}

func (cd *ClientDetails) getDataVolume() string {
	return cd.DataVolume
}

func (cd *ClientDetails) getNetworkID() string {
	return cd.NetworkId
}

func (cd *ClientDetails) getBootnode() string {
	return cd.Bootnode
}

func (cd *ClientDetails) getVersion() string {
	return cd.Version
}

func (pd PortDescription) getHttpPort() string {
	return pd.HttpPort
}

func (pd *PortDescription) getPeerPort() string {
	return pd.PeerPort
}

func (nc *NodeConfigs) getConfigs() *DataVolume {
	return nc.Configs
}

func (nc *NodeConfigs) getKeystore() *DataVolume {
	return nc.Keystore
}

func (nc *NodeConfigs) getNode() *NodeDetails {
	return nc.Node
}

func (nc *NodeConfigs) getExecution() *ClientDetails {
	return nc.Execution
}

func (nc *NodeConfigs) getConsensus() *ClientDetails {
	return nc.Consensus
}

func (nc *NodeConfigs) getValidator() *ClientDetails {
	return nc.Validator
}

func (nc *NodeConfigs) getPort(portName string) *PortDescription {
	if nc.Ports == nil {
		return nil
	}
	portDesc := nc.Ports[portName]
	return &portDesc
}

func (nc *NodeConfigs) GetValSecrets() *ValidatorSecrets {
	return nc.ValidatorCredentials
}

func (nc *NodeConfigs) WriteOrUpdateNodeConfig() error {
	yamlData, err := yaml.Marshal(nc)
	if err != nil {
		return err
	}
	return os.WriteFile(NodeConfigLocation, yamlData, os.ModePerm)
}

func (nc *NodeConfigs) UpdateBootnodes() error {
	chain := GetChainByString(nc.Chain.Name)
	GetChainByString(nc.Chain.Name)
	bootnodes, err := NewBootnodeUpdater(chain).DownloadLatestBootnodes()
	if err != nil {
		return err
	}

	if len(bootnodes) == 0 {
		fmt.Println("No bootnodes available for this chain ", chain.String())
	}

	hasUpdates := false
	if nc.Consensus.Bootnode != bootnodes[0].Consensus {
		fmt.Println("Updating bootnode for the consensus chain...")
		hasUpdates = true
		nc.Consensus.Bootnode = bootnodes[0].Consensus
	}
	if nc.Execution.Bootnode != bootnodes[0].Execution {
		fmt.Println("Updating bootnode for the execution chain...")
		hasUpdates = true
		nc.Execution.Bootnode = bootnodes[0].Execution
	}

	if !hasUpdates {
		fmt.Println("everything up to date")
	} else {
		err := nc.WriteOrUpdateNodeConfig()
		if err != nil {
			return err
		}
	}
	return nil
}

func MustGetNodeConfig() *NodeConfigs {
	if !FileExists(NodeConfigLocation) {
		cobra.CompErrorln("The node was not initialised yet. Call \n   lukso network init --nodeName NAME_OF_YOUR_NODE [--chain CHAIN_NAME] \n to setup a node.")
		os.Exit(1)
	}

	// Search config in home directory with name ".cli" (without extension).
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetConfigName("node_config")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("config file:", viper.ConfigFileUsed())
	} else {
		cobra.CompErrorln(err.Error())
		os.Exit(1)
	}

	config, err := getLoadedNodeConfigs()
	if err != nil {
		cobra.CompErrorln(err.Error())
		os.Exit(1)
	}
	return config
}

func getLoadedNodeConfigs() (*NodeConfigs, error) {
	var nodeConfig NodeConfigs
	err := viper.Unmarshal(&nodeConfig)
	if err != nil {
		return nil, err
	}
	return &nodeConfig, nil
}
