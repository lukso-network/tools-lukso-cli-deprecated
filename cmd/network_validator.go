/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// validatorLogCmd represents the validator command
var validatorCmd = &cobra.Command{
	Use:   "validator",
	Short: "Validator client commands",
}

func init() {
	networkCmd.AddCommand(validatorCmd)
}
