/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os/exec"

	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stops running docker containers",
	Long: `This command stops consensus engine, execution engine, validator client and eth2-stats.
It uses docker-compose file to stop these containers`,
	Run: func(cmd *cobra.Command, args []string) {
		command := exec.Command("sudo", "docker-compose", "down")
		if err := command.Run(); err != nil {
			cobra.CompErrorln(err.Error())
		}
	},
}

func init() {
	networkCmd.AddCommand(stopCmd)
}
