/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/lukso-network/lukso/src/network"
	"github.com/spf13/cobra"
)

// executionCmd represents the execution command
var executionCmd = &cobra.Command{
	Use:     "execution",
	Short:   "Show logs for execution engine",
	Long:    `This command shows log for geth container where execution engine is running`,
	Example: "lukso-cli network log execution --tail 30 -f",
	Run: func(cmd *cobra.Command, args []string) {
		network.ReadLog("lukso-geth", tail, follow)
	},
}

func init() {
	logCmd.AddCommand(executionCmd)
}
