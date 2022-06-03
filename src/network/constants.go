package network

const (
	DefaultDirectoryNameKeystore = "keystore"

	ChainL16Beta = "l16beta"
	ChainL16     = "l16"
	ChainMainNet = "mainnet"
	ChainLocal   = "local"

	ErrMsgValidatorSecretNotPresent = "Validator secret not present, did you call lukso network validator setup first?"
	ConfigDirectory                 = "configs"
	BootnodesDirectory              = "bootnode"
	ConfigBranchName                = "main"
	ConfigRepoName                  = "network-configs"
	BootnodeJSONName                = "bootnodes.json"
	GitUrl                          = "https://raw.githubusercontent.com/lukso-network/"
	NodeConfigLocation              = "./node_config.yaml"
)

var (
	ConfigFiles = []string{"config.yaml", "genesis.json", "genesis.ssz"}
)
