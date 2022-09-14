/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lukso-network/lukso-cli/src/network"
	"github.com/lukso-network/lukso-cli/src/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// balanceCmd represents the balance command
var balanceCmd = &cobra.Command{
	Use:     "balance",
	Short:   "Returns the balance of a given address",
	Long:    `Returns the balance of a given address based on the network given.`,
	Example: "lukso network info balance -a 0x....",
	Run: func(cmd *cobra.Command, args []string) {
		address, _ := cmd.Flags().GetString("address")
		if address == "" {
			cobra.CompError("address must be given")
			return
		}
		// get node conf from --chain param or get default chain
		nodeConf := network.GetDefaultNodeConfigByOptionParam(viper.GetString(CommandOptionChain))
		executionApi, err := readExecutionApiEndpoint(cmd, nodeConf)
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}

		client, err := ethclient.Dial(executionApi)
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}

		balance, err := client.BalanceAt(context.Background(), common.HexToAddress(address), nil)
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}

		utils.ColoredPrintln("Balance:", balance)
	},
}

func init() {
	infoCmd.AddCommand(balanceCmd)
	balanceCmd.Flags().StringP("address", "a", "", "ethereum address")
	balanceCmd.Flags().StringP(CommandOptionExecutionApi, CommandOptionExecutionApiShort, "", "execution api endpoint")
}
