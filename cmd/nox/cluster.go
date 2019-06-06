package main

import (
	"strings"

	"github.com/procore/nox/internal/elastic"
	"github.com/spf13/cobra"
)

var flatSettings bool

// clusterCmd represents the cluster command
var clusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "tool for managing clusters",
}

// clusterCmd represents the cluster command
var clusterInfo = &cobra.Command{
	Use:   "info",
	Short: "tool for managing clusters",
	Run: func(cmd *cobra.Command, args []string) {
		response := elastic.ClusterInfo()
		printResponse(response)
	},
}

var clusterHealth = &cobra.Command{
	Use:   "health",
	Short: "view cluster health",
	Long:  `The cluster health API allows to get a very simple status on the health of the cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		response := elastic.ClusterHealth()
		printResponse(response)
	},
}

var clusterState = &cobra.Command{
	Use:   "state",
	Short: "view cluster state",
	Long:  `The cluster state API allows to get a comprehensive state information of the whole cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		response := elastic.ClusterState()
		printResponse(response)
	},
}

var clusterStats = &cobra.Command{
	Use:   "stats",
	Short: "view cluster stats",
	Long: `The Cluster Stats API allows to retrieve statistics from a cluster wide perspective.
The API returns basic index metrics (shard numbers, store size, memory usage) and information
about the current nodes that form the cluster (number, roles, os, jvm versions, memory usage, cpu and installed plugins).`,
	Run: func(cmd *cobra.Command, args []string) {
		response := elastic.ClusterStats()
		printResponse(response)
	},
}

var clusterPendingTasks = &cobra.Command{
	Use:   "tasks",
	Short: "view pending tasks",
	Long: `The pending cluster tasks API returns a list of any cluster-level changes
(e.g. create index, update mapping, allocate or fail shard) which have not yet been executed.`,
	Run: func(cmd *cobra.Command, args []string) {
		response := elastic.ClusterPendingTasks()
		printResponse(response)
	},
}

var clusterReroute = &cobra.Command{
	Use:   "reroute",
	Short: "run an explicit cluster reroute",
	Long: `The reroute command allows to explicitly execute a cluster reroute
allocation command including specific commands. For example, a shard can be
moved from one node to another explicitly, an allocation can be canceled,
or an unassigned shard can be explicitly allocated on a specific node.`,
	Run: func(cmd *cobra.Command, args []string) {
		response := elastic.ClusterReroute(readFromFile())
		printResponse(response)
	},
}

var settingsCmd = &cobra.Command{
	Use:   "settings",
	Short: "view cluster wide settings",
	Run: func(cmd *cobra.Command, args []string) {
		response := elastic.ClusterSettings(flatSettings)
		printResponse(response)
	},
}

var clusterUpdate = &cobra.Command{
	Use:   "update",
	Short: "update aspects of cluster",
}

var settingsUpdate = &cobra.Command{
	Use:   "settings",
	Short: "update cluster wide settings",
	Long: `Allows to update cluster wide specific settings.
Settings updated can either be persistent (applied across restarts)
or transient (will not survive a full cluster restart).`,
	Run: func(cmd *cobra.Command, args []string) {
		response := elastic.ClusterUpdateSettings(readFromFile(), flatSettings)
		printResponse(response)
	},
}

var toggleRouting = &cobra.Command{
	Use:       "routing [setting]",
	Short:     "toggle setting for routing allocation",
	ValidArgs: []string{"all", "none", "primaries", "new_primaries"},
	Args:      cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		value := strings.ToLower(args[0])
		response := elastic.ToggleRouting(value, flatSettings)
		printResponse(response)
	},
}

func init() {
	rootCmd.AddCommand(clusterCmd)
	clusterCmd.AddCommand(clusterUpdate)
	clusterUpdate.AddCommand(settingsUpdate, toggleRouting)
	clusterCmd.AddCommand(clusterHealth, clusterState, clusterStats, settingsCmd,
		clusterPendingTasks, clusterReroute, clusterInfo)

	clusterUpdate.PersistentFlags().BoolVarP(&flatSettings, "flat", "f", false, "toggle settings are returned in a flat format")

	clusterReroute.Flags().StringVarP(&body, "body", "b", "", "json body to send with this request")
	settingsUpdate.Flags().StringVarP(&body, "body", "b", "", "json body to send with this request")

}
