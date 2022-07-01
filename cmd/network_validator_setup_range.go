/*
Copyright © 2022 The LUKSO authors

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"errors"
	"fmt"
	"github.com/lukso-network/lukso-cli/src/network"
	"github.com/lukso-network/lukso-cli/src/utils"
	"github.com/lukso-network/lukso-cli/src/wallet"
	"github.com/manifoldco/promptui"
	"strings"

	"github.com/spf13/cobra"
)

// setupCmd represents the setup command
var setupRangeCmd = &cobra.Command{
	Use:   "range",
	Short: "Prepare a keystore to join the network",
	Long: `This command prepares wallet, deposit_data and creates a secret.yaml file. These files are necessary to
activate validators. The command gives greater control than "lukso network validator setup" over creating a keystore by describing the position of the keys derived from the mnemonic.`,
	Example: "lukso network validator setup range --from 0 --to 10",
	Run: func(cmd *cobra.Command, args []string) {
		vRange, err := readRangeFromCommand(cmd)
		if err != nil {
			utils.PrintColoredError(err.Error())
			return
		}
		nodeConf := network.MustGetNodeConfig()

		password, err := cmd.Flags().GetString(CommandOptionPassword)
		if err != nil {
			utils.PrintColoredError(err.Error())
			return
		}

		noPrompt, err := cmd.Flags().GetBool(CommandOptionNoPrompt)
		if err != nil {
			utils.PrintColoredError(err.Error())
			return
		}

		existingMnemonic, err := cmd.Flags().GetString(CommandOptionExistingMnemonic)
		if err != nil {
			utils.PrintColoredError(err.Error())
			return
		}

		if noPrompt && password == "" {
			utils.PrintColoredError("need to provide a password with --password if no prompt")
			return
		}

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

		// preparing credentials
		fmt.Println("Creating keystore")
		// create secrets
		if !nodeConf.HasMnemonic() {
			fmt.Println("No mnemonic is present -> need to create one")
			credentials := nodeConf.CreateCredentials()
			// generate mnemonic
			if noPrompt {
				if existingMnemonic != "" {
					credentials.ValidatorMnemonic = existingMnemonic
					credentials.WithdrawalMnemonic = existingMnemonic
				} else {
					err = credentials.GenerateMnemonicWithoutPrompt()
				}
			} else {
				err = credentials.GenerateMnemonic()
			}

			if err != nil {
				utils.PrintColoredErrorWithReason("couldn't create mnemonic", err)
				return
			}
		}

		nodeConf.ValidatorCredentials.ValidatorIndexTo = vRange.To
		nodeConf.ValidatorCredentials.ValidatorIndexFrom = vRange.From
		if err = nodeConf.Save(); err != nil {
			cobra.CompErrorln(err.Error())
			return
		}
		// _ prepared the credentials

		// Everything is alright -> create the keystore
		fmt.Println("Creating keystore with the following configuration")
		nodeConf.ValidatorCredentials.Print()
		if !noPrompt {
			if !promptDoYouWantToContinue("Do you want to continue?") {
				fmt.Println("Creation canceled")
				return
			}
		}

		if !noPrompt && password == "" {
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
			password, err = prompt.Run()
			if err != nil {
				cobra.CompErrorln(err.Error())
				return
			}
		}

		// generate deposit data
		valSecrets := nodeConf.ValidatorCredentials

		err = valSecrets.GenerateDepositDataWithRange(nodeConf.DepositDetails, vRange)
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}

		// generate wallet
		err = valSecrets.GenerateKeystoreWithRange(vRange, password)
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}

		err = nodeConf.Save()
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
		err = nodeConf.Save()
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}

		fmt.Println("Validator wallet was successfully created. Type")
		fmt.Println(utils.ConsoleInBlue("        lukso network validator describe"))
		fmt.Println("to see data related to the validator setup. ")
		fmt.Println("A transaction wallet was created to pay for the deposit transaction. ")
		fmt.Printf("The transaction wallet needs at least [staking amount] + [gas costs] %s before you can create a deposit transaction!\n", nodeConf.GetChain().GetCurrencySymbol())
	},
}

func init() {
	setupCmd.AddCommand(setupRangeCmd)
	setupRangeCmd.Flags().StringP(CommandOptionExistingMnemonic, CommandOptionExistingMnemonicShort, "", "existing mnemonic")
	setupRangeCmd.Flags().Int64P(CommandOptionFrom, CommandOptionFromShort, -1, "from position of validator key")
	setupRangeCmd.Flags().Int64P(CommandOptionTo, CommandOptionToShort, -1, "from position of validator key")
	setupRangeCmd.Flags().StringP(CommandOptionPassword, CommandOptionPasswordShort, "", "password for the keystore")
	setupRangeCmd.Flags().BoolP(CommandOptionNoPrompt, CommandOptionNoPromptShort, false, "no prompt to continue")
}

func promptDoYouWantToContinue(msg string) bool {
	prompt := promptui.Prompt{
		Label:     msg,
		IsConfirm: true,
	}
	result, err := prompt.Run()
	if err != nil {
		return false
	}
	return result == "y"
}
