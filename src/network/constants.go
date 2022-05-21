package network

const (
	DefaultDirectoryNameKeystore = "keystore"

	ChainL16Beta = "l16beta"
	ChainL16     = "l16"
	ChainMainNet = "mainnet"

	CommandOptionChainID            = "chain"
	CommandOptionNodeConf           = "nodeconf"
	CommandOptionNodeName           = "nodeName"
	ErrMsgValidatorSecretNotPresent = "validator secret not present"
	ConfigDirectory                 = "configs"
	ConfigBranchName                = "main"
	ConfigRepoName                  = "network-configs"
	GitUrl                          = "https://raw.githubusercontent.com/lukso-network/"
)

var (
	NetworkSetupFiles = []string{"docker-compose.yml"}
	ConfigFiles       = []string{"config.yaml", "genesis.json", "genesis.ssz"}
)
