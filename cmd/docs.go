/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

// docsGenerateCmd represents the docsGenerate command
var docsGenerateCmd = &cobra.Command{
	Use:   "docs",
	Short: "Generate documents for lukso-cli",
	Long: `This command generates documents for all the available command of lukso-cli.
You will find documentations inside "docs" directory`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Generating docs...")
		err := doc.GenMarkdownTree(rootCmd, "docs")
		if err != nil {
			cobra.CompError(err.Error())
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(docsGenerateCmd)
}
