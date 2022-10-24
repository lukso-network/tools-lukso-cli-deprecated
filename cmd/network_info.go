/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "network",
	Short: "Returns information on the LUKSO network",
}

func init() {
	networkCmd.AddCommand(infoCmd)
}
