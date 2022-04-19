package network

var BetaDefaultValSecrets = &ValidatorSecrets{
	ForkVersion: "0x60000069",
	Deposit: &DepositDetails{
		Amount:              "32000000000",
		ContractAddress:     "0x4242424242424242424242424242424242424242",
		Force:               true,
		DepositFileLocation: "./deposit_data.json",
	},
	Eth1Data: &Eth1Details{
		RPCEndPoint:   "http://34.121.22.6:8545",
		WalletAddress: "",
		WalletPrivKey: "",
	},
}
