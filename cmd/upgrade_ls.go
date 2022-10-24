/*
Copyright © 2022 The LUKSO authors

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either upgrade 3 of the License, or
(at your option) any later upgrade.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"github.com/lukso-network/lukso-cli/src/upgrade"
	"github.com/spf13/cobra"
)

// versionLsCmd represents the list command
var versionLsCmd = &cobra.Command{
	Use:     "ls",
	Short:   "Lists LUKSO CLI versions available. Indicates which upgrade of the CLI has been installed.",
	Example: "lukso upgrade ls",
	RunE: func(cmd *cobra.Command, args []string) error {
		currentVersion := rootCmd.Version
		return upgrade.List(currentVersion)
	},
}

func init() {
	upgradeCmd.AddCommand(versionLsCmd)
}
