/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/lukso-network/lukso-cli/api/gethrpc"
	"github.com/lukso-network/lukso-cli/src/network"
	"github.com/lukso-network/lukso-cli/src/utils"
	"github.com/spf13/cobra"
)

// balanceCmd represents the balance command
var balanceCmd = &cobra.Command{
	Use:     "balance",
	Short:   "returns the balance of a given address",
	Long:    `This command will return the balance of a given address`,
	Example: "lukso network balance -a 0x....",
	Run: func(cmd *cobra.Command, args []string) {
		address, _ := cmd.Flags().GetString("address")
		if address == "" {
			cobra.CompError("address must be given")
			return
		}
		nodeConf, err := network.GetLoadedNodeConfigs()
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}

		client := gethrpc.NewRPCClient(nodeConf.ApiEndpoints.ExecutionApi)

		fmt.Println("Calling ", nodeConf.ApiEndpoints.ExecutionApi)
		balance, err := client.GetBalance(address)

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
