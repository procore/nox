package main

import (
	"github.com/spf13/cobra"
)

var indexCmd = &cobra.Command{
	Use:   "index",
	Short: "tool for managing an index",
}

var toggleOpen = &cobra.Command{
	Use:   "open [index_name]",
	Short: "Open a closed index",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		response := client.IndexToggleOpen(cmd.Name(), args[0])
		printResponse(response)
	},
}

var toggleClose = &cobra.Command{
	Use:   "close [index_name]",
	Short: "Close an open index",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		response := client.IndexToggleOpen(cmd.Name(), args[0])
		printResponse(response)
	},
}

var indexDelete = &cobra.Command{
	Use:   "delete [name]",
	Short: "Delete an index. Use with Caution",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if confirm() {
			response := client.IndexDelete(args[0])
			printResponse(response)
		}
	},
}

func init() {
	rootCmd.AddCommand(indexCmd)
	indexCmd.AddCommand(toggleOpen, toggleClose, indexDelete)
	indexDelete.Flags().BoolVarP(&override, "confirm", "y", false, "Confirm that you want to delete an index")
}
