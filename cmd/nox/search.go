package main

import (
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search [index]",
	Short: "run a search",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		printResponse(client.Search(args[0], readFromFile(), nil))
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringVarP(&body, "body", "b", "", "json body to send with this request")
}
