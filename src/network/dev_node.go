package network

import (
	"encoding/json"
	"fmt"
)

const NodeParamsLocation = "https://raw.githubusercontent.com/lukso-network/network-configs/main/dev/%s/%s"

type NodeParams struct {
	ExecutionAPI   string `json:"executionApi"`
	ConsensusAPI   string `json:"consensusApi"`
	ExecutionStats string `json:"executionStats"`
	ConsensusStats string `json:"consensusStats"`
	NetworkID      string `json:"networkId"`
	MinStakeAmount string `json:"minStakeAmount"`
}

func NewNodeParams() *NodeParams {
	return &NodeParams{}
}

func (n *NodeParams) GetLocation(devConfig string) string {
	return fmt.Sprintf(NodeParamsLocation, devConfig, NodeParamsFileName)
}

func (n *NodeParams) LoadNodeParams(location string) error {
	bytes, err := downloadFileOverHttp(location)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, n)
}
