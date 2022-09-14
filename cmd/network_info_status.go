/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/lukso-network/lukso-cli/src/network"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Describes the status of the network",
	Long: `Describes the status of the network regarding PoS. The Beacon Api node can be chosen and the status perspective can be moved by giving an epoch number.
Leaving out the epoch number gives the latest status.

	
`, Example: "lukso network status",
	Run: func(cmd *cobra.Command, args []string) {
		baseUrl, _ := cmd.Flags().GetString("beaconapi")
		if baseUrl == "" {
			baseUrl = network.GetDefaultNodeConfigByOptionParam(viper.GetString(CommandOptionChain)).ApiEndpoints.ConsensusApi
		}
		epoch, _ := cmd.Flags().GetInt64("epoch")

		err := network.DescribeNetwork(baseUrl, epoch)
		if err != nil {
			cobra.CompError(err.Error())
		}
	},
}

func init() {
	networkCmd.AddCommand(statusCmd)

	statusCmd.Flags().StringP("beaconapi", "b", "", "endpoint of beacon api")
	statusCmd.Flags().Int64P("epoch", "e", -1, "epoch to be statusd - if left out it is the head epoch")
}
