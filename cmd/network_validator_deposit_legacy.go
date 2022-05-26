/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/lukso-network/lukso-cli/src/network"
	"github.com/spf13/cobra"
)

// depositLegacyCmd represents the deposit command
var depositLegacyCmd = &cobra.Command{
	Use:   "deposit_legacy",
	Short: "Send Deposit transactions to activate validator",
	Long: `After preparing wallets and deposit data, this command prepares deposit transactions to the deposit contract
address. Remember it will need your wallet address and private keys. Thus it will deduct balance from your wallet.

This tool is necessary to activate new validators`,
	Example: "lukso network validator deposit_legacy",
	Run: func(cmd *cobra.Command, args []string) {
		nodeConf := network.MustGetNodeConfig()
		valSecrets := nodeConf.GetValSecrets()
		if valSecrets == nil {
			cobra.CompErrorln("no validator credential is presented")
			return
		}
		err := valSecrets.SendDepositTxn(nodeConf.ApiEndpoints.ExecutionApi)
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
	validatorCmd.AddCommand(depositLegacyCmd)
}
