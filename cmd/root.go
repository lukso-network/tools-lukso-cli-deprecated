/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

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
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
