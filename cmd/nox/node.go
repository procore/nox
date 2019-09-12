package main

import (
	"strconv"

	"github.com/spf13/cobra"
)

var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "tool for managing nodes",
}

var nodeList = &cobra.Command{
	Use:   "list",
	Short: "List all nodes",
	Run: func(cmd *cobra.Command, args []string) {
		response := client.NodeList()
		printResponse(response)
	},
}

var nodeStats = &cobra.Command{
	Use:   "stats",
	Short: "node stats",
	Long:  `The cluster nodes stats API allows to retrieve one or more (or all) of the cluster nodes statistics`,
	Run: func(cmd *cobra.Command, args []string) {
		response := client.NodeStats()
		printResponse(response)
	},
}

var nodeShow = &cobra.Command{
	Use:   "get [node_id]",
	Short: "Show info for a specific node",
	Long:  `The cluster nodes info API allows to retrieve one or more (or all) of the cluster nodes information.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		value := args[0]
		response := client.NodeShow(value)
		printResponse(response)
	},
}

var nodeSearchStats = &cobra.Command{
	Use:   "search [node_id]",
	Short: "Get stats associated with search running on nodes",
	Run: func(cmd *cobra.Command, args []string) {
		printResponse(client.NodeSearchStats())
	},
}

var countNodes = &cobra.Command{
	Use:   "count",
	Short: "Count nodes",
	Run: func(cmd *cobra.Command, args []string) {
		response := client.CountNodes()
		printResponse(strconv.Itoa(response))
	},
}

var countNodeType = &cobra.Command{
	Use:   "type [node-type]",
	Short: "Count nodes of a certain type",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		response := client.CountNodeType(args[0])
		printResponse(strconv.Itoa(response))
	},
}

func init() {
	rootCmd.AddCommand(nodeCmd)
	nodeCmd.AddCommand(countNodes, nodeList, nodeStats, nodeShow)
	countNodes.AddCommand(countNodeType)
	nodeStats.AddCommand(nodeSearchStats)
}
