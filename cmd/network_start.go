/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/lukso-network/lukso-cli/src/network"
	"github.com/lukso-network/lukso-cli/src/utils"
	"github.com/spf13/cobra"
	"os"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:     "start",
	Short:   "Starts consensus, execution and eth2-stats clients",
	Long:    `Starts LUKSO node with the consensus engine, execution engine and eth2-stats clients.`,
	Example: "lukso network start",
	Run: func(cmd *cobra.Command, args []string) {
		updateEnv()
		err := network.StartArchNode()
		if err != nil {
			cobra.CompErrorln(err.Error())
			os.Exit(1)
		}
		utils.ColoredPrintln("LUKSO Network", "successfully started")
	},
}

func init() {
	networkCmd.AddCommand(startCmd)
}

func updateEnv() {
	fmt.Println("update env")
	err := network.GenerateEnvFile()
	if err != nil {
		cobra.CompErrorln(err.Error())
		return
	}
}
