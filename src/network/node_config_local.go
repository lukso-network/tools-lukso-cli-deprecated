package network

var DefaultLocalNodeConfigs = &NodeConfigs{
	Chain: &ChainConfig{
		Name: ChainLocal,
		ID:   "190578",
	},
	Configs: &DataVolume{
		Volume: "./configs",
	},
	Keystore: &DataVolume{
		Volume: "./keystore",
	},
	Node: &NodeDetails{
		IP:   myPublicIP,
		Name: myHostName,
	},
	Execution: &ClientDetails{
		StatsAddress: "127.0.0.1",
		Verbosity:    "3",
		Version:      "v0.2.0-dev",
		Etherbase:    "0x7781121fd00A009670E31b76A2bf99b3A2D6878D",
		DataVolume:   "./data/execution_data",
		NetworkId:    "19191919",
		Bootnode:     "",
	},
	Consensus: &ClientDetails{
		StatsAddress: "127.0.0.1",
		Verbosity:    "info",
		Version:      "v0.2.3-dev",
		DataVolume:   "./data/consensus_data",
		Bootnode:     "",
	},
	Validator: &ClientDetails{
		DataVolume: "./data/validator_data",
	},
	ValidatorCredentials: BetaDefaultValSecrets,
	Ports: map[string]PortDescription{
		"geth": {
			HttpPort: "8545",
			PeerPort: "30303",
		},
	},
	ApiEndpoints: &NodeApi{
		ConsensusApi: "http://localhost:3500",
		ExecutionApi: "http://localhost:8545",
	},
}
