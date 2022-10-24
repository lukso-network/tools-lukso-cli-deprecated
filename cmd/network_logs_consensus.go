/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/lukso-network/lukso-cli/src/network"
	"github.com/spf13/cobra"
)

// consensusCmd represents the consensus command
var consensusCmd = &cobra.Command{
	Use:     "consensus",
	Short:   "Show logs for consensus client",
	Long:    `Returns logs for the Prysm Beacon client, where consensus engine is running`,
	Example: "lukso network logs consensus --tail 30 -f",
	Run: func(cmd *cobra.Command, args []string) {
		network.ReadLog("prysm_beacon", tail, follow)
	},
}

func init() {
	logsCmd.AddCommand(consensusCmd)
}
