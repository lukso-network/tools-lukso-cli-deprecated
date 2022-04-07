/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"
	"github.com/lukso-network/lukso-cli/src/network"
	"os/exec"

	"github.com/spf13/cobra"
)

// validatorStartCmd represents the validator command
var validatorStartCmd = &cobra.Command{
	Use:   "validator",
	Short: "This subcommand starts a validator client",
	RunE: func(cmd *cobra.Command, args []string) error {
		configFilePath, err := network.GetConfigPath()
		if err != nil {
			cobra.CompError(err.Error())
			return err
		}
		keystorePath, err := network.GetKeyStorePath()
		if err != nil {
			cobra.CompError(err.Error())
			return err
		}
		if !network.FileExists(configFilePath) {
			return errors.New("config file path is invalid")
		}
		if !network.FileExists(keystorePath) {
			return errors.New("keystore path is invalid")
		}
		fmt.Println("spinning validator client. You may need super user (sudo) password")
		command := exec.Command("sudo", "docker-compose", "up", "-d", "prysm_validator")
		if err := command.Run(); err != nil {
			cobra.CompErrorln(fmt.Errorf("found error while running docker. Make sure your docker is running. %s", err).Error())
		}
		return nil
	},
}

func init() {
	startCmd.AddCommand(validatorStartCmd)
}
