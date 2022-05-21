/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/lukso-network/lukso-cli/src/network"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// restartCmd represents the restart command
var restartCmd = &cobra.Command{
	Use:     "restart",
	Short:   "Restart all running docker containers",
	Example: "lukso network restart",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("stopping running containers")
		err := network.DownDockerServices()
		if err != nil {
			cobra.CompErrorln(err.Error())
			os.Exit(1)
		}
		// take a small break between stop and start
		time.Sleep(2 * time.Second)
		fmt.Println("starting docker containers")
		err = network.StartArchNode()
		if err != nil {
			cobra.CompErrorln(err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	networkCmd.AddCommand(restartCmd)
}
