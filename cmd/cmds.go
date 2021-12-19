/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"github.com/spf13/cobra"
)

// cmdsCmd represents the cmds command
var cmdsCmd = &cobra.Command{
	Use:   "cmds",
	Short: "Command runner management utility",
	Long:  `Command runner management utility`,
	Args:  cobra.NoArgs,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("cmds called")
	// },
}

func init() {
	rootCmd.AddCommand(cmdsCmd)

}

// func printEvents(m map[string][]string) {
// 	for evt, cmds := range m {
// 		for i, cmd := range cmds {
// 			fmt.Printf("%s(%d): %s\n", evt, i, cmd)
// 		}
// 	}
// }
