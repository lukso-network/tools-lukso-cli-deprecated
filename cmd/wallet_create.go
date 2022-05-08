/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/lukso-network/lukso-cli/src/wallet"
	"io/ioutil"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a wallet and a password file ",
	Long: `This command will create a wallet and password file in a target directory. Optionally a password and a label for the filenames can be given:

lukso wallet create -p [PASSWORD] -d [TARGET_DIRECTORY] -l [LABEL]`,
	Run: func(cmd *cobra.Command, args []string) {
		dir, _ := cmd.Flags().GetString("dir")
		password, _ := cmd.Flags().GetString("password")
		label, _ := cmd.Flags().GetString("label")

		if password == "" {
			password = wallet.CreateRandomPassword()
		}

		store := keystore.NewKeyStore(dir, keystore.StandardScryptN, keystore.StandardScryptP)

		a, err := store.NewAccount(password)

		if err != nil {
			fmt.Println("error while creating new account", err.Error())
			return
		}

		filename := a.URL.String()
		passwordFilename := strings.Replace(a.URL.Path, a.URL.Scheme, "", 1)

		if label != "" {

			if dir == "" {
				filename = fmt.Sprintf("%v.json", label)
			} else {
				filename = fmt.Sprintf("%v/%v.json", dir, label)
			}
			err = os.Rename(strings.Replace(a.URL.Path, a.URL.Scheme, "", 1), filename)

			if err != nil {
				fmt.Println("error while renaming wallet file", err.Error())
				return
			}

			passwordFilename = label
		}

		// write password file
		if dir == "" {
			err = ioutil.WriteFile(fmt.Sprintf("%v_password.txt", passwordFilename), []byte(password), os.ModePerm)
		} else {
			err = ioutil.WriteFile(fmt.Sprintf("%v/%v_password.txt", dir, passwordFilename), []byte(password), os.ModePerm)
		}

		if err != nil {
			fmt.Println("error while writing password file", err.Error())
		}

		fmt.Println("Successfully created wallet...")
		fmt.Print("Account", strings.ToLower(a.Address.String()))
		fmt.Print("Location", filename)
	},
}

func init() {
	walletCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("label", "l", "", "indicates the name of the wallet and password file")
	createCmd.Flags().StringP("dir", "d", "", "is the target directory of the wallet and password file")
	createCmd.Flags().StringP("password", "p", "", "password of the wallet stored in password file")
}
