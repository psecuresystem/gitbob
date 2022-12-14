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
	"github.com/psecuresystem/gitb/helpers"
	"github.com/spf13/cobra"
)

// kCmd represents the k command
var kCmd = &cobra.Command{
	Use:   "k",
	Short: "Kinda like gitk",
	Long: `Cli command  to view refs`,
	Run: func(cmd *cobra.Command, args []string) {
		helpers.IterRefs()
	},
}

func init() {
	rootCmd.AddCommand(kCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
