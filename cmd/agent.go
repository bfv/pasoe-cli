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
	"strings"

	"github.com/bfv/pasoe-cli/logic"
	"github.com/spf13/cobra"
)

// agentCmd represents the agent command
var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("agent called")
	},
}

var agentAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add agents",
	Long:  ``,
	Run:   addAgents,
}

var agentKillCmd = &cobra.Command{
	Use:   "kill",
	Short: "Kill agents",
	Long:  ``,
	Run:   killAgents,
}

var agentListCmd = &cobra.Command{
	Use:   "list",
	Short: "List agents",
	Long:  ``,
	Run:   listAgents,
}

func init() {
	agentCmd.AddCommand(agentAddCmd)
	agentCmd.AddCommand(agentKillCmd)
	agentCmd.AddCommand(agentListCmd)

	agentCmd.PersistentFlags().String("app", "", "Application name")

	// add cmd
	agentAddCmd.Flags().IntP("number", "n", 1, "number of agents to add")

	// kill cmd
	agentKillCmd.Flags().IntP("number", "n", 1, "number of agents to kill")
	agentKillCmd.Flags().Bool("all", false, "kill all agents")

	rootCmd.AddCommand(agentCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// agentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// agentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func addAgents(cmd *cobra.Command, args []string) {
	apps := getApps(cmd)
	n, _ := cmd.Flags().GetInt("number")
	logic.AddAgents(instance, apps, n)
}

func killAgents(cmd *cobra.Command, args []string) {

	params := logic.KillAgentParams{
		Apps: getApps(cmd),
		Ids:  args,
	}

	n, _ := cmd.Flags().GetInt("number")
	killAll, _ := cmd.Flags().GetBool("all")
	if killAll {
		n = -1
	}
	if len(args) > 0 {
		n = -1
	}

	params.Number = n

	logic.KillAgents(instance, params)
}

func listAgents(cmd *cobra.Command, args []string) {
	apps := getApps(cmd)
	logic.ListAgents(instance, apps)
}

func getApps(cmd *cobra.Command) []string {
	var apps []string
	appString, err := cmd.Flags().GetString("app")
	if err == nil && appString != "" {
		apps = strings.Split(appString, ",")
	}
	return apps
}
