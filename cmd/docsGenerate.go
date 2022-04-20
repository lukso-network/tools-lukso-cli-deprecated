/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// docsGenerateCmd represents the docsGenerate command
var docsGenerateCmd = &cobra.Command{
	Use:   "docsGenerate",
	Short: "Generate documents for lukso-cli",
	Long: `This command generates documents for all the available command of lukso-cli.
You will find documentations inside "docs" directory`,
	Run: func(cmd *cobra.Command, args []string) {
		generateDocuments()
	},
}

func init() {
	networkCmd.AddCommand(docsGenerateCmd)
}
