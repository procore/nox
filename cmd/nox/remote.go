package main

import (
	"github.com/spf13/cobra"
)

// remoteCmd represents the node command
var remoteCmd = &cobra.Command{
	Use:   "remote",
	Short: "tool for managing remote clusters",
	Long:  `The cluster remote info API allows to retrieve all of the configured remote cluster information`,
	Run: func(cmd *cobra.Command, args []string) {
		response := client.RemoteInfo()
		printResponse(response)
	},
}

func init() {
	rootCmd.AddCommand(remoteCmd)
}
