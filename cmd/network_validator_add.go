/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/lukso-network/lukso-cli/src/network"
	"github.com/lukso-network/lukso-cli/src/utils"
	"github.com/lukso-network/lukso-cli/src/wallet"
	"github.com/spf13/cobra"
)

// validatorAddCmd represents the validator command
var validatorAddCmd = &cobra.Command{
	Use:     "add",
	Short:   "Generates additional validator keys",
	Example: "lukso network validator add",
	Run: func(cmd *cobra.Command, args []string) {
		passwordFile, _ := cmd.Flags().GetString("passwordFile")

		p, err := wallet.ReadPasswordFile(passwordFile)
		if err != nil {
			utils.PrintColoredError(err.Error())
		}
		network.
			NewAddValidatorProcess(network.MustGetNodeConfig(), p).
			Add()
	},
}

func init() {
	validatorCmd.AddCommand(validatorAddCmd)
	// TODO: change passwordFile to password everywhere
	validatorAddCmd.Flags().StringP("passwordFile", "p", "", "the location of the password file")
	validatorAddCmd.MarkFlagRequired("passwordFile")
}
