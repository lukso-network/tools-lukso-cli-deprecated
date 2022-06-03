/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"
	"github.com/lukso-network/lukso-cli/src/network"
	"github.com/lukso-network/lukso-cli/src/utils"
	"github.com/lukso-network/lukso-cli/src/wallet"
	"github.com/manifoldco/promptui"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Prepare a validator to join the network",
	Long: `This command prepares wallet, deposit_data and creates a secret.yaml file. These files are necessary to
activate validators`,
	Example: "lukso network validator setup",
	Run: func(cmd *cobra.Command, args []string) {
		// Checks
		// load node config
		nodeConf := network.MustGetNodeConfig()

		// check if keystore is empty
		keystorePath := nodeConf.Keystore.Volume
		isKeystoreEmpty, err := utils.IsDirectoryEmpty(keystorePath)
		if err != nil {
			// if directory doesn't exist, ignore it
			if strings.Contains(err.Error(), "no such file or directory") {
			} else {
				cobra.CompErrorln(err.Error())
				return
			}
		}
		if !isKeystoreEmpty {
			cobra.CompError("The keystore directory is not empty. In order to setup the validator you need an empty keystore directory. \nConsider setting up the node in a different location.\n")
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
			cobra.CompErrorln(err.Error())
			return
		}

		// set number of validators
		prompt = promptui.Prompt{
			Label: "How Many Validators Do You Want To Run",
		}

		numOfValString, err := prompt.Run()
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}
		numOfVal, err := strconv.Atoi(numOfValString)
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}

		// create secrets
		valSecrets := nodeConf.CreateCredentials()
		// generate mnemonic
		err = valSecrets.GenerateMnemonic()
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}

		// generate deposit data
		err = valSecrets.GenerateDepositData(nodeConf.DepositDetails, numOfVal)
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}

		// generate wallet
		err = valSecrets.GenerateWallet(numOfVal, password)
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}

		walletInfo, err := wallet.CreateWallet("transaction_wallet", "", "transaction_wallet")
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}

		nodeConf.TransactionWallet = &network.TransactionWallet{
			PublicKey:  walletInfo.PubKey,
			PrivateKey: walletInfo.PrivKey,
		}

		// push node config
		err = nodeConf.WriteOrUpdateNodeConfig()
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}
		fmt.Println("Validator wallet was successfully created. Type")
		fmt.Println(utils.ConsoleInBlue("        lukso network validator describe"))
		fmt.Println("to see data related to the validator setup. ")
		fmt.Println("A transaction wallet was created to pay for the deposit transaction. ")
		fmt.Println("The transaction wallet needs at least [staking amount] + [gas costs] LyX before you can create a deposit transaction!")

	},
}

func init() {
	validatorCmd.AddCommand(setupCmd)
}
