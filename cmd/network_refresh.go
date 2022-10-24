package cmd

import (
	"fmt"
	"github.com/lukso-network/lukso-cli/src/network"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// refreshCmd represents the refresh command
var refreshCmd = &cobra.Command{
	Use:     "refresh",
	Short:   "Restarts the network and updates the configuration files. This command might be removed in the future ",
	Example: "lukso network refresh",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO:
		fmt.Println("stopping running containers")
		err := network.DownDockerServices()
		if err != nil {
			cobra.CompErrorln(err.Error())
			os.Exit(1)
		}
		updateEnv()
		updateCmd.Run(cmd, args)
		// take a small break between stop and start
		time.Sleep(2 * time.Second)
		fmt.Println("starting clients")
		err = network.StartArchNode()
		if err != nil {
			cobra.CompErrorln(err.Error())
			os.Exit(1)
		}
		network.StartValidatorNode()
	},
}

func init() {
	networkCmd.AddCommand(refreshCmd)
}
