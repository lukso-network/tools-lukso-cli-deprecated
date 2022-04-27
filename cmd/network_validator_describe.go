/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/lukso-network/lukso-cli/src/network"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// validatorDescribeCmd represents the describe command
var validatorDescribeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Show detailed status of the validators",
	Long:  `It shows validator count, addresses and transaction status.`,
	Run: func(cmd *cobra.Command, args []string) {
		valSecrets, err := network.GetValSecrets(viper.GetString("chainId"))
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}
		depositData, err := network.ParseDepositDataFromFile(valSecrets.Deposit.DepositFileLocation)
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}
		fmt.Println("Transaction wallet address:", valSecrets.Eth1Data.WalletAddress)
		fmt.Println("Number of validators:", len(depositData))
		err = valSecrets.GetTxStatus()
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}
	},
}

func init() {
	validatorCmd.AddCommand(validatorDescribeCmd)
}
