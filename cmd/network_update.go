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
		err = config.UpdateBootnodes()
		if err != nil {
			fmt.Printf("couldn't update bootnodes reason: %s\n", err.Error())
		} else {
			fmt.Println("Successfully updated bootnodes -> restart your nodes to make the changes effective")
		}
	},
}

func init() {
	networkCmd.AddCommand(updateCmd)
}
