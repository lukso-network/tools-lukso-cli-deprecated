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
	"github.com/lukso-network/lukso-cli/src/utils"
	"github.com/lukso-network/lukso-cli/src/version"
	"github.com/spf13/cobra"
)

// versionInstallCmd represents the 'install' command
var versionInstallCmd = &cobra.Command{
	Use:     "install",
	Short:   "Install a LUKSO CLI version locally.",
	Example: "lukso version install --version v0.4.2",
	RunE: func(cmd *cobra.Command, args []string) error {
		upgrade, _ := cmd.Flags().GetBool(CommandOptionUpgrade)
		if upgrade {
			latestVersion, err := version.GetLatestVersion()
			if err != nil {
				utils.PrintColoredError(err.Error())
				return err
			}
			err = version.Install(latestVersion)
			if err != nil {
				utils.PrintColoredError(err.Error())
				return err
			}
			return nil
		}

		// Install specified version
		specifiedVersion, err := cmd.Flags().GetString(CommandOptionVersion)
		if specifiedVersion == "" {
			utils.PrintColoredError("please specify the version you want to install")
			return nil
		}
		if err != nil {
			utils.PrintColoredError(err.Error())
			return err
		}
		err = version.Install(specifiedVersion)
		if err != nil {
			utils.PrintColoredError(err.Error())
			return err
		}
		return nil
	},
}

func init() {
	versionCmd.AddCommand(versionInstallCmd)

	versionInstallCmd.Flags().StringP(CommandOptionVersion, "v", "", "Install the specified LUKSO CLI version.")
	versionInstallCmd.Flags().BoolP(CommandOptionUpgrade, "u", false, "Upgrade to the latest LUKSO CLI version.")
}
