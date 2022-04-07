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
	"github.com/lukso-network/lukso-cli/src"
	"github.com/lukso-network/lukso-cli/src/network"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initCmd represents the setup command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes node by downloading configs and scripts",
	Long: `This command downloads network starter scripts and config files
from the github repository. It also updates node name and IP address in the .env file`,
	Example: "lukso-cli network init --nodeconf ./node_config.yaml --chainId l16 --nodeName my_node --docker",
	RunE: func(cmd *cobra.Command, args []string) error {

		viper.Set(src.ViperKeyNetworkName, viper.GetString("chainId"))
		return network.SetupNetwork(viper.GetString("nodeName"))
	},
}

func init() {
	networkCmd.AddCommand(initCmd)

	initCmd.Flags().String("chainId", src.DefaultNetworkID, "provide chainId for the LUKSO network")
	initCmd.Flags().Bool("docker", true, "use docker or not")
	initCmd.Flags().String("nodeName", "", "set node name")

	viper.BindPFlag("chainId", initCmd.Flags().Lookup("chainId"))
	viper.BindPFlag("docker", initCmd.Flags().Lookup("docker"))
	viper.BindPFlag("nodeName", initCmd.Flags().Lookup("nodeName"))
}
