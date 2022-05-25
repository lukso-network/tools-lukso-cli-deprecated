/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/lukso-network/lukso-cli/src/network"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// describeCmd represents the describe command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "updates the node with the latest params",
	Long: `Updates the with the latest params [bootnodes, client versions,...]. You need to restart the nodes to make the changes become effective.

	
`, Example: "lukso network update",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := network.GetLoadedNodeConfigs()
		if err != nil {
			cobra.CompError(err.Error())
			return
		}

		// trying to "repair node conf" to make it downwards compatible
		if config.Chain == nil {
			chain := network.GetChainByString(viper.GetString(network.CommandOptionChain))
			defaultConfig := network.GetDefaultNodeConfig(chain)
			config.Chain = defaultConfig.Chain
			err := config.WriteOrUpdateNodeConfig()
			if err != nil {
				cobra.CompError(err.Error())
				return
			}
		}

		chain := network.GetChainByString(config.Chain.Name)
		bootnodes, err := network.NewBootnodeUpdater(chain).DownloadLatestBootnodes()
		if err != nil {
			cobra.CompError(err.Error())
			return
		}

		if len(bootnodes) == 0 {
			fmt.Println("No bootnodes available for this chain ", chain.String())
		}

		hasUpdates := false
		if config.Consensus.Bootnode != bootnodes[0].Consensus {
			fmt.Println("Updating bootnode for the consensus chain...")
			hasUpdates = true
			config.Consensus.Bootnode = bootnodes[0].Consensus
		}
		if config.Execution.Bootnode != bootnodes[0].Execution {
			fmt.Println("Updating bootnode for the execution chain...")
			hasUpdates = true
			config.Execution.Bootnode = bootnodes[0].Execution
		}

		if !hasUpdates {
			fmt.Println("everything up to date")
		} else {
			err = config.WriteOrUpdateNodeConfig()
			if err != nil {
				cobra.CompError(fmt.Sprintf("Couldn't update bootnodes reason: %s", err.Error()))
				return
			}
			fmt.Println("Successfully updated bootnodes -> restart your nodes to make the changes effective")
		}
	},
}

func init() {
	networkCmd.AddCommand(updateCmd)
}
