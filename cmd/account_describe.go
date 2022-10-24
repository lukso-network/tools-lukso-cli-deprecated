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

// accountDecryptCmd represents the create command
var accountDecryptCmd = &cobra.Command{
	Use:     "decrypt",
	Short:   "Decrypts and describes an account with taking in a password file.",
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
	accountCmd.AddCommand(accountDecryptCmd)
	accountDecryptCmd.Flags().StringP("walletFile", "w", "", "the location of the wallet file")
	accountDecryptCmd.Flags().StringP("passwordFile", "p", "", "the location of the password file")

	accountDecryptCmd.MarkFlagRequired("walletFile")
	accountDecryptCmd.MarkFlagRequired("passwordFile")
}
