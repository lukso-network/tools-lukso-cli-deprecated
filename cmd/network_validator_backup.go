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

// validatorBackupCmd represents the describe command
var validatorBackupCmd = &cobra.Command{
	Use:     "backup",
	Short:   "Creates a backup file containing all validator keys",
	Example: "lukso network validator backup",
	Run: func(cmd *cobra.Command, args []string) {
		nodeConf := network.MustGetNodeConfig()
		credentials := nodeConf.ValidatorCredentials
		if credentials == nil || credentials.IsEmpty() {
			utils.PrintColoredError(network.ErrMsgValidatorSecretNotPresent)
			return
		}
		wallet := nodeConf.TransactionWallet
		if wallet == nil || wallet.IsEmpty() {
			utils.PrintColoredError(network.ErrMsgTransactionWalletNotPresent)
			return
		}

		err := nodeConf.CreateNodeRecovery().Save()
		if err != nil {
			utils.PrintColoredErrorWithReason("couldn't save validator credentials or transaction wallet in recovery file", err)
			return
		}

		fmt.Println("A file ./node_recovery.json was created. Store this in a save place.")
		fmt.Println("You can recover your keystore with")
		utils.Coloredln("   lukso network validator recover --path [PATH_TO_FILE]")
		fmt.Println("Make sure to NEVER run 2 nodes with the same keystore as you could be prone to slashing.")
	},
}

func init() {
	validatorCmd.AddCommand(validatorBackupCmd)
}
