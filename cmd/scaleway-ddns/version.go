package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version = "undefined"
	commit  = "dev"
	date    = "undefined"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(fmt.Sprintf("scaleway-ddns %s (commit %s, built at %s)", version, commit, date))
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
