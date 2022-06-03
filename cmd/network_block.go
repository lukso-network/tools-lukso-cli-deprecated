/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lukso-network/lukso-cli/src/network"
	"github.com/lukso-network/lukso-cli/src/utils"
	"github.com/spf13/cobra"
	"math/big"
)

// blockCmd represents the block command
var blockCmd = &cobra.Command{
	Use:     "block",
	Short:   "returns a block at number",
	Long:    `This command will return the execution block at the given position`,
	Example: "lukso network block -n 100",
	Run: func(cmd *cobra.Command, args []string) {
		number, _ := cmd.Flags().GetInt64("number")
		baseUrl, _ := cmd.Flags().GetString(CommandOptionExecutionApi)

		if baseUrl == "" {
			baseUrl = network.MustGetNodeConfig().ApiEndpoints.ExecutionApi
		}

		client, err := ethclient.Dial(baseUrl)
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}

		blockNumber := big.NewInt(number)
		if number == -2 {
			blockNumber = nil
			fmt.Println("Fetching latest block")
		} else {
			fmt.Println("Fetching block", number)
		}

		block, err := client.BlockByNumber(context.Background(), blockNumber)
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}

		utils.ColoredPrintln("Block:", block.Number())
		utils.ColoredPrintln("Hash:", block.Hash())
		utils.ColoredPrintln("#Transactions:", block.Transactions().Len())
	},
}

func init() {
	networkCmd.AddCommand(blockCmd)
	blockCmd.Flags().Int64P("number", "n", -2, "block number of geth block")
	blockCmd.Flags().StringP(CommandOptionExecutionApi, CommandOptionExecutionApiShort, "", "executionApi endpoint")
}
