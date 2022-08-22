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
	"github.com/bfv/pasoe-cli/logic"
	"github.com/spf13/cobra"
)

// agentSessionCmd represents the agentsession command
var agentSessionCmd = &cobra.Command{
	Use:     "agentsession",
	Aliases: []string{"as"},
	Short:   "A command handling agent sessions",
	Long:    ``,
	Run:     listAgentSessions, // default to list sub command
}

var agentSessionListCmd = &cobra.Command{
	Use:   "list",
	Short: "List agent sessions",
	Long:  ``,
	Run:   listAgentSessions,
}

func init() {
	rootCmd.AddCommand(agentSessionCmd)

	agentSessionCmd.AddCommand(agentSessionListCmd)
	agentSessionCmd.PersistentFlags().String("app", "", "Application name")

	agentSessionListCmd.Flags().Int("pid", -1, "Pid")
	agentSessionListCmd.Flags().Int("threshold", -1, "Treshold in MiB")
	agentSessionListCmd.Flags().BoolP("verbose", "v", false, "Verbose output")
	agentSessionListCmd.Flags().Bool("ato", false, "List entries above threshold only")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// agentSessionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// agentSessionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func listAgentSessions(cmd *cobra.Command, args []string) {

	apps := getApps(cmd)

	pid, _ := cmd.Flags().GetInt("pid")
	threshold, _ := cmd.Flags().GetInt("threshold")
	aboveTresholdOnly, _ := cmd.Flags().GetBool("ato")
	verbose, _ := cmd.Flags().GetBool("verbose")

	logic.ListAgentSessions(instance, apps, logic.ListAgentSessionParams{
		Pid:               pid,
		Treshold:          threshold,
		AboveTresholdOnly: aboveTresholdOnly,
		Verbose:           verbose,
	})
}
