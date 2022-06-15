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
	"errors"
	"fmt"
	"github.com/lukso-network/lukso-cli/src/network"
	"github.com/lukso-network/lukso-cli/src/utils"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initCmd represents the setup command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes node by downloading configs and scripts",
	Long: `This command downloads network starter scripts and config files
from the github repository. It also updates node name and IP address in the .env file`,
	Example: "lukso network init --chain l16 --nodeName my_node",
	RunE: func(cmd *cobra.Command, args []string) error {
		chain := network.GetChainByString(viper.GetString(CommandOptionChain))

		// default is Mainnet but not supported yet
		if chain == network.MainNet {
			utils.PrintColoredError("you are trying to init a mainnet node...you are ahead of your time")
			return nil
		}

		statsName, _ := cmd.Flags().GetString(CommandOptionNodeName)
		devConfig, _ := cmd.Flags().GetString(CommandOptionDevConfig)

		if chain == network.Dev && devConfig == "" {
			utils.PrintColoredError("you need to provide a dev location if you want to use a dev chain: lukso network init --statsName [NAME] -d test --chain dev")
			return nil
		}

		if statsName == "" {
			// set number of validators
			prompt := promptui.Prompt{
				Label: "Enter the name of your node (how it appears on the stat pages)",
				Validate: func(input string) error {
					if input == "" {
						return errors.New("your node name cannot be empty")
					}
					return nil
				},
			}

			name, err := prompt.Run()
			if err != nil {
				utils.PrintColoredErrorWithReason("couldn't read stats name", err)
				return nil
			}
			statsName = name
		}

		nodeConf := network.GetDefaultNodeConfigsIfDoesntExist(chain)
		if nodeConf == nil {
			fmt.Println("A node is already setup in this location. Choose another location to setup a node for a different chain or modify this node by editing ./node_conf.yaml.")
			return nil
		}

		// Get IP And HostName
		nodeDetails, err := network.GetIPAndHostName(statsName)
		if err != nil {
			utils.PrintColoredError(fmt.Sprintf("\ncouldn't get ip or host name, reason %s", err.Error()))
			return nil
		}
		nodeConf.Node = nodeDetails
		// download node params
		nodeParams := network.NewNodeParamsLoader()
		fmt.Println("devConfig:", devConfig)

		location := ""
		if chain == network.Dev {
			location = nodeParams.GetDevLocation(devConfig)
		} else {
			location = nodeParams.GetLocation(chain)
		}

		fmt.Printf("Loading node params from  %s ...", location)
		err = nodeParams.LoadNodeParams(location)
		if err != nil {
			fmt.Println("unsuccessful")
			utils.PrintColoredError(fmt.Sprintf("couldn't load node params for dev chain, reason: %s", err.Error()))
			return nil
		}
		nodeConf.ApiEndpoints = &network.NodeApi{
			ConsensusApi: nodeParams.ConsensusAPI,
			ExecutionApi: nodeParams.ExecutionAPI,
		}
		nodeConf.Chain.ID = nodeParams.NetworkID
		nodeConf.Execution.StatsAddress = nodeParams.ExecutionStats
		nodeConf.Consensus.StatsAddress = nodeParams.ConsensusStats
		nodeConf.DepositDetails.Amount = nodeParams.MinStakeAmount
		nodeConf.Execution.Version = nodeParams.GethVersion
		nodeConf.Consensus.Version = nodeParams.PrysmVersion

		fmt.Println("success")

		err = nodeConf.Save()
		if err != nil {
			utils.PrintColoredError(fmt.Sprintf("\ncouldn't save %s, reason %s", network.NodeConfigLocation, err.Error()))
			return nil
		}

		// when chain is dev -> download the settings for this chain
		if chain == network.Dev {
			err = network.SetupDevNetwork(devConfig)
		} else {
			err = network.SetupNetwork(chain)
		}
		if err != nil {
			cobra.CompError(err.Error())
			return nil
		}

		if chain == network.Dev {
			_, err = nodeConf.UpdateDevBootnodes(devConfig)
		} else {
			_, err = nodeConf.UpdateBootnodes()
		}
		if err != nil {
			fmt.Println("couldn't update bootnodes, reason:", err.Error())
		}

		updateEnv()

		fmt.Printf("You successfully prepared the node for chain %s!!!\n", chain.String())
		return nil
	},
}

func init() {
	networkCmd.AddCommand(initCmd)
	initCmd.Flags().String(CommandOptionNodeName, "", "name of your node as it appears in the stats services")
	initCmd.Flags().StringP(CommandOptionDevConfig, "d", "", "location of the dev configuration")
}
