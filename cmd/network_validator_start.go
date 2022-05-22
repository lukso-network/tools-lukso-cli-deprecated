/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/lukso-network/lukso-cli/src/network"
	"github.com/spf13/cobra"
)

// validatorStartCmd represents the validator command
var validatorStartAltCmd = &cobra.Command{
	Use:     "start",
	Short:   "This subcommand starts a validator client",
	Example: "lukso network validator start",
	RunE: func(cmd *cobra.Command, args []string) error {
		return network.StartValidatorNode()
	},
}

func init() {
	validatorCmd.AddCommand(validatorStartAltCmd)
}
