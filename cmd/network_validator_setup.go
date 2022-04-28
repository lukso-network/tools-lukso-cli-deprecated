/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"github.com/lukso-network/lukso/src"
	"github.com/lukso-network/lukso/src/network"
	"github.com/manifoldco/promptui"
	"strconv"

	"github.com/spf13/cobra"
)

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Prepare a validator to join the network",
	Long: `This command prepares wallet, deposit_data and creates a secret.yaml file. These files are necessary to
activate validators`,
	Run: func(cmd *cobra.Command, args []string) {
		prompt := promptui.Prompt{
			Label: "Keystore password",
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
			cobra.CompErrorln(err.Error())
			return
		}
		prompt = promptui.Prompt{
			Label: "Number of Validators",
		}
		validatorNumber, err := prompt.Run()
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}
		nodeConf, err := network.GetLoadedNodeConfigs()
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}
		valSecrets := nodeConf.GetValSecrets()
		if valSecrets == nil {
			cobra.CompErrorln(src.ErrMsgValidatorSecretNotPresent)
			return
		}
		err = valSecrets.GenerateMnemonic()
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}
		numOfVal, err := strconv.Atoi(validatorNumber)
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}
		err = valSecrets.GenerateDepositData(numOfVal)
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}
		err = valSecrets.GenerateWallet(numOfVal, password)
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}
		err = nodeConf.WriteOrUpdateNodeConfig()
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}
	},
}

func init() {
	validatorCmd.AddCommand(setupCmd)
}
