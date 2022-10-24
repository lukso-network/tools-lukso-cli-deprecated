/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/lukso-network/lukso-cli/src/network"
	"github.com/spf13/cobra"
)

// executionCmd represents the execution command
var executionCmd = &cobra.Command{
	Use:     "execution",
	Short:   "Show logs for execution client",
	Long:    `Returns logs for geth client, where execution engine is running`,
	Example: "lukso network logs execution --tail 30 -f",
	Run: func(cmd *cobra.Command, args []string) {
		network.ReadLog("lukso-geth", tail, follow)
	},
}

func init() {
	logsCmd.AddCommand(executionCmd)
}
