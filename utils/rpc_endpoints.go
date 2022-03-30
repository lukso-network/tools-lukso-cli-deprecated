package utils

import (
	"encoding/json"
	"github.com/lukso-network/lukso-cli/utils/gethrpc"
)

type AdminPeersResponse = []Peer

type AdminNodeInfo struct {
	Enode      string `json:"enode"`
	Enr        string `json:"enr"`
	ID         string `json:"id"`
	IP         string `json:"ip"`
	ListenAddr string `json:"listenAddr"`
	Name       string `json:"name"`
	Ports      struct {
		Discovery int `json:"discovery"`
		Listener  int `json:"listener"`
	} `json:"ports"`
	Protocols struct {
		Eth struct {
			Config struct {
				BerlinBlock             int    `json:"berlinBlock"`
				ByzantiumBlock          int    `json:"byzantiumBlock"`
				ChainID                 int    `json:"chainId"`
				ConstantinopleBlock     int    `json:"constantinopleBlock"`
				Eip150Block             int    `json:"eip150Block"`
				Eip150Hash              string `json:"eip150Hash"`
				Eip155Block             int    `json:"eip155Block"`
				Eip158Block             int    `json:"eip158Block"`
				HomesteadBlock          int    `json:"homesteadBlock"`
				IstanbulBlock           int    `json:"istanbulBlock"`
				LondonBlock             int    `json:"londonBlock"`
				MergeForkBlock          int    `json:"mergeForkBlock"`
				PetersburgBlock         int    `json:"petersburgBlock"`
				TerminalTotalDifficulty int    `json:"terminalTotalDifficulty"`
			} `json:"config"`
			Difficulty int    `json:"difficulty"`
			Genesis    string `json:"genesis"`
			Head       string `json:"head"`
			Network    int    `json:"network"`
		} `json:"eth"`
		Snap struct {
		} `json:"snap"`
	} `json:"protocols"`
}

type Peer struct {
	Enode   string   `json:"enode"`
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Caps    []string `json:"caps"`
	Network struct {
		LocalAddress  string `json:"localAddress"`
		RemoteAddress string `json:"remoteAddress"`
		Inbound       bool   `json:"inbound"`
		Trusted       bool   `json:"trusted"`
		Static        bool   `json:"static"`
	} `json:"network"`
	Protocols struct {
		Eth struct {
			Version    int    `json:"version"`
			Difficulty int    `json:"difficulty"`
			Head       string `json:"head"`
		} `json:"eth"`
		Snap struct {
			Version int `json:"version"`
		} `json:"snap"`
	} `json:"protocols"`
}

/*
	admin_peers
*/
func AdminPeersRequest(client *gethrpc.Instance) ([]Peer, error) {
	response := make([]Peer, 0)
	result, err := gethrpc.CheckRPCError(client.Call("admin_peers"))
	if err != nil {
		return response, err
	}
	bytes, err := result.ToJSONBytes()
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
	admin_nodeInfo
*/
func AdminNodeInfoRequest(client *gethrpc.Instance) (*AdminNodeInfo, error) {
	response := &AdminNodeInfo{}
	result, err := gethrpc.CheckRPCError(client.Call("admin_nodeInfo"))
	if err != nil {
		return response, err
	}
	bytes, err := result.ToJSONBytes()
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(bytes, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func BlockNumber(client *gethrpc.Instance) (int64, error) {
	return client.RequestInt64("eth_blockNumber")
}
