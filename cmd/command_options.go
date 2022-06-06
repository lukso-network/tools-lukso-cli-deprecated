package cmd

import (
	"github.com/lukso-network/lukso-cli/src/network"
	"github.com/spf13/cobra"
)

const (
	CommandOptionConsensusApi        = "consensusApi"
	CommandOptionConsensusApiShort   = "c"
	CommandOptionExecutionApi        = "executionApi"
	CommandOptionExecutionApiShort   = "e"
	CommandOptionDepositAddress      = "address"
	CommandOptionDepositAddressShort = "a"

	CommandOptionChain     = "chain"
	CommandOptionNodeName  = "nodeName"
	CommandOptionDevConfig = "devConfig"
)

func readConsensusApiEndpoint(cmd *cobra.Command, nodeConf *network.NodeConfigs) (string, error) {
	flag, err := cmd.Flags().GetString(CommandOptionConsensusApi)
	if err != nil {
		return "", err
	}
	// user entered value
	if flag != "" {
		return flag, nil
	}

	return nodeConf.ApiEndpoints.ConsensusApi, nil
}

func readExecutionApiEndpoint(cmd *cobra.Command, nodeConf *network.NodeConfigs) (string, error) {
	flag, err := cmd.Flags().GetString(CommandOptionExecutionApi)
	if err != nil {
		return "", err
	}
	// user entered value
	if flag != "" {
		return flag, nil
	}

	return nodeConf.ApiEndpoints.ExecutionApi, nil
}

func readDepositAddress(cmd *cobra.Command, nodeConf *network.NodeConfigs) (string, error) {
	flag, err := cmd.Flags().GetString(CommandOptionDepositAddress)
	if err != nil {
		return "", err
	}
	// user entered value
	if flag != "" {
		return flag, nil
	}

	return nodeConf.DepositDetails.ContractAddress, nil
}
