/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/lukso-network/lukso-cli/src"
	"github.com/lukso-network/lukso-cli/src/network"
	"github.com/spf13/cobra"
)

// validatorDescribeCmd represents the describe command
var validatorDescribeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Show detailed status of the validators",
	Long:  `It shows validator count, addresses and transaction status.`,
	Run: func(cmd *cobra.Command, args []string) {
		nodeConf, err := network.GetLoadedNodeConfigs()
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}
		valSecrets := nodeConf.GetValSecrets()
		if valSecrets == nil {
			cobra.CompErrorln(src.ErrMsgValidatorSecretNotPresent)
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
