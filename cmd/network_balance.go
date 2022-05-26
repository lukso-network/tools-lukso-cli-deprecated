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
)

// balanceCmd represents the balance command
var balanceCmd = &cobra.Command{
	Use:     "balance",
	Short:   "returns the balance of a given address",
	Long:    `This command will return the balance of a given address based on the network given.`,
	Example: "lukso network balance -a 0x....",
	Run: func(cmd *cobra.Command, args []string) {
		address, _ := cmd.Flags().GetString("address")
		if address == "" {
			cobra.CompError("address must be given")
			return
		}
		nodeConf := network.MustGetNodeConfig()

		client, err := ethclient.Dial(nodeConf.ApiEndpoints.ExecutionApi)
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
	networkCmd.AddCommand(balanceCmd)
	balanceCmd.Flags().StringP("address", "a", "", "ethereum address")
}
