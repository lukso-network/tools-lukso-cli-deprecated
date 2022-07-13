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
	Short:   "Creates a recovery file",
	Long:    `Creates a recovery file that can be used to recreate the node somewhere else`,
	Example: "lukso network validator backup",
	Run: func(cmd *cobra.Command, args []string) {
		nodeConf := network.MustGetNodeConfig()
		credentials := nodeConf.ValidatorCredentials
		wallet := nodeConf.TransactionWallet
		if credentials == nil || credentials.IsEmpty() {
			utils.PrintColoredError(network.ErrMsgValidatorSecretNotPresent)
			return
		}

		err := credentials.CreateNodeRecovery().Save()
		if err != nil {
			utils.PrintColoredErrorWithReason("couldn't save validator credentials in recovery file", err)
			return
		}

		err = wallet.CreateNodeRecovery().Append()
		//if err != nil {
		//	utils.PrintColoredErrorWithReason("couldn't save wallet info in recovery file", err)
		//	return
		//}
		fmt.Println("A file ./node_recovery.json was created. Store this in a save place.")
		fmt.Println("You can recover your keystore with")
		utils.Coloredln("   lukso network validator recover --path [PATH_TO_FILE]")
		fmt.Println("Make sure to NEVER run 2 nodes with the same keystore as you could be prone to slashing.")
	},
}

func init() {
	validatorCmd.AddCommand(validatorBackupCmd)
}
