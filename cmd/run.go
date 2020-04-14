/*
Copyright © 2020 NAME HERE icjp2600@gmail.com

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
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"os"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run http server",
	Long:  `run http server`,
	Run:   RunCmd,
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringP("port", "p", "", "http server port")
	runCmd.Flags().BoolP("verbose", "v", false, "verbose debug information")
}

func RunCmd(cmd *cobra.Command, args []string) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("service", "gol").Logger()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	port, _ := cmd.Flags().GetString("port")
	verbose, _ := cmd.Flags().GetBool("verbose")
	if verbose {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	server := NewRest(logger, port)
	if err := server.Run(); err != nil {
		logger.Err(err)
	}
}
