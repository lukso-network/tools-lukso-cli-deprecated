/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/lukso-network/lukso-cli/api"
	"github.com/lukso-network/lukso-cli/api/gethrpc"
	"github.com/lukso-network/lukso-cli/src/network"
	"github.com/lukso-network/lukso-cli/src/utils"
	"github.com/spf13/cobra"
)

// nodeCmd represents the node command
var nodeCmd = &cobra.Command{
	Use:     "node",
	Short:   "gets a description of peers",
	Long:    `shows the peers of this node`,
	Example: "lukso network describe node",
	Run: func(cmd *cobra.Command, args []string) {
		nodeConf := network.MustGetNodeConfig()

		utils.ColoredPrintln("Chain", nodeConf.Chain.Name)
		utils.ColoredPrintln("NetworkId", nodeConf.Chain.ID)

		c := gethrpc.NewRPCClient("http://localhost:8545")

		nodeInfo, err := api.AdminNodeInfoRequest(c)

		if err != nil {
			fmt.Println("Error: ", err.Error())
		}

		peers, err := api.AdminPeersRequest(c)
		if err != nil {
			fmt.Println("Error: ", err.Error())
		}

		blocknumber, err := api.BlockNumber(c)
		if err != nil {
			fmt.Println("Error: ", err.Error())
		}

		utils.Coloredln("Execution:")
		utils.Coloredln(".........................................")
		utils.ColoredPrintln("Bootnode: ", nodeConf.Consensus.Bootnode)
		utils.ColoredPrintln("Enode: ", nodeInfo.Enode)
		utils.ColoredPrintln("Peers: ", len(peers))
		utils.ColoredPrintln("Latest Block:", blocknumber)

		utils.Coloredln(".........................................")
		utils.Coloredln("Consensus:")
		utils.Coloredln(".........................................")
		utils.ColoredPrintln("Bootnode: ", nodeConf.Execution.Bootnode)
		utils.ColoredPrintln("Peers: ", len(peers))
		utils.ColoredPrintln("Latest Block:", blocknumber)

	},
}

func init() {
	describeCmd.AddCommand(nodeCmd)
}
