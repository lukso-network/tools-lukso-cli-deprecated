/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/lukso-network/lukso-cli/src/network"
	"github.com/spf13/cobra"
	"os"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:     "stop",
	Short:   "Stops all running clients",
	Long:    `Stops consensus client, execution client, validator client and eth2-stats client.`,
	Example: "lukso network stop",
	Run: func(cmd *cobra.Command, args []string) {
		err := network.DownDockerServices()
		if err != nil {
			cobra.CompErrorln(err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	networkCmd.AddCommand(stopCmd)
}
