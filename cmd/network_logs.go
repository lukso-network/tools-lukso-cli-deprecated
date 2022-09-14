/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

var (
	tail   string
	follow bool
)

// logsCmd represents the log command
var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Shows logs of the different clients",
}

func init() {
	networkCmd.AddCommand(logsCmd)
	logsCmd.PersistentFlags().BoolVarP(&follow, "follow", "f", false, "follow log output as a stream")
	logsCmd.PersistentFlags().StringVar(&tail, "tail", "5", "display last part of the log")
}
