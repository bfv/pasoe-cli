/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
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
