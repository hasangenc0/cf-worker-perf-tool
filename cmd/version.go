package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	version = "1.0.0"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version number of this cli tool",
	Run:   ShowVersion,
}

func ShowVersion(cmd *cobra.Command, args []string) {
	showVersion()
}

func showVersion() {
	fmt.Printf("cf-worker-perf-tool version: %s \n", version)
}
