/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lukso-network/lukso-cli/api/beaconapi"
	"github.com/lukso-network/lukso-cli/src/network"
	"github.com/lukso-network/lukso-cli/src/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
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
		epoch, _ := cmd.Flags().GetInt64("epoch")

		nodeConf := network.LoadNodeConfOrDefault(network.GetChainByString(viper.GetString(network.CommandOptionChain)))
		baseUrl := nodeConf.ApiEndpoints.ConsensusApi
		vc := nodeConf.ValidatorCredentials

		// should never happen
		if vc == nil {
			cobra.CompErrorln("couldn't find contract details")
			return
		}

		events, err := network.NewDepositEvents(vc.Deposit.ContractAddress, nodeConf.ApiEndpoints.ExecutionApi)
		if err != nil {
			cobra.CompErrorln(fmt.Sprintf("couldn't load deposit data from contract, reason: %s", err.Error()))
			return
		}

		// key was provided in command option
		if key != "" {
			err := describe(&events, []string{key}, baseUrl, epoch)
			if err != nil {
				cobra.CompErrorln(err.Error())
			}
			return
		}

		valSecrets := nodeConf.GetValSecrets()
		if valSecrets == nil {
			cobra.CompErrorln(network.ErrMsgValidatorSecretNotPresent)
			return
		}

		utils.ColoredPrintln("Transaction wallet address:", valSecrets.Eth1Data.WalletAddress)
		client, err := ethclient.Dial(nodeConf.ApiEndpoints.ExecutionApi)
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}

		balance, err := client.BalanceAt(context.Background(), common.HexToAddress(valSecrets.Eth1Data.WalletAddress), nil)
		if err != nil {
			cobra.CompErrorln(err.Error())
		} else {
			utils.ColoredPrintln("Transaction wallet balance:", utils.WeiToString(balance, true))
		}

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

		err = describe(&events, pubKeys, baseUrl, epoch)
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

func describe(events *network.DepositEvents, pubKeys []string, baseUrl string, epoch int64) error {
	beaconClient := beaconapi.NewBeaconClient(baseUrl)
	for _, d := range pubKeys {
		pubKey := maybeAddHexPrefix(d)
		amount := events.Amount(common.HexToAddress(pubKey))
		if amount == 0 {
			fmt.Printf("Validator %s has no deposits yet...\n", pubKey)
			// nothing to do anymore, validator doesn't have a deposit yet
			continue
		}
		utils.ColoredPrintln("StakeAmount", amount)

		fmt.Println("....................................................................................")
		state, status, err := beaconClient.ValidatorState(pubKey, epoch)
		if status == http.StatusNotFound {
			fmt.Printf("Validator %s is pending...\n", pubKey)
			continue
		}
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
	validatorDescribeCmd.Flags().Int64P("epoch", "e", -1, "epoch to be described - if left out it is the head epoch")
	validatorDescribeCmd.Flags().StringP("key", "k", "", "validator key to be described - keep empty to describe your validators defined in keystore")
}
