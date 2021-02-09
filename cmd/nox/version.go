package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "the current nox version",
	Long:  `the current git tag for nox releases`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(Version)
	},
}
