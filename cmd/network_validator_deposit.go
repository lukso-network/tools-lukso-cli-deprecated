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

const CommandOptionGasPrice = "gasPrice"
const CommandOptionDry = "dry"

// depositCmd represents the deposit command
var depositCmd = &cobra.Command{
	Use:   "deposit",
	Short: "Send Deposit transactions to activate validator",
	Long: `After preparing wallets and deposit data, this command prepares deposit transactions to the deposit contract
address. Remember it will need your wallet address and private keys. Thus it will deduct balance from your wallet.

This tool is necessary to activate new validators`,
	Example: "lukso network validator deposit",
	Run: func(cmd *cobra.Command, args []string) {
		dry, _ := cmd.Flags().GetBool(CommandOptionDry)
		if dry {
			fmt.Println("THIS IS A DRY RUN")
		}
		nodeConf := network.MustGetNodeConfig()

		valSecrets := nodeConf.GetValSecrets()
		if valSecrets == nil || valSecrets.Deposit == nil || valSecrets.Eth1Data == nil {
			utils.PrintColoredError("no validator credentials exist. Did you forget to setup your validators?")
			utils.Coloredln("    lukso network validator setup")
			return
		}

		gasPrice, err := cmd.Flags().GetInt64(CommandOptionGasPrice)
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

		events, err := network.NewDepositEvents(vc.Deposit.ContractAddress, nodeConf.ApiEndpoints.ExecutionApi)
		if err != nil {
			cobra.CompErrorln(fmt.Sprintf("couldn't load deposit data from contract, reason: %s", err.Error()))
			return
		}

		fmt.Println("Past deposit events loaded", len(events.Events))

		totalDeposits, err := network.Deposit(&events, valSecrets.Deposit.DepositFileLocation, valSecrets.Deposit.ContractAddress, valSecrets.Eth1Data.WalletPrivKey, nodeConf.ApiEndpoints.ExecutionApi, gasPrice, dry)
		if err != nil {
			cobra.CompErrorln(err.Error())
			return
		}
		if !dry {
			fmt.Printf("You successfully deposited %d key(s)! Your keys need to be activated which takes around 8h. You can check the status by calling:\n", totalDeposits)
			fmt.Println(utils.ConsoleInBlue("        lukso network validator describe"))
		} else {
			fmt.Println("THIS WAS A DRY RUN - you didn't deposit any keys")
		}
	},
}

func init() {
	validatorCmd.AddCommand(depositCmd)

	depositCmd.Flags().Int64P(CommandOptionGasPrice, "g", 1000000, "set the gas price for transactions")
	depositCmd.Flags().BoolP(CommandOptionDry, "d", false, "don't run the transactions but just prepare it")
}
