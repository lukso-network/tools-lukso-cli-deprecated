/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lukso-network/lukso-cli/src/network"
	"github.com/lukso-network/lukso-cli/src/utils"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"strings"
)

// validatorRecoverCmd represents the describe command
var validatorRecoverCmd = &cobra.Command{
	Use:     "recover",
	Short:   "Recovers a keystore from a recovery file",
	Long:    `Creates a recovery file that can be used to recreate the node somewhere els`,
	Example: "lukso network validator recover --path [PATH_TO_RECOVERY_FILE]",
	Run: func(cmd *cobra.Command, args []string) {
		nodeConf := network.MustGetNodeConfig()
		keystorePath := nodeConf.Keystore.Volume
		isKeystoreEmpty, err := utils.IsDirectoryEmpty(keystorePath)
		if err != nil {
			// if directory doesn't exist, ignore it
			if strings.Contains(err.Error(), "no such file or directory") {
			} else {
				utils.PrintColoredError(fmt.Sprintf("couldn't determine if keystore directory exists, reason: %v", err.Error()))
				return
			}
		}
		if !isKeystoreEmpty {
			utils.PrintColoredError("The keystore directory is not empty. In order to setup the validator you need an empty keystore directory. \nConsider setting up the node in a different location.\n")
			return
		}

		path, err := cmd.Flags().GetString(CommandOptionPath)
		if err != nil {
			utils.PrintColoredError(fmt.Sprintf("couldn't get flags, reason: %v", err.Error()))
			return
		}

		nr, err := network.LoadNodeRecovery(path)
		if err != nil {
			utils.PrintColoredError(fmt.Sprintf("couldn't load recover file, reason: %v", err.Error()))
			return
		}

		fmt.Println("Loaded the recovery file: ")
		data, _ := json.MarshalIndent(nr, "", "    ")
		fmt.Println(string(data))
		if !promptDoYouWantToContinue("Do you really want to recover the keystore?") {
			fmt.Println("Recovery  canceled...")
			return
		}

		// choose password
		prompt := promptui.Prompt{
			Label: "Choose A Password For Your Keystore",
			Validate: func(s string) error {
				if len(s) < 6 {
					return errors.New("password must have more than 6 characters")
				}
				return nil
			},
			Mask: '*',
		}
		password, err := prompt.Run()
		if err != nil {
			utils.PrintColoredError(fmt.Sprintf("couldn't read password, reason: %v", err.Error()))
			return
		}

		credentials := new(network.ValidatorCredentials).FromNodeRecovery(*nr)

		// generate deposit data
		err = credentials.GenerateDepositDataWithRange(nodeConf.DepositDetails, credentials.ValidatorRange())
		if err != nil {
			utils.PrintColoredError(fmt.Sprintf("couldn't create deposit data, reason: %v", err.Error()))
			return
		}

		// generate wallet
		err = credentials.GenerateKeystoreWithRange(credentials.ValidatorRange(), password)
		if err != nil {
			utils.PrintColoredError(fmt.Sprintf("couldn't create keystore, reason: %v", err.Error()))
			return
		}

		nodeConf.ValidatorCredentials = credentials

		transactionWallet := new(network.TransactionWallet).FromNodeRecovery(*nr)
		nodeConf.TransactionWallet = transactionWallet
		err = nodeConf.Save()
		if err != nil {
			utils.PrintColoredError(fmt.Sprintf("couldn't save nodeConf, reason: %v", err.Error()))
			return
		}

		fmt.Println("Successfully recovered your keystore!!!!!")
	},
}

func init() {
	validatorCmd.AddCommand(validatorRecoverCmd)

	validatorRecoverCmd.Flags().StringP(CommandOptionPath, CommandOptionPathShort, network.NodeRecoveryFileLocation, "path to recovery file")
}
