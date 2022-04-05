/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/lukso-network/lukso-cli/src"
	"github.com/lukso-network/lukso-cli/src/network"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os/exec"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Spin up consensus, execution and eth2-stats docker container",
	Long: `start command spins up LUKSO node using .env and docker-compose file. It spins up
consensus engine, execution engine and eth2-stats containers.`,
	Run: func(cmd *cobra.Command, args []string) {
		configDirName := viper.GetString(src.ViperKeyConfigDir)
		if configDirName != "" && network.FileExists(configDirName) {
			fmt.Println("You may need to provide super user (sudo) password to run docker (if needed)")
			command := exec.Command("sudo", "docker-compose", "up", "-d", "init-geth", "geth", "prysm_beacon", "eth2stats-client")
			if err := command.Run(); err != nil {
				cobra.CompErrorln(fmt.Errorf("found error while running docker. Make sure your docker is running. %s", err).Error())
			}
		} else {
			fmt.Println("Config files are not present. Can't start docker containers")
		}
	},
}

func init() {
	networkCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
