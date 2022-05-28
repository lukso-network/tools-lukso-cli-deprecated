/*
Copyright © 2022 The LUKSO authors

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"
	"github.com/lukso-network/lukso-cli/src/network"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// initCmd represents the setup command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes node by downloading configs and scripts",
	Long: `This command downloads network starter scripts and config files
from the github repository. It also updates node name and IP address in the .env file`,
	Example: "lukso network init --chain l16 --nodeName my_node",
	RunE: func(cmd *cobra.Command, args []string) error {
		chain := network.GetChainByString(viper.GetString(network.CommandOptionChain))
		nodeName := viper.GetString(network.CommandOptionNodeName)

		isGenerated, err := network.GenerateDefaultNodeConfigsIfDoesntExist(chain)
		if err != nil {
			cobra.CompErrorln(err.Error())
			os.Exit(1)
		}

		if !isGenerated {
			fmt.Println("A node is already setup in this location. Choose another location to setup a node for a different chain or modify this node by editing ./node_conf.yaml.")
			return nil
		}
		config := network.MustGetNodeConfig()

		err = network.SetupNetwork(chain, nodeName)
		if err != nil {
			cobra.CompError(err.Error())
		}

		_, err = config.UpdateBootnodes()
		if err != nil {
			fmt.Println("couldn't update bootnodes, reason:", err.Error())
		}

		fmt.Printf("You successfully prepared the node for chain %s!!!\n", chain.String())
		return nil
	},
}

func init() {
	networkCmd.AddCommand(initCmd)
	initCmd.Flags().String(network.CommandOptionNodeName, "", "set node name")
	viper.BindPFlag(network.CommandOptionNodeName, initCmd.Flags().Lookup(network.CommandOptionNodeName))
}
