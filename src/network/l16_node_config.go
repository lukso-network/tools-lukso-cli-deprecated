package network

import "os"

// Generate node IP and Hostname dynamically
var myPublicIP, _ = getPublicIP()
var myHostName, _ = os.Hostname()

var DefaultL16NodeConfigs = &NodeConfigs{
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
		StatsAddress: "ethstats.l16.d.lukso.dev",
		Verbosity:    "3",
		Etherbase:    "0x7781121fd00A009670E31b76A2bf99b3A2D6878D",
		DataVolume:   "./data/execution_data",
		NetworkId:    "19051978",
		Bootnode:     "enode://45f5c9d2bec9253f82f44a385692f22a69fdddb5bc5ada29176c3e5977d659529387770e77548dcaad668ce0ef0c74f994ca61a0d55bdfbdffb3813b98c3f7ea@34.91.62.48:30303",
	},
	Consensus: &ClientDetails{
		StatsAddress: "35.202.229.165:9090",
		Verbosity:    "info",
		Version:      "v0.1.4-dev",
		DataVolume:   "./data/consensus_data",
		Bootnode:     "enr:-MK4QEkSFE8VA-mbzlHLCgyfpvfh7gFPs1AQNrLxNtHSMq58ChKNtwDz7hVkW8qYUGs71tmSeP1buAGJYYZvj_Hp_8yGAX8tWHtih2F0dG5ldHOIAAAAAAAAAACEZXRoMpC3QoawYgAAcf__________gmlkgnY0gmlwhCJ2PI-Jc2VjcDI1NmsxoQMOCac5xZmiy984EK04FgtH-ijxkjdvcXZZvayai71U94hzeW5jbmV0cwCDdGNwgjLIg3VkcIIu4A",
	},
	Validator: &ClientDetails{
		DataVolume: "./data/validator_data",
	},
	Ports: map[string]PortDescription{
		"geth": {
			HttpPort: "8545",
			PeerPort: "30303",
		},
	},
}
