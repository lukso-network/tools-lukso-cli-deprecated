package src

const (
	ConfigBranchName = "l16-dev"
	NetworkVersion   = "17"
	DefaultNetworkID = "l16"
	ConfigRepoName   = "network-configs"
	GitUrl           = "https://raw.githubusercontent.com/lukso-network/"
)

// Keys that are used in Viper
const (
	ViperKeyNetworkName = "NETWORK_NAME"
)

var (
	NetworkSetupFiles = []string{".env", "docker-compose.yml", "secrets.env", "send_deposit.sh"}
	ConfigFiles       = []string{"config.yaml", "genesis.json", "genesis.ssz"}
)
