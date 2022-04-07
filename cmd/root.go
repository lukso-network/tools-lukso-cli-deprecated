/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "Command Line Interface to spin up a LUKSO node",
	Long: `A Command Line Interface to spin up and maintain different components of LUKSO network. This CLI
is helpful to spin up a full node as well as it monitors log for all the components (e.g. execution engine, consensus engine, eth2stats client and validator client).
One can also create deposit data, validator credentials and submit deposit transactions to the LUKSO network.
`,
	Version: "v0.0.1",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "nodeconf", "", "config file (default is $HOME/.lukso/node_config.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		configFilePath := path.Join(home, ".lukso")

		// Search config in home directory with name ".cli" (without extension).
		viper.AddConfigPath(configFilePath)
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
