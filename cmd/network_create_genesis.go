/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/lukso-network/lukso-cli/src/network"
	"github.com/spf13/cobra"
)

// createGenesisCmd represents the createGenesis command
var createGenesisCmd = &cobra.Command{
	Use:   "createGenesis",
	Short: "creates validator wallets according to a hardcoded config",
	Long: `Creates validator wallets given a mnemonic and a configType. The config type must correspond to a hardcoded configuration.
The configuration defines the validator set from 1...n that can be consumed by a validator node.`,
	Run: func(cmd *cobra.Command, args []string) {
		mnemonic, _ := cmd.Flags().GetString("mnemonic")
		configType, _ := cmd.Flags().GetString("configType")

		err := network.CreateValidatorWallets(mnemonic, configType)
		if err != nil {
			fmt.Println("could not create wallets", err.Error())
		}
	},
}

func init() {
	networkCmd.AddCommand(createGenesisCmd)
	createGenesisCmd.Flags().StringP("mnemonic", "m", "", "all wallets are dereived from this wallet")
	createGenesisCmd.MarkFlagRequired("mnemonic")
	createGenesisCmd.Flags().StringP("configType", "c", "", "indicates the way the wallets are chunked")
	createGenesisCmd.MarkFlagRequired("configType")
}
