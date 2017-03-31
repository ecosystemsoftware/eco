// Copyright 2017 Jonathan Pincas

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// 	http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ghost

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "ghost [command] [arguments]",
	Short: "ghost command line tool",
	Long: `Use to initialise or launch the ghost server or create new users or bundles.
	Use the bare command 'ghost' to create a new config.json or verify an existing one.`,
	RunE: createConfigIfNotExists,
}

// serveCmd represents the serve command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Ping test",
	Long:  `Pings the DB and returns OK if the connection is ready.`,
	RunE:  ping,
}

func init() {

	RootCmd.AddCommand(pingCmd)
	RootCmd.PersistentFlags().StringP("pgpw", "p", "", "Postgres superuser password")
	RootCmd.PersistentFlags().StringP("configfile", "c", "config", "Name of config file (without extension)")
	RootCmd.PersistentFlags().BoolP("noprompt", "n", false, "Override prompt for confirmation")
	viper.BindPFlags(RootCmd.PersistentFlags())

}

// initConfig reads in config file and ENV variables if set.
func createConfigIfNotExists(cmd *cobra.Command, args []string) error {

	viper.SetConfigName(viper.GetString("configfile"))

	if err := viper.ReadInConfig(); err == nil {
		LogFatal(LogEntry{"ghost.CONFIG", true, "Config file already exists:" + viper.ConfigFileUsed()})
	} else {
		if err := createDefaultConfigFile(viper.GetString("configfile")); err != nil {
			LogFatal(LogEntry{"ghost.CONFIG", false, "Error creating config file: " + err.Error()})
		} else {
			//Otherwise create one
			Log(LogEntry{"ghost.CONFIG", true, "Config file created"})
		}
	}

	return nil
}

func ping(cmd *cobra.Command, args []string) error {

	Config.Setup(viper.GetString("configfile"))

	//Attempt to open a db connection
	db := SuperUserDBConfig.ReturnDBConnection("")
	defer db.Close()
	//IF we get this far, just exit with success
	Log(LogEntry{"PING", true, "Ping test passed"})
	os.Exit(0)

	return nil

}
