/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/lukso-network/lukso-cli/src/network"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// validatorDescribeByKeyCmd represents the check command
var validatorDescribeByKeyCmd = &cobra.Command{
	Use:     "byKey",
	Short:   "Show detailed status of the validators as deposited in the DepositDetails Contract",
	Long:    `Show detailed status of the validators as deposited in the DepositDetails Contract.`,
	Example: "lukso network validator check",
	Run: func(cmd *cobra.Command, args []string) {
		// get node conf from --chain param or get default chain
		nodeConf := network.GetDefaultNodeConfigByOptionParam(viper.GetString(CommandOptionChain))

		key, err := cmd.Flags().GetString("key")
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}
		consensusApi, err := readConsensusApiEndpoint(cmd, nodeConf)
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}
		executionApi, err := readExecutionApiEndpoint(cmd, nodeConf)
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}
		address, err := readDepositAddress(cmd, nodeConf)
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}

		err = network.DescribeValidatorKey([]string{key}, address, executionApi, consensusApi, nil)
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}
	},
}

func init() {
	validatorDescribeCmd.AddCommand(validatorDescribeByKeyCmd)
	validatorDescribeByKeyCmd.Flags().StringP(CommandOptionExecutionApi, CommandOptionExecutionApiShort, "", "execution api endpoint")
	validatorDescribeByKeyCmd.Flags().StringP(CommandOptionConsensusApi, CommandOptionConsensusApiShort, "", "consensus api endpoint")
	validatorDescribeByKeyCmd.Flags().StringP("key", "k", "", "validator key to be described - keep empty to describe your validators defined in keystore")
	validatorDescribeByKeyCmd.Flags().StringP(CommandOptionDepositAddress, CommandOptionDepositAddressShort, "", "the address of the deposit contract")

	validatorDescribeByKeyCmd.MarkFlagRequired("key")
}
