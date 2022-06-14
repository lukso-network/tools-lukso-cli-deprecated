/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/lukso-network/lukso-cli/src/network"
	"github.com/spf13/cobra"
)

// validatorAddCmd represents the validator command
var validatorAddCmd = &cobra.Command{
	Use:     "add",
	Short:   "This adds validators to your existing setup",
	Example: "lukso network validator add",
	Run: func(cmd *cobra.Command, args []string) {
		network.
			NewAddValidatorProcess(network.MustGetNodeConfig()).
			Add()
	},
}

func init() {
	validatorCmd.AddCommand(validatorAddCmd)
}
