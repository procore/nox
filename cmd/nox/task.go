package main

import (
	"strings"

	"github.com/spf13/cobra"
)

var nodes string

// taskCmd represents the node command
var taskCmd = &cobra.Command{
	Use:   "tasks",
	Short: "information about cluster tasks",
	Long: `The task management API allows to retrieve
information about the tasks currently executing on one or more nodes in the cluster`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		response := client.GetTask(strings.Join(args, ""), nodes)
		printResponse(response)
	},
}

func init() {
	rootCmd.AddCommand(taskCmd)
	taskCmd.Flags().StringVarP(&nodes, "nodes", "n", "", "comma separated list of nodes to run operation on")
}
