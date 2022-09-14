/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/lukso-network/lukso-cli/src/network"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// networkCmd represents the network command
var networkCmd = &cobra.Command{
	Use:   "network",
	Short: "LUKSO network related functionality",
}

func init() {
	rootCmd.AddCommand(networkCmd)
	//TODO: Make all usage sentences with uppercase
	//networkCmd.PersistentFlags().StringVar(&cfgFile, network.CommandOptionNodeConf, "", "config file (default is MY_NODE_DIRECTORY/node_config.yaml)")
	networkCmd.PersistentFlags().String(CommandOptionChain, network.DefaultNetworkID, "Provide chain you want to target [l16,...]")
	viper.BindPFlag(CommandOptionChain, networkCmd.PersistentFlags().Lookup(CommandOptionChain))
}
