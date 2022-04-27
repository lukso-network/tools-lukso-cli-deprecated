/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/lukso-network/lukso-cli/src/network"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// depositCmd represents the deposit command
var depositCmd = &cobra.Command{
	Use:   "deposit",
	Short: "Send Deposit transactions to activate validator",
	Long: `After preparing wallets and deposit data, this command prepares deposit transactions to the deposit contract
address. Remember it will need your wallet address and private keys. Thus it will deduct balance from your wallet.

This tool is necessary to activate new validators`,
	Run: func(cmd *cobra.Command, args []string) {
		valSecrets, err := network.GetValSecrets(viper.GetString("chainId"))
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}
		err = valSecrets.SendDepositTxn()
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}
		err = valSecrets.WriteToFile("./secrets.yaml")
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}
	},
}

func init() {
	validatorCmd.AddCommand(depositCmd)
}
