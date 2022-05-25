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
	"os"
)

// blockCmd represents the block command
var blockCmd = &cobra.Command{
	Use:     "block",
	Short:   "returns a block at number",
	Long:    `This command will return the execution block at the given position`,
	Example: "lukso network block -n 100",
	Run: func(cmd *cobra.Command, args []string) {
		number, _ := cmd.Flags().GetInt64("number")

		nodeConf := network.MustGetNodeConfig()
		client := gethrpc.NewRPCClient(nodeConf.ApiEndpoints.ExecutionApi)

		fmt.Println("Calling ", nodeConf.ApiEndpoints.ExecutionApi)
		block, err := client.GetBlock(number)

		if err != nil {
			cobra.CompErrorln(err.Error())
			os.Exit(1)
		}

		utils.ColoredPrintln("Block:", block.Number)
		utils.ColoredPrintln("Hash:", block.Hash)
		utils.ColoredPrintln("#Transactions:", block.NumberOfTransactions)
	},
}

func init() {
	networkCmd.AddCommand(blockCmd)
	blockCmd.Flags().Int64P("number", "n", 0, "block number of geth block")
}
