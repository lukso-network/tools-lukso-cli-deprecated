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

import "github.com/spf13/cobra"

// upgradeCmd represents the upgrades command
var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrades the LUKSO CLI to a specific version",
}

func init() {
	rootCmd.AddCommand(upgradeCmd)
}
