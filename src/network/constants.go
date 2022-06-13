package network

const (
	ChainL16Beta = "l16beta"
	ChainL16     = "l16"
	ChainMainNet = "mainnet"
	ChainLocal   = "local"
	ChainDev     = "dev"

	ErrMsgValidatorSecretNotPresent = "Validator secret not present, did you call lukso network validator setup first?"
	ConfigDirectory                 = "configs"
	BootnodesDirectory              = "bootnode"
	ConfigBranchName                = "main"
	ConfigRepoName                  = "network-configs"
	NodeParamsFileName              = "node_params.json"
	BootnodeJSONName                = "bootnodes.json"
	GitUrl                          = "https://raw.githubusercontent.com/lukso-network/"
	NodeConfigLocation              = "./node_config.yaml"
	NodeRecoveryFileLocation        = "./node_recovery.json"
)
