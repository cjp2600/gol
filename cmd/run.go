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
	"context"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
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
	runCmd.Flags().StringP("http-port", "p", "8012", "http server port")
	runCmd.Flags().StringP("grpc-port", "g", "6012", "tcp server port")
	runCmd.Flags().BoolP("verbose", "v", false, "verbose output")
}

func RunCmd(cmd *cobra.Command, args []string) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("service", "gol").Logger()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	httpPort, err := cmd.Flags().GetString("http-port")
	grpcPort, err := cmd.Flags().GetString("grpc-port")
	verbose, err := cmd.Flags().GetBool("verbose")

	if err != nil {
		logger.Info().Msgf("flag parse error %s", err)
	}

	if verbose {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	server := NewServer(logger, httpPort, grpcPort)

	go func() {
		err := server.RunGRPC()
		log.Fatal().Msgf("error %s", err)
	}()

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
		<-stop
		logger.Info().Msg("Stop gol application")

		cancel()
		os.Exit(1)
	}()

	if err := server.RunHttp(ctx); err != nil {
		log.Fatal().Msgf("error %s", err)
	}
}
