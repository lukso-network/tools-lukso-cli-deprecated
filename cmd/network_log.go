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

// logCmd represents the log command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Subcommand \"log\" for LUKSO log related things",
}

func init() {
	networkCmd.AddCommand(logCmd)
	logCmd.PersistentFlags().BoolVarP(&follow, "follow", "f", false, "follow log output as a stream")
	logCmd.PersistentFlags().StringVar(&tail, "tail", "5", "display last part of the log")
}
