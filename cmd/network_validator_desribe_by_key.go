/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/lukso-network/lukso-cli/api/beaconapi"
	"github.com/lukso-network/lukso-cli/src/network"
	"github.com/lukso-network/lukso-cli/src/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
)

// validatorDescribeByKeyCmd represents the check command
var validatorDescribeByKeyCmd = &cobra.Command{
	Use:     "byKey",
	Short:   "Show detailed status of the validators as deposited in the Deposit Contract",
	Long:    `Show detailed status of the validators as deposited in the Deposit Contract.`,
	Example: "lukso network validator check",
	Run: func(cmd *cobra.Command, args []string) {
		// get node conf from --chain param or get default chain
		nodeConf := network.GetDefaultNodeConfigByOptionParam(viper.GetString(network.CommandOptionChain))

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

		err = describeValidatorKey([]string{key}, address, executionApi, consensusApi, nil)
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}
	},
}

func describeValidatorKey(keys []string, contractAddress string, executionApi string, consensusApi string, depositEvents *network.DepositEvents) (err error) {
	fmt.Println("Configuration")
	fmt.Println("........................................................................................................................................................................")
	utils.ColoredPrintln("Consensus Api:", consensusApi)
	utils.ColoredPrintln("Execution Api:", executionApi)
	utils.ColoredPrintln("Contract Address:", contractAddress)
	fmt.Println("........................................................................................................................................................................")

	beaconClient := beaconapi.NewBeaconClient(consensusApi)
	fmt.Println("Getting all deposits from contract....")

	if depositEvents == nil {
		e, err := network.NewDepositEvents(contractAddress, executionApi)
		if err != nil {
			return err
		}
		depositEvents = &e
	}

	for _, k := range keys {
		key := maybeAddHexPrefix(k)
		fmt.Printf("Checking state of validator key: %v.......", key)
		amount := depositEvents.Amount(key)
		if amount == 0 {
			fmt.Println("   not deposited yet")
			fmt.Println("")
			continue
		}

		state, status, err := beaconClient.ValidatorState(key, -1)
		if status == http.StatusNotFound {
			fmt.Println("  is pending")
			fmt.Println("")
			continue
		}
		if err != nil {
			return err
		}
		fmt.Println("s")
		utils.ColoredPrintln("ValidatorKey", state.Data.Validator.Pubkey)
		utils.ColoredPrintln("Index:", state.Data.Index)
		utils.ColoredPrintln("Status:", state.Data.Status)
		utils.ColoredPrintln("Balance:", state.Data.Balance)
		utils.ColoredPrintln("Effective Balance:", state.Data.Validator.EffectiveBalance)
		utils.ColoredPrintln("Activation Epoch:", state.Data.Validator.ActivationEpoch)
		utils.ColoredPrintln("Activation Eligibility Epoch:", state.Data.Validator.ActivationEligibilityEpoch)
		utils.ColoredPrintln("Exit Epoch:", state.Data.Validator.ExitEpoch)
		utils.ColoredPrintln("Withdrawable Epoch:", state.Data.Validator.WithdrawableEpoch)
		utils.ColoredPrintln("Withdrawal Credentials", state.Data.Validator.WithdrawalCredentials)
		utils.ColoredPrintln("Is Slashed:", state.Data.Validator.Slashed)
	}

	return nil
}

func init() {
	validatorDescribeCmd.AddCommand(validatorDescribeByKeyCmd)
	validatorDescribeByKeyCmd.Flags().StringP(CommandOptionExecutionApi, CommandOptionExecutionApiShort, "", "execution api endpoint")
	validatorDescribeByKeyCmd.Flags().StringP(CommandOptionConsensusApi, CommandOptionConsensusApiShort, "", "consensus api endpoint")
	validatorDescribeByKeyCmd.Flags().StringP("key", "k", "", "validator key to be described - keep empty to describe your validators defined in keystore")
	validatorDescribeByKeyCmd.Flags().StringP(CommandOptionDepositAddress, CommandOptionDepositAddressShort, "", "the address of the deposit contract")

	validatorDescribeByKeyCmd.MarkFlagRequired("key")
}
