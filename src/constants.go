package src

const (
	ConfigBranchName = "main"
	NetworkVersion   = "19"
	DefaultNetworkID = "l16"
	ConfigRepoName   = "network-configs"
	GitUrl           = "https://raw.githubusercontent.com/lukso-network/"
)

// Supported ChainIDs
const (
	L16Network = "l16"
)

// Keys that are used in Viper
const (
	ViperKeyNetworkName = "NETWORK_NAME"
)

var (
	NetworkSetupFiles = []string{"docker-compose.yml"}
	ConfigFiles       = []string{"config.yaml", "genesis.json", "genesis.ssz"}
)

const (
	ErrMsgValidatorSecretNotPresent = "validator secret not present"
)
