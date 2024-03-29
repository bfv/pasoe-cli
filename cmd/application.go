/*
Copyright © 2022 Bronco Oostermeyer <dev@bfv.io>

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

	"github.com/bfv/pasoe-cli/logic"
	"github.com/spf13/cobra"
)

// applicationCmd represents the application command
var applicationCmd = &cobra.Command{
	Use:     "application",
	Aliases: []string{"app"},
	Short:   "Command for action related to (OEABL) applications",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("application called")
	},
}

var applicationListCmd = &cobra.Command{
	Use:   "list",
	Short: "Command for listing to (OEABL) applications",
	Long:  ``,
	Run:   listApplications,
}

func init() {

	rootCmd.AddCommand(applicationCmd)
	applicationCmd.AddCommand(applicationListCmd)

	applicationListCmd.Flags().BoolP("verbose", "v", false, "verbose ouput, lists webapps and enabled transports")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// applicationCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// applicationCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func listApplications(cmd *cobra.Command, args []string) {
	verbose, _ := cmd.Flags().GetBool("verbose")
	logic.ListApplications(instance, verbose)
}
