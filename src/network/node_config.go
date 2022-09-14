package network

import (
	"fmt"
	"github.com/lukso-network/lukso-cli/api/beaconapi"
	"github.com/lukso-network/lukso-cli/src/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

func GetDefaultNodeConfigByOptionParam(chain string) *NodeConfigs {
	return GetDefaultNodeConfig(GetChainByString(chain))
}

func GetDefaultNodeConfig(chain Chain) *NodeConfigs {
	switch chain {
	case L16:
		return DefaultL16NodeConfigs
	case Local:
		return DefaultLocalNodeConfigs
	case Dev:
		return DefaultDevNodeConfigs
	default:
		return DefaultL16NodeConfigs
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

type TransactionWallet struct {
	PublicKey  string `yaml:""`
	PrivateKey string `yaml:""`
}

func (config *NodeConfigs) getPort(portName string) *PortDescription {
	if config.Ports == nil {
		return nil
	}
	portDesc := config.Ports[portName]
	return &portDesc
}

func (config *NodeConfigs) CreateCredentials() *ValidatorCredentials {
	config.ValidatorCredentials = &ValidatorCredentials{}
	return config.ValidatorCredentials
}

func (config *NodeConfigs) WriteOrUpdateNodeConfig() error {
	yamlData, err := yaml.Marshal(config)
	if err != nil {
		return err
	}
	return os.WriteFile(NodeConfigLocation, yamlData, os.ModePerm)
}

func (config *NodeConfigs) UpdateExternalIP() (bool, error) {
	fmt.Println("Fetching external IP....")
	oldIP := config.Node.IP
	ip, err := getPublicIP()
	if err != nil {
		return false, err
	}

	if oldIP == ip {
		return false, nil
	} else {
		config.Node.IP = ip
		err := config.WriteOrUpdateNodeConfig()
		if err != nil {
			return false, err
		}
		return true, nil
	}
}

func (config *NodeConfigs) UpdateBootnodes() (bool, error) {
	chain := GetChainByString(config.Chain.Name)
	bootnodes, err := NewBootnodeUpdater(chain).DownloadLatestBootnodes()
	if err != nil {
		return false, err
	}

	if len(bootnodes) == 0 {
		fmt.Println("No bootnodes available for this chain ", chain.String())
	}

	hasUpdates := false
	if config.Consensus.Bootnode != bootnodes[0].Consensus {
		fmt.Println("Updating bootnode for the consensus chain...")
		hasUpdates = true
		config.Consensus.Bootnode = bootnodes[0].Consensus
	}
	if config.Execution.Bootnode != bootnodes[0].Execution {
		fmt.Println("Updating bootnode for the execution chain...")
		hasUpdates = true
		config.Execution.Bootnode = bootnodes[0].Execution
	}

	if !hasUpdates {
		return false, nil
	} else {
		err := config.WriteOrUpdateNodeConfig()
		if err != nil {
			return false, err
		}
		return true, nil
	}
}

func (config *NodeConfigs) UpdateDevBootnodes(devLocation string) (bool, error) {
	chain := GetChainByString(config.Chain.Name)
	GetChainByString(config.Chain.Name)
	bootnodes, err := NewBootnodeUpdaterDev(chain, devLocation).DownloadLatestBootnodes()
	if err != nil {
		return false, err
	}

	if len(bootnodes) == 0 {
		fmt.Println("No bootnodes available for this chain ", chain.String())
	}

	hasUpdates := false
	if config.Consensus.Bootnode != bootnodes[0].Consensus {
		fmt.Println("Updating bootnode for the consensus chain...")
		hasUpdates = true
		config.Consensus.Bootnode = bootnodes[0].Consensus
	}
	if config.Execution.Bootnode != bootnodes[0].Execution {
		fmt.Println("Updating bootnode for the execution chain...")
		hasUpdates = true
		config.Execution.Bootnode = bootnodes[0].Execution
	}

	if !hasUpdates {
		return false, nil
	} else {
		err := config.WriteOrUpdateNodeConfig()
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
	//TODO: Change from node_config to client_configuration
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

func (config *NodeConfigs) HasMnemonic() bool {
	if config.ValidatorCredentials == nil {
		return false
	}

	return config.ValidatorCredentials.ValidatorMnemonic != ""
}

func (config *NodeConfigs) GetChain() Chain {
	return GetChainByString(config.Chain.Name)
}

func GetENRFromBootNode(endpoint string) (string, error) {
	response, err := beaconapi.NewBeaconClient(endpoint).Identity()
	if err != nil {
		utils.PrintColoredErrorWithReason("couldn't get bootnode enr", err)
		return "", err
	}
	return strings.TrimSuffix(response.Data.Enr, "=="), nil
}

func (config *NodeConfigs) CreateNodeRecovery() NodeRecovery {
	var tw TransactionWallet
	if config.TransactionWallet == nil {
		tw = TransactionWallet{
			PublicKey:  "",
			PrivateKey: "",
		}
	} else {
		tw = *config.TransactionWallet
	}
	return NodeRecovery{
		*config.ValidatorCredentials,
		tw,
	}
}
