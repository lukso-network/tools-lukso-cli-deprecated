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

// describeCmd represents the describe command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "updates the node with the latest params",
	Long: `Updates the with the latest params [bootnodes, client versions,...]. You need to restart the nodes to make the changes become effective.

	
`, Example: "lukso network update",
	Run: func(cmd *cobra.Command, args []string) {
		hasUpdates := false
		fmt.Println("Searching for updates")

		nodeConf := network.MustGetNodeConfig()
		chain := network.GetChainByString(nodeConf.Chain.Name)
		chainId := nodeConf.Chain.ID

		wasIPUpdated, err := nodeConf.UpdateExternalIP()
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

		wasBootnodeUpdated := false
		if chain == network.Dev {
			wasBootnodeUpdated, err = nodeConf.UpdateDevBootnodes(chainId)
		} else {
			wasBootnodeUpdated, err = nodeConf.UpdateBootnodes()
		}
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

		if chain != network.L16Beta {
			nodeParamsLoader := network.NewNodeParamsLoader()

			location := ""
			if chain == network.Dev {
				location = nodeParamsLoader.GetDevLocation(chainId)
			} else {
				location = nodeParamsLoader.GetLocation(chain)
			}

			fmt.Printf("Loading node params from  %s ...", location)
			err := nodeParamsLoader.LoadNodeParams(location)
			if err != nil {
				fmt.Println("unsuccessful")
				utils.PrintColoredError(fmt.Sprintf("couldn't load node params for chain, reason: %s", err.Error()))
			} else {
				hasUpdates = true
				nodeConf.ApiEndpoints = &network.NodeApi{
					ConsensusApi: nodeParamsLoader.ConsensusAPI,
					ExecutionApi: nodeParamsLoader.ExecutionAPI,
				}
				nodeConf.Chain.ID = nodeParamsLoader.NetworkID
				nodeConf.Execution.StatsAddress = nodeParamsLoader.ExecutionStats
				nodeConf.Consensus.StatsAddress = nodeParamsLoader.ConsensusStats
				nodeConf.DepositDetails.Amount = nodeParamsLoader.MinStakeAmount
				nodeConf.Execution.Version = nodeParamsLoader.GethVersion
				nodeConf.Consensus.Version = nodeParamsLoader.PrysmVersion
				err = nodeConf.Save()
				fmt.Println("")
				if err != nil {
					fmt.Println("couldn't update node params, reason:", err.Error())
				}
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
