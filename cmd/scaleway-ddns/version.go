package main

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	version       = "undefined"
	commit        = "dev"
	date          = "undefined"
	golangVersion = runtime.Version()
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("scaleway-ddns %s (%s, commit %s, built at %s)\n", version, golangVersion, commit, date)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
