/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/lukso-network/lukso-cli/src/network"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var cfgFile string

// networkCmd represents the network command
var networkCmd = &cobra.Command{
	Use:   "network",
	Short: "subcommand \"network\" for LUKSO network related things",
}

func init() {
	rootCmd.AddCommand(networkCmd)
	cobra.OnInitialize(initConfig)

	networkCmd.PersistentFlags().StringVar(&cfgFile, network.CommandOptionNodeConf, "", "config file (default is MY_NODE_DIRECTORY/node_config.yaml)")
	networkCmd.PersistentFlags().String(network.CommandOptionChain, network.DefaultNetworkID, "provide chain you want to target [l16,...]")

	viper.BindPFlag("chainId", networkCmd.PersistentFlags().Lookup(network.CommandOptionChain))
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		nodeConfigFileLocation := "./node_config.yaml"
		if !network.FileExists(nodeConfigFileLocation) {
			fmt.Println("No node_config.yaml found for this network. Generating node_config.yaml")
			chain := network.GetChainByString(viper.GetString(network.CommandOptionChain))
			err := network.GenerateDefaultNodeConfigs(chain)
			if err != nil {
				cobra.CompErrorln(err.Error())
				os.Exit(1)
			}
		}

		// Search config in home directory with name ".cli" (without extension).
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName("node_config")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		cobra.CompErrorln(err.Error())
		os.Exit(1)
	}
}
