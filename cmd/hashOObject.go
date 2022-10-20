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
	"os"

	"github.com/psecuresystem/gitb/helpers"
	"github.com/spf13/cobra"
)

// hashObjectCmd represents the hashObject command
var hashObjectCmd = &cobra.Command{
	Use:   "hash-object",
	Short: "This command hashes file and saves it in object database",
	Long: `This command reads the content of a file, generates a sha256 hash for it and stores the content in .gitb/objects/sha256hash`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.ParseFlags(args)
		fileName, _ := cmd.Flags().GetString("file")
		if fileName == "" {
			fmt.Println("File to hash is required")
			return
		}
		if _, err := os.Stat(".gitb"); os.IsNotExist(err) {
			fmt.Println("No gitb repository detected")
			return
		}
		fileContent, _ := os.ReadFile(fileName)
		helpers.HashObject(string(fileContent), "blob")
	},
}

func init() {
	rootCmd.AddCommand(hashObjectCmd)

	hashObjectCmd.PersistentFlags().String("file", "", "The file to hash")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hashObjectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hashObjectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
