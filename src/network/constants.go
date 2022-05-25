package network

const (
	DefaultDirectoryNameKeystore = "keystore"

	ChainL16Beta = "l16beta"
	ChainL16     = "l16"
	ChainMainNet = "mainnet"
	ChainLocal   = "local"

	CommandOptionChain              = "chain"
	CommandOptionNodeConf           = "nodeconf"
	CommandOptionNodeName           = "nodeName"
	ErrMsgValidatorSecretNotPresent = "validator secret not present"
	ConfigDirectory                 = "configs"
	BootnodesDirectory              = "bootnode"
	ConfigBranchName                = "main"
	ConfigRepoName                  = "network-configs"
	BootnodeJSONName                = "bootnodes.json"
	GitUrl                          = "https://raw.githubusercontent.com/lukso-network/"
)

var (
	NetworkSetupFiles = []string{"docker-compose.yml"}
	ConfigFiles       = []string{"config.yaml", "genesis.json", "genesis.ssz"}
)
