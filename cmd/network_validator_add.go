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
	Short:   "This adds validators to your existing setup",
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
	validatorAddCmd.Flags().StringP("passwordFile", "p", "", "the location of the password file")
	validatorAddCmd.MarkFlagRequired("passwordFile")
}
