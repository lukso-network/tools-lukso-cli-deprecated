package network

import (
	"encoding/json"
	"fmt"
)

const BootNodeLocation = "https://raw.githubusercontent.com/lukso-network/network-configs/main/devnets/%s/%s/%s"

type Bootnode struct {
	Consensus string `json:"consensus"`
	Execution string `json:"execution"`
}

type BootnodeUpdater struct {
	chainIdentifier string
}

func NewBootnodeUpdater(chain Chain) BootnodeUpdater {
	return BootnodeUpdater{chain.String()}
}

func NewBootnodeUpdaterDev(chain Chain, devLocation string) BootnodeUpdater {
	return BootnodeUpdater{fmt.Sprintf("%s/%s", chain.String(), devLocation)}
}

func (b BootnodeUpdater) DownloadLatestBootnodes() ([]Bootnode, error) {
	url := fmt.Sprintf(BootNodeLocation, b.chainIdentifier, BootnodesDirectory, BootnodeJSONName)

	fmt.Println("Fetching bootnodes from", url)
	bytes, err := downloadFileOverHttp(url)
	if err != nil {
		return nil, err
	}
	var result []Bootnode
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
