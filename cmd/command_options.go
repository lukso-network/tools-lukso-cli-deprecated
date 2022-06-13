package cmd

import (
	"fmt"
	"github.com/lukso-network/lukso-cli/src/network"
	"github.com/lukso-network/lukso-cli/src/network/types"
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

	CommandOptionFrom      = "from"
	CommandOptionFromShort = "f"
	CommandOptionTo        = "to"
	CommandOptionToShort   = "t"

	CommandOptionMaxGasFee        = "maxGasFee"
	CommandOptionMaxGasFeeShort   = "m"
	CommandOptionPriorityFee      = "priorityFee"
	CommandOptionPriorityFeeShort = "p"

	CommandOptionPath      = "path"
	CommandOptionPathShort = "p"
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

func readRangeFromCommand(cmd *cobra.Command) (types.ValidatorRange, error) {
	vRange := types.ValidatorRange{}
	from, err := cmd.Flags().GetInt64(CommandOptionFrom)
	if err != nil {
		return vRange, err
	}
	to, err := cmd.Flags().GetInt64(CommandOptionTo)
	if err != nil {
		return vRange, err
	}

	// default to 0
	if from == -1 {
		from = 0
	}

	if to == -1 {
		return vRange, fmt.Errorf("--to not given")
	}

	if to <= from {
		return vRange, fmt.Errorf("--to (%d) must be greater than --from (%d)", to, from)
	}

	vRange.From = from
	vRange.To = to
	return vRange, nil
}
