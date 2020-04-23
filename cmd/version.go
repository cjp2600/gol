package cmd

import (
	"fmt"
	"github.com/cjp2600/gol/core"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Gol-server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Set version flag
		gitsha := " (" + core.GitSHA + ")"
		if gitsha == " (0000000)" {
			gitsha = ""
		}
		versionLine := `gol-server version: ` + core.Version + gitsha
		fmt.Println(versionLine)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
