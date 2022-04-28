/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/lukso-network/lukso/src/network"
	"github.com/spf13/cobra"
)

// validatorStartCmd represents the validator command
var validatorStartCmd = &cobra.Command{
	Use:   "validator",
	Short: "This subcommand starts a validator client",
	RunE: func(cmd *cobra.Command, args []string) error {
		return network.StartValidatorNode()
	},
}

func init() {
	startCmd.AddCommand(validatorStartCmd)
}
