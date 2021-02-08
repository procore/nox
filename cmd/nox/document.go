package main

import (
	"github.com/spf13/cobra"
)

var fields string
var source bool

var documentCMD = &cobra.Command{
	Use:     "document",
	Aliases: []string{"doc"},
	Short:   "tool for managing documents",
}

var documentIndex = &cobra.Command{
	Use:   "index [index] [id]",
	Short: "Index a document",
	Long:  "update a typed JSON document in a specific index, making it searchable",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		printResponse(client.DocumentIndex(args[0], args[1], readFromFile()))
	},
}

var documentGet = &cobra.Command{
	Use:   "get [index] [id]",
	Short: "Get a document",
	Long:  "The get API allows to get a typed JSON document from the index based on its id.",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		printResponse(client.DocumentGet(args[0], args[1], source, fields))
	},
}

var documentDelete = &cobra.Command{
	Use:   "delete [index] [id]",
	Short: "Delete a document",
	Long:  "The delete API allows to delete a typed JSON document from a specific index based on its id",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if confirm() {
			printResponse(client.DocumentDelete(args[0], args[1]))
		}
	},
}

var documentDeleteByQuery = &cobra.Command{
	Use:   "delete_by_query [index]",
	Short: "Delete the documents that are the result of a query",
	Long:  "The simplest usage of delete_by_query just performs a deletion on every document that match a query",
	Run: func(cmd *cobra.Command, args []string) {
		if confirm() {
			printResponse(client.DocumentDeleteByQuery(args[0], readFromFile()))
		}
	},
}

var documentUpdate = &cobra.Command{
	Use:   "update [index] [id]",
	Short: "Update a document",
	Long: `The update API allows to update a document based on a script provided.
The operation gets the document (collocated with the shard) from the index,
runs the script (with optional script language and parameters), and index
back the result (also allows to delete, or ignore the operation).`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		printResponse(client.DocumentUpdate(args[0], args[1], readFromFile()))
	},
}

var documentMultiGet = &cobra.Command{
	Use:   "mget [index]",
	Short: "Get multiple documents",
	Long: `Multi GET API allows to get multiple
documents based on an index, type (optional) and
id (and possibly routing). The response includes a docs
array with all the fetched documents in order corresponding to the original
multi-get request (if there was a failure for a specific get,
an object containing this error is included in place in the response instead).
The structure of a successful get is similar in structure to a document provided by the get API.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		printResponse(client.DocumentMultiGet(args[0], readFromFile(), fields))
	},
}

var documentReindex = &cobra.Command{
	Use:   "reindex",
	Short: "Copy documents from one index to another",
	Long: `Reindex does not attempt to set up the destination index.
It does not copy the settings of the source index. You should set
up the destination index prior to running a _reindex action, including
setting up mappings, shard counts, replicas, etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		if confirm() {
			printResponse(client.DocumentReindex(readFromFile()))
		}
	},
}

func init() {
	rootCmd.AddCommand(documentCMD)
	documentCMD.AddCommand(documentGet, documentIndex, documentDelete,
		documentDeleteByQuery, documentMultiGet, documentReindex, documentUpdate)

	documentDeleteByQuery.Flags().StringVarP(&body, "body", "b", "", "json body to send with this request")
	documentUpdate.Flags().StringVarP(&body, "body", "b", "", "json body to send with this request")
	documentIndex.Flags().StringVarP(&body, "body", "b", "", "json body to send with this request")
	documentMultiGet.Flags().StringVarP(&body, "body", "b", "", "json body to send with this request")
	documentGet.Flags().StringVarP(&fields, "fields", "f", "", "comma separated list of fields to retrieve")
	documentGet.Flags().BoolVarP(&source, "source", "s", false, "retrieve only the _source object")
	documentMultiGet.Flags().StringVarP(&fields, "fields", "f", "", "comma separated list of fields to retrieve")
	documentDelete.Flags().BoolVarP(&override, "confirm", "y", false, "Confirm that you want to delete a document")
	documentDeleteByQuery.Flags().BoolVarP(&override, "confirm", "y", false, "Confirm that you want to delete a document")
}
