package network

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
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
	case Dev:
		return DefaultDevNodeConfigs
	default:
		return DefaultL16BetaNodeConfigs
	}
}

type DataVolume struct {
	Volume string `yaml:""`
}

type NodeDetails struct {
	IP   string `yaml:""`
	Name string `yaml:""`
}

type ClientDetails struct {
	StatsAddress string `yaml:""`
	Verbosity    string `yaml:""`
	Etherbase    string `yaml:""`
	DataVolume   string `yaml:""`
	NetworkId    string `yaml:""`
	Bootnode     string `yaml:""`
	Version      string `yaml:""`
}

type PortDescription struct {
	HttpPort string `yaml:""`
	PeerPort string `yaml:""`
}

type NodeApi struct {
	ConsensusApi string `yaml:""`
	ExecutionApi string `yaml:""`
}

type ChainConfig struct {
	Name string `yaml:""`
	ID   string `yaml:""`
}

type NodeConfigs struct {
	Chain                *ChainConfig          `yaml:""`
	Configs              *DataVolume           `yaml:""`
	Keystore             *DataVolume           `yaml:""`
	Node                 *NodeDetails          `yaml:""`
	Execution            *ClientDetails        `yaml:""`
	Consensus            *ClientDetails        `yaml:""`
	Validator            *ClientDetails        `yaml:""`
	ValidatorCredentials *ValidatorCredentials `yaml:""`
	ApiEndpoints         *NodeApi              `yaml:""`
	TransactionWallet    *TransactionWallet    `yaml:""`
	DepositDetails       *DepositDetails       `yaml:""`

	Ports map[string]PortDescription `yaml:""`
}

type NodeConfigsV0 struct {
	Configs              *DataVolume           `yaml:""`
	Keystore             *DataVolume           `yaml:""`
	Node                 *NodeDetails          `yaml:""`
	Execution            *ClientDetails        `yaml:""`
	Consensus            *ClientDetails        `yaml:""`
	Validator            *ClientDetails        `yaml:""`
	ValidatorCredentials *ValidatorCredentials `yaml:""`
	ApiEndpoints         *NodeApi              `yaml:""`

	Ports map[string]PortDescription `yaml:""`
}

type TransactionWallet struct {
	PublicKey  string `yaml:""`
	PrivateKey string `yaml:""`
}

func (nc *NodeConfigs) getPort(portName string) *PortDescription {
	if nc.Ports == nil {
		return nil
	}
	portDesc := nc.Ports[portName]
	return &portDesc
}

func (nc *NodeConfigs) CreateCredentials() *ValidatorCredentials {
	nc.ValidatorCredentials = &ValidatorCredentials{}
	return nc.ValidatorCredentials
}

func (nc *NodeConfigs) WriteOrUpdateNodeConfig() error {
	yamlData, err := yaml.Marshal(nc)
	if err != nil {
		return err
	}
	return os.WriteFile(NodeConfigLocation, yamlData, os.ModePerm)
}

func (nc *NodeConfigs) UpdateExternalIP() (bool, error) {
	fmt.Println("Fetching external IP....")
	oldIP := nc.Node.IP
	ip, err := getPublicIP()
	if err != nil {
		return false, err
	}

	if oldIP == ip {
		return false, nil
	} else {
		nc.Node.IP = ip
		err := nc.WriteOrUpdateNodeConfig()
		if err != nil {
			return false, err
		}
		return true, nil
	}
}

func (nc *NodeConfigs) UpdateBootnodes() (bool, error) {
	chain := GetChainByString(nc.Chain.Name)
	bootnodes, err := NewBootnodeUpdater(chain).DownloadLatestBootnodes()
	if err != nil {
		return false, err
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
		return false, nil
	} else {
		err := nc.WriteOrUpdateNodeConfig()
		if err != nil {
			return false, err
		}
		return true, nil
	}
}

func (nc *NodeConfigs) InitDevBootnodes(devLocation string) (bool, error) {
	chain := GetChainByString(nc.Chain.Name)
	GetChainByString(nc.Chain.Name)
	bootnodes, err := NewBootnodeUpdaterDev(chain, devLocation).DownloadLatestBootnodes()
	if err != nil {
		return false, err
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
		return false, nil
	} else {
		err := nc.WriteOrUpdateNodeConfig()
		if err != nil {
			return false, err
		}
		return true, nil
	}
}

/*
	EnsureNetworkIsInitialised

	Checks if the node_conf.yaml file was created otherwise it exits
*/
func EnsureNetworkIsInitialised() {
	if !FileExists(NodeConfigLocation) {
		cobra.CompErrorln("The node was not initialised yet. Call \n   lukso network init --nodeName NAME_OF_YOUR_NODE [--chain CHAIN_NAME] \n to setup a node.")
		os.Exit(1)
	}
}

func MustGetNodeConfig() *NodeConfigs {
	EnsureNetworkIsInitialised()

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

	config, err := readNodeConfigsFromFile()
	if err != nil {
		cobra.CompErrorln(err.Error())
		os.Exit(1)
	}
	return config
}

func LoadNodeConf() (*NodeConfigs, error) {
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetConfigName("node_config")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("config file:", viper.ConfigFileUsed())
	} else {
		return nil, err
	}

	return readNodeConfigsFromFile()
}

func readNodeConfigsFromFile() (*NodeConfigs, error) {
	return getConf()
	//var nodeConfig NodeConfigs
	//err := viper.Unmarshal(&nodeConfig)
	//if err != nil {
	//	return nil, err
	//}
	//return &nodeConfig, nil
}

func LoadNodeConfV0() (*NodeConfigsV0, error) {
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetConfigName("node_config")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
	} else {
		return nil, err
	}

	var nodeConfig NodeConfigsV0
	err := viper.Unmarshal(&nodeConfig)
	if err != nil {
		return nil, err
	}
	return &nodeConfig, nil
}

func (n NodeConfigsV0) Upgrade(chain Chain) *NodeConfigs {
	defaultConfig := GetDefaultNodeConfig(chain)
	return &NodeConfigs{
		Chain:                defaultConfig.Chain,
		Configs:              n.Configs,
		Keystore:             n.Keystore,
		Node:                 n.Node,
		Execution:            n.Execution,
		Consensus:            n.Consensus,
		Validator:            n.Validator,
		ValidatorCredentials: n.ValidatorCredentials,
		ApiEndpoints:         n.ApiEndpoints,
		Ports:                n.Ports,
	}
}

func getConf() (*NodeConfigs, error) {
	yamlFile, err := ioutil.ReadFile(NodeConfigLocation)
	if err != nil {
		return nil, err
	}
	node := &NodeConfigs{}
	err = yaml.Unmarshal(yamlFile, node)
	if err != nil {
		return nil, err
	}

	return node, nil
}

func (config *NodeConfigs) Save() error {
	yamlData, err := yaml.Marshal(config)
	if err != nil {
		return err
	}
	return os.WriteFile(NodeConfigLocation, yamlData, os.ModePerm)
}
