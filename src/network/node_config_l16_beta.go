package network

import "os"

// Generate node IP and Hostname dynamically
var myPublicIP, _ = getPublicIP()
var myHostName, _ = os.Hostname()

var DefaultL16BetaNodeConfigs = &NodeConfigs{
	Chain: &ChainConfig{
		Name: ChainL16Beta,
		ID:   "83748374",
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
		StatsAddress: "35.204.4.181",
		Verbosity:    "3",
		Version:      "v0.2.0-dev",
		Etherbase:    "0x7781121fd00A009670E31b76A2bf99b3A2D6878D",
		DataVolume:   "./data/execution_data",
		NetworkId:    "83748374",
		Bootnode:     "enode://c2bb19ce658cfdf1fecb45da599ee6c7bf36e5292efb3fb61303a0b2cd07f96c20ac9b376a464d687ac456675a2e4a44aec39a0509bcb4b6d8221eedec25aca2@34.91.51.22:30303",
	},
	Consensus: &ClientDetails{
		StatsAddress: "34.141.143.70:9090",
		Verbosity:    "info",
		Version:      "v0.2.3-dev",
		DataVolume:   "./data/consensus_data",
		Bootnode:     "enr:-MK4QDjZAzD165YMiCM6WrCdRN3EGDViJqP_v1fyOE-YuOYAQ102NnF8GjTJG_-yUUMwWKuEzRMJMIoObbgDB3jtZA6GAYCJeizfh2F0dG5ldHOIAAAAAAAAAACEZXRoMpCBQMXLYgAAcf__________gmlkgnY0gmlwhCJbpv2Jc2VjcDI1NmsxoQOGOP29TdRKCr40kT-ZgsvrMYIs6EIwuYGEq11txAj_VohzeW5jbmV0cwCDdGNwgjLIg3VkcIIu4A",
	},
	Validator: &ClientDetails{
		DataVolume: "./data/validator_data",
	},
	ValidatorCredentials: &ValidatorCredentials{},
	Ports: map[string]PortDescription{
		"geth": {
			HttpPort: "8545",
			PeerPort: "30303",
		},
	},
	ApiEndpoints: &NodeApi{
		ConsensusApi: "https://beacon.beta.l16.lukso.network",
		ExecutionApi: "https://rpc.beta.l16.lukso.network",
	},
	DepositDetails: &DepositDetails{
		Amount:              "32000000000",
		ContractAddress:     "0x4242424242424242424242424242424242424242",
		DepositFileLocation: "./deposit_data.json",
		ForkVersion:         "0x60000069",
	},
}
