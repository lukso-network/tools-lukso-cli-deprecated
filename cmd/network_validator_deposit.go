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

const CommandOptionDry = "dry"

// depositCmd represents the deposit command
var depositCmd = &cobra.Command{
	Use:     "deposit",
	Short:   "Send deposit transactions to activate validators",
	Long:    `Prepares and sends deposit transactions to activate new validators.`,
	Example: "lukso network validator deposit",
	Run: func(cmd *cobra.Command, args []string) {
		dry, _ := cmd.Flags().GetBool(CommandOptionDry)
		if dry {
			fmt.Println("THIS IS A DRY RUN")
		}
		nodeConf := network.MustGetNodeConfig()

		credentials := nodeConf.ValidatorCredentials
		if credentials == nil {
			utils.PrintColoredError("no validator credentials exist. Did you forget to setup your validators?")
			utils.Coloredln("    lukso network validator setup")
			return
		}

		maxGasFee, err := cmd.Flags().GetInt64(CommandOptionMaxGasFee)
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}

		priorityGasFee, err := cmd.Flags().GetInt64(CommandOptionPriorityFee)
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}

		vc := nodeConf.ValidatorCredentials
		// should never happen
		if vc == nil {
			cobra.CompErrorln("couldn't find contract details")
			return
		}

		events, err := network.NewDepositEvents(nodeConf.DepositDetails.ContractAddress, nodeConf.ApiEndpoints.ExecutionApi)
		if err != nil {
			cobra.CompErrorln(fmt.Sprintf("couldn't load deposit data from contract, reason: %s", err.Error()))
			return
		}

		fmt.Println("Past deposit events loaded", len(events.Events))

		totalDeposits, err := network.Deposit(&events, nodeConf.DepositDetails.DepositFileLocation, nodeConf.DepositDetails.ContractAddress, nodeConf.TransactionWallet.PrivateKey, nodeConf.ApiEndpoints.ExecutionApi, maxGasFee, priorityGasFee, dry)
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}
		if !dry {
			if totalDeposits == 0 {
				fmt.Printf("You didn't manage to deposit any keys:\n")
			} else {
				fmt.Printf("You successfully deposited %d key(s)! Your keys need to be activated which takes around 8h. You can check the status by calling:\n", totalDeposits)
				fmt.Println(utils.ConsoleInBlue("        lukso network validator describe"))
			}
		} else {
			fmt.Println("THIS WAS A DRY RUN - you didn't deposit any keys")
		}
	},
}

func init() {
	validatorCmd.AddCommand(depositCmd)

	depositCmd.Flags().Int64P(CommandOptionMaxGasFee, CommandOptionMaxGasFeeShort, 2500000014, "set the max gas price for transactions")
	depositCmd.Flags().Int64P(CommandOptionPriorityFee, CommandOptionPriorityFeeShort, 2500000000, "set priority price for transactions")
	depositCmd.Flags().BoolP(CommandOptionDry, "d", false, "don't run the transactions but just prepare it")
}
