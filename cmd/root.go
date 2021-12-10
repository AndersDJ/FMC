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
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.FMC.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	flags := rootCmd.Flags()
	persistent := rootCmd.PersistentFlags()
	persistent.StringP("databases", "d", "./fmc.db", "database path")
	flags.BoolP("toggle", "t", false, "Help message for toggle")

	addServerFlags(flags)

}

func addServerFlags(flags *pflag.FlagSet) {
	flags.StringP("address", "a", "127.0.0.1", "address to listen")
	flags.StringP("log", "l", "stdout", "output log")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".FMC" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".FMC")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

var rootCmd = &cobra.Command{
	Use:   "FMC",
	Short: "A stylish web-based file browser",
	Long: `File Browser CLI lets you create the database to use with File Browser,
manage your users and all the configurations without acessing the
web interface.

If you've never run File Browser, you'll need to have a database for
it. Don't worry: you don't need to setup a separate database server.
We're using Bolt DB which is a single file database and all managed
by ourselves.

For this specific command, all the flags you have available (except
"config" for the configuration file), can be given either through
environment variables or configuration files.

If you don't set "config", it will look for a configuration file called
.filebrowser.{json, toml, yaml, yml} in the following directories:

- ./
- $HOME/
- /etc/filebrowser/

The precedence of the configuration values are as follows:

- flags
- environment variables
- configuration file
- database values
- defaults

The environment variables are prefixed by "FB_" followed by the option
name in caps. So to set "database" via an env variable, you should
set FB_DATABASE.

Also, if the database path doesn't exist, File Browser will enter into
the quick setup mode and a new database will be bootstraped and a new
user created with the credentials from options "username" and "password".`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}
