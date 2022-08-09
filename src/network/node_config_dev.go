package network

var DefaultDevNodeConfigs = &NodeConfigs{
	Chain: &ChainConfig{
		Name: ChainDev,
		ID:   "00000000",
	},
	Configs: &DataVolume{
		Volume: "./configs",
	},
	Keystore: &DataVolume{
		Volume: "./keystore",
	},
	Node: &NodeDetails{},
	Execution: &ClientDetails{
		StatsAddress: "EXECUTION_STATS_ADDRESS",
		Verbosity:    "3",
		Version:      "v0.2.0-dev",
		Etherbase:    "0xCce25D1620bD5E33485B808B30D3c805eA28dBe3",
		DataVolume:   "./data/execution_data",
		Bootnode:     "EXECUTION_BOOTNODE",
	},
	Consensus: &ClientDetails{
		StatsAddress: "CONSENSUS_STATS_ADDRESS",
		Verbosity:    "info",
		Version:      "v0.2.3-dev",
		DataVolume:   "./data/consensus_data",
		Bootnode:     "CONSENSUS_BOOTNODE",
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
		ConsensusApi: "ConsensusApi",
		ExecutionApi: "ExecutionApi",
	},
	DepositDetails: &DepositDetails{
		Amount:              "Amount",
		ContractAddress:     "0x4242424242424242424242424242424242424242",
		DepositFileLocation: "./deposit_data.json",
		ForkVersion:         "0x60000069",
	},
}
