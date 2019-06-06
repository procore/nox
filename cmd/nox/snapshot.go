package main

import (
	"strings"

	"github.com/procore/nox/internal/elastic"
	"github.com/spf13/cobra"
)

var r string
var f string
var cleanNumber int

var snapshotCmd = &cobra.Command{
	Use:   "snapshot",
	Short: "tool for managing snapshots",
}

var snapshotRepoRegister = &cobra.Command{
	Use:   "register [name]",
	Short: "Register a new snapshot repository",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		response := elastic.SnapshotRepoRegister(strings.Join(args, ""), readFromFile())
		printResponse(response)
	},
}

var snapshotStart = &cobra.Command{
	Use:   "start [name]",
	Short: "Kick off a snapshot",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		response := elastic.SnapshotStart(strings.Join(args, ""), r, f, readFromFile())
		printResponse(response)
	},
}

var snapshotList = &cobra.Command{
	Use:   "list",
	Short: "List your available snapshots",
	Run: func(cmd *cobra.Command, args []string) {
		response := elastic.SnapshotList(r)
		printResponse(response)
	},
}

var snapshotGet = &cobra.Command{
	Use:   "get [name]",
	Short: "Get details about a snapshot",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		response := elastic.SnapshotGet(strings.Join(args, ""), r, f)
		printResponse(response)
	},
}

var snapshotRestore = &cobra.Command{
	Use:   "restore [name]",
	Short: "Restore a snapshot",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		response := elastic.SnapshotRestore(strings.Join(args, ""), r, f, readFromFile())
		printResponse(response)
	},
}

var snapshotDelete = &cobra.Command{
	Use:   "delete [name]",
	Short: "Delete a snapshot",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if confirm() {
			response := elastic.SnapshotDelete(strings.Join(args, ""), r, f)
			printResponse(response)
		}
	},
}

var snapshotClean = &cobra.Command{
	Use:   "clean [# number]",
	Short: "Delete all but a certain number of snapshots",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if confirm() {
			elastic.SnapshotClean(cleanNumber, r, f)
		}
	},
}

func init() {
	rootCmd.AddCommand(snapshotCmd)
	snapshotCmd.AddCommand(snapshotRepoRegister, snapshotStart, snapshotRestore, snapshotClean, snapshotList,
		snapshotGet, snapshotDelete)

	snapshotCmd.PersistentFlags().StringVarP(&f, "frequency", "f", "", "Frequency of the snapshot")
	snapshotCmd.PersistentFlags().StringVarP(&r, "repo", "r", "", "Repository for the snapshot")

	snapshotList.Flags().StringVarP(&f, "frequency", "f", "", "Frequency of the snapshot")
	snapshotList.Flags().StringVarP(&r, "repo", "r", "", "Repository for the snapshot")

	snapshotGet.Flags().StringVarP(&f, "frequency", "f", "", "Frequency of the snapshot")
	snapshotGet.Flags().StringVarP(&r, "repo", "r", "", "Repository for the snapshot")

	snapshotDelete.Flags().StringVarP(&f, "frequency", "f", "", "Frequency of the snapshot")
	snapshotDelete.Flags().StringVarP(&r, "repo", "r", "", "Repository for the snapshot")

	snapshotStart.Flags().StringVarP(&body, "body", "b", "", "json body to send with this request")
	snapshotRestore.Flags().StringVarP(&body, "body", "b", "", "json body to send with this request")
	snapshotDelete.Flags().BoolVarP(&override, "confirm", "y", false, "Confirm that you want to destroy a snapshot")
	snapshotClean.Flags().BoolVarP(&override, "confirm", "y", false, "Confirm that you want to destroy a snapshot")
	snapshotClean.Flags().IntVarP(&cleanNumber, "number", "n", 3, "number of snapshots to keep")
	snapshotRepoRegister.Flags().StringVarP(&body, "body", "b", "", "json body to send with this request")
}
