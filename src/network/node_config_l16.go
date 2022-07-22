package network

var DefaultL16NodeConfigs = &NodeConfigs{
	Chain: &ChainConfig{
		Name: ChainL16,
		ID:   "2828",
	},
	Configs: &DataVolume{
		Volume: "./configs",
	},
	Keystore: &DataVolume{
		Volume: "./keystore",
	},
	Node: &NodeDetails{},
	Execution: &ClientDetails{
		StatsAddress: "",
		Verbosity:    "3",
		Version:      "v0.2.1",
		Etherbase:    "0x7781121fd00A009670E31b76A2bf99b3A2D6878D",
		DataVolume:   "./data/execution_data",
		Bootnode:     "",
	},
	Consensus: &ClientDetails{
		StatsAddress: "",
		Verbosity:    "info",
		Version:      "v0.2.3-dev",
		DataVolume:   "./data/consensus_data",
		Bootnode:     "",
	},
	Validator: &ClientDetails{
		DataVolume: "./data/validator_data",
	},
	ValidatorCredentials: &ValidatorCredentials{},
	TransactionWallet:    &TransactionWallet{},
	Ports: map[string]PortDescription{
		"geth": {
			HttpPort: "8545",
			PeerPort: "30303",
		},
	},
	ApiEndpoints: &NodeApi{
		ConsensusApi: "https://beacon.l16.lukso.network",
		ExecutionApi: "https://rpc.l16.lukso.network",
	},
	DepositDetails: &DepositDetails{
		Amount:              "220000000000",
		ContractAddress:     "0x000000000000000000000000000000000000cafe",
		DepositFileLocation: "./deposit_data.json",
		ForkVersion:         "0x60000069",
	},
}
