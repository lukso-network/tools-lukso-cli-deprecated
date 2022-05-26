/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/lukso-network/lukso-cli/src/network"
	"github.com/lukso-network/lukso-cli/src/utils"
	"github.com/spf13/cobra"
)

const CommandOptionGasPrice = "gasPrice"

// depositCmd represents the deposit command
var depositCmd = &cobra.Command{
	Use:   "deposit",
	Short: "Send Deposit transactions to activate validator",
	Long: `After preparing wallets and deposit data, this command prepares deposit transactions to the deposit contract
address. Remember it will need your wallet address and private keys. Thus it will deduct balance from your wallet.

This tool is necessary to activate new validators`,
	Example: "lukso network validator deposit",
	Run: func(cmd *cobra.Command, args []string) {
		nodeConf := network.MustGetNodeConfig()
		valSecrets := nodeConf.GetValSecrets()
		if valSecrets == nil {
			cobra.CompErrorln("no validator credentials exist")
			return
		}

		gasPrice, err := cmd.Flags().GetInt64(CommandOptionGasPrice)
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}
		totalDeposits, err := network.Deposit(valSecrets.Deposit.DepositFileLocation, valSecrets.Deposit.ContractAddress, valSecrets.Eth1Data.WalletPrivKey, nodeConf.ApiEndpoints.ExecutionApi, gasPrice)
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}
		fmt.Printf("You successfully deposited %d key(s)! Your keys need to be activated which takes around 8h. You can check the status by calling:\n", totalDeposits)
		fmt.Println(utils.ConsoleInBlue("        lukso network validator describe"))
	},
}

func init() {
	validatorCmd.AddCommand(depositCmd)

	depositCmd.Flags().Int64P(CommandOptionGasPrice, "g", 1000000, "set the gas price for transactions")
}
