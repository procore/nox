package main

import (
	"github.com/procore/nox/internal/elastic"
	"github.com/spf13/cobra"
)

// remoteCmd represents the node command
var remoteCmd = &cobra.Command{
	Use:   "remote",
	Short: "tool for managing remote clusters",
	Long:  `The cluster remote info API allows to retrieve all of the configured remote cluster information`,
	Run: func(cmd *cobra.Command, args []string) {
		response := elastic.RemoteInfo()
		printResponse(response)
	},
}

func init() {
	rootCmd.AddCommand(remoteCmd)
}
