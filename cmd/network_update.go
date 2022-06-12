/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/lukso-network/lukso-cli/src/network"
	"github.com/spf13/cobra"
)

// describeCmd represents the describe command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "updates the node with the latest params",
	Long: `Updates the with the latest params [bootnodes, client versions,...]. You need to restart the nodes to make the changes become effective.

	
`, Example: "lukso network update",
	Run: func(cmd *cobra.Command, args []string) {
		hasUpdates := false
		fmt.Println("Searching for updates")

		config := network.MustGetNodeConfig()

		wasIPUpdated, err := config.UpdateExternalIP()
		if err != nil {
			fmt.Printf("couldn't update external IP reason: %s\n", err.Error())
		} else {
			if wasIPUpdated {
				fmt.Println("Successfully updated IP -> restart your nodes to make the changes effective")
				hasUpdates = true
			} else {
				fmt.Println("External IP is up to date")
			}
		}

		wasBootnodeUpdated, err := config.UpdateBootnodes()
		if err != nil {
			fmt.Printf("couldn't update bootnodes reason: %s\n", err.Error())
		} else {
			if wasBootnodeUpdated {
				fmt.Println("Successfully updated bootnodes")
				hasUpdates = true
			} else {
				fmt.Println("Bootnodes are up to date")
			}
		}

		if hasUpdates {
			fmt.Println("Successfully updated your node -> restart your nodes to make the changes effective!")
		} else {
			fmt.Println("Everything up to date!")
		}

	},
}

func init() {
	networkCmd.AddCommand(updateCmd)
}
