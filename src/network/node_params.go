package network

import (
	"encoding/json"
	"fmt"
)

const NodeParamsDevLocation = "https://raw.githubusercontent.com/lukso-network/network-configs/main/dev/%s/%s"
const NodeParamsLocation = "https://raw.githubusercontent.com/lukso-network/network-configs/main/%s/%s"

type NodeParams struct {
	ExecutionAPI   string `json:"executionApi"`
	ConsensusAPI   string `json:"consensusApi"`
	ExecutionStats string `json:"executionStats"`
	ConsensusStats string `json:"consensusStats"`
	NetworkID      string `json:"networkId"`
	MinStakeAmount string `json:"minStakeAmount"`
	GethVersion    string `json:"gethVersion"`
	PrysmVersion   string `json:"prysmVersion"`
}

func NewNodeParamsLoader() *NodeParams {
	return &NodeParams{}
}

func (n *NodeParams) GetDevLocation(devConfig string) string {
	return fmt.Sprintf(NodeParamsDevLocation, devConfig, NodeParamsFileName)
}

func (n *NodeParams) GetLocation(chain Chain) string {
	return fmt.Sprintf(NodeParamsLocation, chain, NodeParamsFileName)
}

func (n *NodeParams) LoadNodeParams(location string) error {
	bytes, err := downloadFileOverHttp(location)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, n)
}
