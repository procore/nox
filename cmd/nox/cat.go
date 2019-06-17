package main

import (
	"strings"

	"github.com/procore/nox/internal/elastic"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// catCmd represents the cat command
var catCmd = &cobra.Command{
	Use:   "cat",
	Short: "Query the cat apis",
	Long: `JSON is great… for computers. Even if it’s pretty-printed,
trying to find relationships in the data is tedious. Human eyes,
especially when looking at an ssh terminal, need compact and aligned text.
The cat API aims to meet this need.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		viper.Set("pretty", false)
		configESClient()
	},
}

var catAliases = &cobra.Command{
	Use:   "aliases",
	Short: "list configured aliases",
	Long:  `Aliases shows information about currently configured aliases to indices including filter and routing infos.`,
	Run: func(cmd *cobra.Command, args []string) {
		response := elastic.Cat(cmd.Name())
		printResponse(response)
	},
}

var catAllocation = &cobra.Command{
	Use:   "allocation",
	Short: "display shard allocation",
	Long:  `Allocation provides a snapshot of how many shards are allocated to each data node and how much disk space they are using.`,
	Run: func(cmd *cobra.Command, args []string) {
		response := elastic.Cat(cmd.Name())
		printResponse(response)
	},
}

var catCount = &cobra.Command{
	Use:   "count [index]",
	Short: "get document counts",
	Long:  `Count provides quick access to the document count of the entire cluster, or individual indices`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var response string
		if len(args) > 0 {
			response = elastic.CatCountIndex(args[0])
		} else {
			response = elastic.Cat(cmd.Name())
		}
		printResponse(response)
	},
}

var catFielddata = &cobra.Command{
	Use:   "fielddata",
	Short: "info on fieldata heap usage",
	Long:  `Fielddata shows how much heap memory is currently being used by fielddata on every data node in the cluster.`,
	Run: func(cmd *cobra.Command, args []string) {
		response := elastic.Cat(cmd.Name())
		printResponse(response)
	},
}

var catHealth = &cobra.Command{
	Use:   "health",
	Short: "info on cluster health",
	Long:  `Health is a terse, one-line representation of the same information from ` + "`cluster health`",
	Run: func(cmd *cobra.Command, args []string) {
		response := elastic.Cat(cmd.Name())
		printResponse(response)
	},
}

var catIndices = &cobra.Command{
	Use:   "indices",
	Short: "summary of all indices",
	Long:  `The indices command provides a cross-section of each index. This information spans nodes.`,
	Run: func(cmd *cobra.Command, args []string) {
		response := elastic.Cat(cmd.Name())
		printResponse(response)
	},
}

var catMaster = &cobra.Command{
	Use:   "master",
	Short: "master node summary",
	Long:  `Displays the master’s node ID, bound IP address, and node name.`,
	Run: func(cmd *cobra.Command, args []string) {
		response := elastic.Cat(cmd.Name())
		printResponse(response)
	},
}

var catNodeattrs = &cobra.Command{
	Use:   "nodeattrs",
	Short: "Shows custom node attributes",
	Run: func(cmd *cobra.Command, args []string) {
		response := elastic.Cat(cmd.Name())
		printResponse(response)
	},
}

var catNodes = &cobra.Command{
	Use:   "nodes",
	Short: "Shows the cluster topology",
	Run: func(cmd *cobra.Command, args []string) {
		response := elastic.Cat(cmd.Name())
		printResponse(response)
	},
}

var catNodeType = &cobra.Command{
	Use:   "type [node-type]",
	Short: "Get info on only a certain type of node",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		matches := elastic.CatNodeType(args[0])
		response := strings.Join(matches, "\n")
		printResponse(response)
	},
}

var catPendingTasks = &cobra.Command{
	Use:   "pending_tasks",
	Short: "List pending tasks",
	Run: func(cmd *cobra.Command, args []string) {
		response := elastic.Cat(cmd.Name())
		printResponse(response)
	},
}

var catPlugins = &cobra.Command{
	Use:   "plugins",
	Short: "list plugins",
	Long:  `The plugins command provides a view per node of running plugins. This information spans nodes`,
	Run: func(cmd *cobra.Command, args []string) {
		response := elastic.Cat(cmd.Name())
		printResponse(response)
	},
}

var catRecovery = &cobra.Command{
	Use:   "recovery",
	Short: "compact view of the JSON recovery API",
	Long: `The recovery command is a view of index shard recoveries, both on-going and previously completed.
A recovery event occurs anytime an index shard moves to a different node in the cluster.
This can happen during a snapshot recovery, a change in replication level, node failure, or on node startup.
This last type is called a local store recovery and is the normal way for shards to be loaded from disk when a node starts up.`,
	Run: func(cmd *cobra.Command, args []string) {
		response := elastic.Cat(cmd.Name())
		printResponse(response)
	},
}

var catRepositories = &cobra.Command{
	Use:   "repositories",
	Short: "Shows the snapshot repositories registered in the cluster",
	Run: func(cmd *cobra.Command, args []string) {
		response := elastic.Cat(cmd.Name())
		printResponse(response)
	},
}

var catThreadPool = &cobra.Command{
	Use:   "thread_pool",
	Short: "thread pool info",
	Long:  `Shows cluster wide thread pool statistics per node`,
	Run: func(cmd *cobra.Command, args []string) {
		response := elastic.Cat(cmd.Name())
		printResponse(response)
	},
}

var catShards = &cobra.Command{
	Use:   "shards",
	Short: "shard info",
	Long: `The shards command is the detailed view of what nodes contain which shards.
It will tell you if it’s a primary or replica, the number of docs, the bytes it takes on disk, and the node where it’s located.`,
	Run: func(cmd *cobra.Command, args []string) {
		response := elastic.Cat(cmd.Name())
		printResponse(response)
	},
}

var catSegments = &cobra.Command{
	Use:   "segments",
	Short: "segment info",
	Long:  `Provides low level information about the segments in the shards of an index.`,
	Run: func(cmd *cobra.Command, args []string) {
		response := elastic.Cat(cmd.Name())
		printResponse(response)
	},
}

var catSnapshots = &cobra.Command{
	Use:   "snapshots [repository]",
	Short: "List snapshots in a repo",
	Long: `Shows all snapshots that belong to a specific repository.
To find a list of available repositories to query, the command ` + "`cat repositories`" + ` can be used`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		response := elastic.CatSnapshots(strings.Join(args, ""))
		printResponse(response)
	},
}

var catTemplates = &cobra.Command{
	Use:   "templates",
	Short: "Provides information about existing templates",
	Run: func(cmd *cobra.Command, args []string) {
		response := elastic.Cat(cmd.Name())
		printResponse(response)
	},
}

func init() {
	rootCmd.AddCommand(catCmd)
	catCmd.AddCommand(catAliases, catAllocation,
		catCount, catFielddata, catHealth, catIndices,
		catMaster, catNodeattrs, catNodes,
		catPendingTasks, catPlugins, catRecovery,
		catRepositories, catThreadPool, catShards,
		catSegments, catSnapshots, catTemplates)
	catNodes.AddCommand(catNodeType)
}
