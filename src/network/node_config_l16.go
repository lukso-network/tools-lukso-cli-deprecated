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
		Etherbase:    "0xCce25D1620bD5E33485B808B30D3c805eA28dBe3",
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
