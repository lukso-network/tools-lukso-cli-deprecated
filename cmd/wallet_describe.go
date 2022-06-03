/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/lukso-network/lukso-cli/src/utils"
	"github.com/lukso-network/lukso-cli/src/wallet"
	"github.com/spf13/cobra"
)

// walletDescribeCmd represents the create command
var walletDescribeCmd = &cobra.Command{
	Use:     "describe",
	Short:   "Describes a wallet by reading the wallet and  password file ",
	Long:    "This command will describe a wallet and password file in a target directory.",
	Example: "lukso wallet create -p [PASSWORD] -d [TARGET_DIRECTORY] -l [LABEL]",
	Run: func(cmd *cobra.Command, args []string) {
		walletFile, _ := cmd.Flags().GetString("walletFile")
		passwordFile, _ := cmd.Flags().GetString("passwordFile")

		p, err := wallet.ReadPasswordFile(passwordFile)
		if err != nil {
			utils.PrintColoredError(err.Error())
		}
		key, err := wallet.KeyFromWalletAndPasswordFile(walletFile, p)
		if err != nil {
			utils.PrintColoredError(err.Error())
		}

		fmt.Println("Public Key: ", wallet.PublicKeyFromKey(key))
		fmt.Println("Private Key: ", wallet.PrivateKeyFromKey(key))
	},
}

func init() {
	walletCmd.AddCommand(walletDescribeCmd)
	walletDescribeCmd.Flags().StringP("walletFile", "w", "", "the location of the wallet file")
	walletDescribeCmd.Flags().StringP("passwordFile", "p", "", "the location of the password file")

	walletDescribeCmd.MarkFlagRequired("walletFile")
	walletDescribeCmd.MarkFlagRequired("passwordFile")
}
