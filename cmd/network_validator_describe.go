/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/lukso-network/lukso-cli/api/beaconapi"
	"github.com/lukso-network/lukso-cli/api/gethrpc"
	"github.com/lukso-network/lukso-cli/src/network"
	"github.com/lukso-network/lukso-cli/src/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strings"
)

// validatorDescribeCmd represents the describe command
var validatorDescribeCmd = &cobra.Command{
	Use:     "describe",
	Short:   "Show detailed status of the validators",
	Long:    `It shows validator count, addresses and transaction status.`,
	Example: "lukso network validator describe",
	Run: func(cmd *cobra.Command, args []string) {
		key, _ := cmd.Flags().GetString("key")
		baseUrl, _ := cmd.Flags().GetString("beaconapi")
		if baseUrl == "" {
			baseUrl = network.GetDefaultNodeConfigByOptionParam(viper.GetString(network.CommandOptionChain)).ApiEndpoints.ConsensusApi
		}
		epoch, _ := cmd.Flags().GetInt64("epoch")

		if key != "" {
			err := describe([]string{key}, baseUrl, epoch)
			if err != nil {
				cobra.CompErrorln(err.Error())
			}
			return
		}

		nodeConf, err := network.GetLoadedNodeConfigs()
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}

		baseUrl = nodeConf.ApiEndpoints.ConsensusApi

		valSecrets := nodeConf.GetValSecrets()
		if valSecrets == nil {
			cobra.CompErrorln(network.ErrMsgValidatorSecretNotPresent)
			return
		}

		utils.ColoredPrintln("Transaction wallet address:", valSecrets.Eth1Data.WalletAddress)
		client := gethrpc.NewRPCClient(nodeConf.ApiEndpoints.ExecutionApi)
		balance, err := client.GetBalance(valSecrets.Eth1Data.WalletAddress)
		if err != nil {
			cobra.CompErrorln(err.Error())
		}
		utils.ColoredPrintln("Transaction wallet balance:", balance)

		depositData, err := network.ParseDepositDataFromFile(valSecrets.Deposit.DepositFileLocation)
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}

		utils.ColoredPrintln("Number of validators:", len(depositData))

		pubKeys := make([]string, len(depositData))
		for k, d := range depositData {
			pubKeys[k] = d.PubKey
		}

		err = describe(pubKeys, baseUrl, epoch)
		if err != nil {
			cobra.CompErrorln(err.Error())
		}
	},
}

func maybeAddHexPrefix(address string) string {
	if strings.Contains(address, "0x") {
		return address
	}

	return fmt.Sprintf("0x%s", address)
}

func describe(pubKeys []string, baseUrl string, epoch int64) error {
	beaconClient := beaconapi.NewBeaconClient(baseUrl)
	for _, d := range pubKeys {
		pubKey := maybeAddHexPrefix(d)
		fmt.Println("....................................................................................")
		state, err := beaconClient.ValidatorState(pubKey, epoch)
		if err != nil {
			return err
		}
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
	validatorCmd.AddCommand(validatorDescribeCmd)
	validatorDescribeCmd.Flags().StringP("beaconapi", "b", "", "endpoint of beacon api")
	validatorDescribeCmd.Flags().Int64P("epoch", "e", -1, "epoch to be described - if left out it is the head epoch")
	validatorDescribeCmd.Flags().StringP("key", "k", "", "validator key to be described - keep empty to describe your validators defined in keystore")
}
