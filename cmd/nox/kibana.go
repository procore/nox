package main

import (
	"log"
	"strings"

	"github.com/procore/nox/internal/elastic"
	"github.com/softpunks/goin"
	"github.com/spf13/cobra"
)

var commandFile string

// kibanaCmd represents the kibana command
var kibanaCmd = &cobra.Command{
	Use:     "kibana",
	Aliases: []string{"k"},
	Short:   "Run a kibana style query",
	Long: `The kibana command opens an editor that allows you to enter any query that is
compatible with the kibana dev tools console. The first line must be the HTTP keyword
and the endpoind. The second line starts your JSON request body. Example:

GET <index_name>/_search/
{
  "query": {
    "bool": {
      "must": [
        {
          "multi_match": {
            "query": "<query>",
          }
        }
      ],
    }
  }
}`,
	Run: func(cmd *cobra.Command, args []string) {
		var fileName string
		if commandFile == "" {
			fileName = tempBufferFileName
		} else {
			fileName = commandFile
		}
		// TODO: this is currently deleting the specified file
		// this is non ideal since the whole point of the option
		// is to have repeatable kibana commands
		a, err := goin.ReadLinesFromFile(fileName)
		if err != nil {
			log.Fatal(err)
		}
		c, a := a[0], a[1:]

		s := strings.Join(a, "")
		printResponse(elastic.KibanaQuery(c, s))
	},
}

func init() {
	rootCmd.AddCommand(kibanaCmd)
	kibanaCmd.Flags().StringVarP(&commandFile, "file", "f", "", "file containing kibana command")
}
