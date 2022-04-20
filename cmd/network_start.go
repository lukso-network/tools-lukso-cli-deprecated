/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/lukso-network/lukso-cli/src/network"
	"github.com/spf13/cobra"
	"os"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Spin up consensus, execution and eth2-stats docker container",
	Long: `start command spins up LUKSO node using .env and docker-compose file. It spins up
consensus engine, execution engine and eth2-stats containers.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := network.StartArchNode()
		if err != nil {
			cobra.CompErrorln(err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	networkCmd.AddCommand(startCmd)
}