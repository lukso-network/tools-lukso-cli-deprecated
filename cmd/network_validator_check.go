/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/lukso-network/lukso-cli/src/network"
	"github.com/lukso-network/lukso-cli/src/utils"
	"github.com/spf13/cobra"
)

// validatorCheckCmd represents the check command
var validatorCheckCmd = &cobra.Command{
	Use:     "check",
	Short:   "Show detailed status of the validators as deposited in the Deposit Contract",
	Long:    `Show detailed status of the validators as deposited in the Deposit Contract.`,
	Example: "lukso network validator check",
	Run: func(cmd *cobra.Command, args []string) {

		nodeConf := network.MustGetNodeConfig()
		baseUrl := nodeConf.ApiEndpoints.ConsensusApi
		vc := nodeConf.ValidatorCredentials

		events, err := network.NewDepositEvents(vc.Deposit.ContractAddress, nodeConf.ApiEndpoints.ExecutionApi)
		if err != nil {
			cobra.CompErrorln(fmt.Sprintf("couldn't load deposit data from contract, reason: %s", err.Error()))
			return
		}

		fmt.Println("Beacon API: ", baseUrl)
		fmt.Println("Checking status of all deposited keys")
		for _, d := range events.Events {
			fmt.Println("........................................................................................................................................................................")
			utils.ColoredPrintln("Index", d.Index)
			err := describe(&events, []string{d.Pubkey}, baseUrl, -1)

			if err != nil {
				utils.PrintColoredError(err.Error())
			}
		}
	},
}

func init() {
	validatorCmd.AddCommand(validatorCheckCmd)
}
