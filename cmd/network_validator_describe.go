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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
		if len(depositData) > 0 {
			fmt.Println("Public keys for validators")
		}
		for _, valData := range depositData {
			fmt.Println(valData.PubKey)
		}
	},
}

func init() {
	validatorCmd.AddCommand(validatorDescribeCmd)
}
