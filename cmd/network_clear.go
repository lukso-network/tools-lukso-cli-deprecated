/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/lukso-network/lukso/src/network"
	"github.com/spf13/cobra"
	"os"
)

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Remove logs and data directory for all clients",
	Long: `This command is responsible to remove data directory and logs for all the running clients (e.g. consensus
engine, execution engine and validator client)`,
	Run: func(cmd *cobra.Command, args []string) {
		err := network.Clear()
		if err != nil {
			cobra.CompErrorln(err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	networkCmd.AddCommand(clearCmd)
}
