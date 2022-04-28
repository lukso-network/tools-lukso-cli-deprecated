/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/lukso-network/lukso/src/network"

	"github.com/spf13/cobra"
)

// validatorLogCmd represents the validator command
var validatorLogCmd = &cobra.Command{
	Use:   "validator",
	Short: "Show logs for validator client",
	Long:  `This command shows log for prysm-validator container where validator client is running`,
	Run: func(cmd *cobra.Command, args []string) {
		network.ReadLog("prysm_validator", tail, follow)
	},
}

func init() {
	logCmd.AddCommand(validatorLogCmd)
}
