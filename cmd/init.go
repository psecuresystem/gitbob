/*
Copyright © 2022 Vision Onyeaku <visiondaniels32@gmail.com>

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
	"os"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a gitb repository",
	Long: `This command is used to initialize a new gitb repository in the root of the project. 
	If a repo already exists, it uses that one. 
	By repo I mean a .gitb folder`,
	Run: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()
		if _, err := os.Stat(".gitb"); os.IsNotExist(err) {
			os.Mkdir(".gitb", os.ModePerm)
			os.MkdirAll(".gitb/objects", os.ModePerm)
			os.MkdirAll(".gitb/refs", os.ModePerm)
			os.MkdirAll(".gitb/refs/tags", os.ModePerm)
			fmt.Printf("Empty Repository Created in %s/.gitb\n", wd)
		} else {
			fmt.Printf("Re Initialized Repository in %s/.gitb\n", wd)
		}
		
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
