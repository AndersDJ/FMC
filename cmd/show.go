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
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use: "show",
	// Short: "A brief description of your command",
	Short: "show time",
	Long:  `show time for long`,

	Run: ShowTime,
}

func init() {
	rootCmd.AddCommand(showCmd)

}

func ShowTime(cmd *cobra.Command, args []string) {
	for _, v := range args {
		fmt.Println(v)
	}
	fmt.Println(time.Now())
}
