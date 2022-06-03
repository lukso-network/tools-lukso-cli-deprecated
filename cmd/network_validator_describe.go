/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lukso-network/lukso-cli/src/network"
	"github.com/lukso-network/lukso-cli/src/utils"
	"github.com/spf13/cobra"
)

// validatorDescribeCmd represents the describe command
var validatorDescribeCmd = &cobra.Command{
	Use:     "describe",
	Short:   "Show detailed status of the validators",
	Long:    `It shows validator count, addresses and transaction status.`,
	Example: "lukso network validator describe",
	Run: func(cmd *cobra.Command, args []string) {
		nodeConf := network.MustGetNodeConfig()
		vc := nodeConf.ValidatorCredentials

		// should never happen
		if vc == nil {
			cobra.CompErrorln("couldn't find contract details")
			return
		}

		events, err := network.NewDepositEvents(nodeConf.DepositDetails.ContractAddress, nodeConf.ApiEndpoints.ExecutionApi)
		if err != nil {
			cobra.CompErrorln(fmt.Sprintf("couldn't load deposit data from contract, reason: %s", err.Error()))
			return
		}

		valSecrets := nodeConf.GetCredentials()
		if valSecrets == nil {
			cobra.CompErrorln(network.ErrMsgValidatorSecretNotPresent)
			return
		}

		if nodeConf.TransactionWallet == nil {
			cobra.CompErrorln("Transaction wallet does not exist, did you forget to run lukso network validator setup?")
			return
		}

		utils.ColoredPrintln("Transaction wallet address:", nodeConf.TransactionWallet.PublicKey)
		client, err := ethclient.Dial(nodeConf.ApiEndpoints.ExecutionApi)
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}

		balance, err := client.BalanceAt(context.Background(), common.HexToAddress(nodeConf.TransactionWallet.PublicKey), nil)
		if err != nil {
			cobra.CompErrorln(err.Error())
		} else {
			utils.ColoredPrintln("Transaction wallet balance:", utils.WeiToString(balance, true))
		}

		depositData, err := network.ParseDepositDataFromFile(nodeConf.DepositDetails.DepositFileLocation)
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}

		utils.ColoredPrintln("Number of validators:", len(depositData))

		pubKeys := make([]string, len(depositData))
		for k, d := range depositData {
			pubKeys[k] = d.PubKey
		}
		fmt.Println("........................................................................................................................................................................")
		err = network.DescribeValidatorKey(pubKeys, nodeConf.DepositDetails.ContractAddress, nodeConf.ApiEndpoints.ExecutionApi, nodeConf.ApiEndpoints.ConsensusApi, &events)
		if err != nil {
			cobra.CompErrorln(err.Error())
		}
	},
}

func init() {
	validatorCmd.AddCommand(validatorDescribeCmd)
}
