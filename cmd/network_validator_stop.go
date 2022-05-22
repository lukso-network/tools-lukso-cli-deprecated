/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/lukso-network/lukso-cli/src/network"
	"github.com/spf13/cobra"
)

// validatorStopCmd represents the validator command
var validatorStopCmd = &cobra.Command{
	Use:     "stop",
	Short:   "This subcommand stops the validator client",
	Example: "lukso network validator start",
	RunE: func(cmd *cobra.Command, args []string) error {
		return network.StopValidatorNode()
	},
}

func init() {
	validatorCmd.AddCommand(validatorStopCmd)
}
