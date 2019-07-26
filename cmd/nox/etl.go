package main

import (
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/Jeffail/gabs"
	"github.com/procore/nox/cmd/nox/internal/algorithms"
	"github.com/spf13/cobra"
	"gopkg.in/cheggaaa/pb.v1"
)

var maxSlices int

var sourceIndex string
var destinationIndex string
var keepAlive string
var sourceFields string
var batchSize int
var dryRun bool

var searchOptions map[string]string

var wg sync.WaitGroup
var bar *pb.ProgressBar

var etlCmd = &cobra.Command{
	Use:   "etl [source_index]",
	Short: "Run a manual reindex with an extra data transform step",
	Long: `This command essentially performs a manual document reindex on an Elasticsearch cluster.
What makes this different is that you can perform a data transform step before the data is saved to the index.
This allows for Elasticsearch data ETL between indicies, and to new indices.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		searchOptions = make(map[string]string)
		searchOptions["scroll"] = keepAlive

		sourceIndex = args[0]
		if destinationIndex == "" {
			destinationIndex = sourceIndex
		}

		total := client.Count(sourceIndex, `{"query": { "match_all": {} } }`)
		bar = pb.StartNew(total)
		bar.ShowTimeLeft = true
		bar.ShowSpeed = true

		for i := 0; i < maxSlices; i++ {
			wg.Add(1)
			go readScroll(i)
		}

		wg.Wait()
		bar.FinishPrint("Done!")
	},
}

func init() {
	rootCmd.AddCommand(etlCmd)
	etlCmd.Flags().StringVar(&destinationIndex, "destination", "", "destination index name, defaults to source index")
	etlCmd.Flags().StringVar(&keepAlive, "scroll", "1m", "how long to keep the scroll context alive")
	etlCmd.Flags().StringVar(&sourceFields, "fields", "", "comma separated list of fields you want to retrieve from the _source object on search")
	etlCmd.Flags().IntVar(&batchSize, "batch", 1000, "size of thethe batches of documents per scroll")

	etlCmd.Flags().IntVar(&maxSlices, "threads", 2, "number of slices/threads to split the scroll into")
	etlCmd.Flags().BoolVar(&dryRun, "dry-run", false, "toggle dry run for etl command")
}

func readScroll(i int) {
	var page *gabs.Container
	var sid string
	var scrollSize int

	defer wg.Done()

	body := `{ `
	if sourceFields != "" {
		s := strings.Split(sourceFields, ",")
		body += `"_source": [ `
		for i, v := range s {
			body += `"` + v + `"`
			if i != len(s)-1 {
				body += `, `
			}
		}
		body += ` ], `
	}
	body += `"size": ` + strconv.Itoa(batchSize) + `, "slice": { "id": ` + strconv.Itoa(i) + `, "max": ` + strconv.Itoa(maxSlices) + ` }, `
	body += `"query": { "match_all": {} },`
	body += `"sort": [ "_doc"] }`

	page = parseResponse(client.Search(sourceIndex, body, searchOptions))
	sid = page.Path("_scroll_id").Data().(string)
	hits, err := page.Path("hits.hits").Children()
	if err != nil {
		log.Fatal(err)
	}
	updateDoc(hits)
	bar.Add(len(hits))
	scrollSize = int(page.Path("hits.total").Data().(float64)) - batchSize
	for scrollSize > 0 {
		page = parseResponse(client.Scroll(sid, keepAlive))
		sid = page.Path("_scroll_id").Data().(string)
		hits, err := page.Path("hits.hits").Children()
		if err != nil {
			log.Fatal(err)
		}
		updateDoc(hits)
		bar.Add(len(hits))
		scrollSize -= len(hits)
	}
	parseResponse(client.DeleteScroll(sid))
}

// TODO: Allow plugins for user defined etl functions
// TODO: reflection to allow choosing function from cli option
func updateDoc(h []*gabs.Container) {
	var bulkBody []string
	for _, hit := range h {
		doc := hit.Path("_source")
		if doc, ok := algorithms.TextStripBackSlash(doc); ok {
			docString := doc.String()
			bulkBody = append(bulkBody, `{ "update": { "_index": "`+destinationIndex+`", "_type": "`+hit.Path("_type").Data().(string)+`", "_id": "`+hit.Path("_id").Data().(string)+`", "_retry_on_conflict": 5}}
{"doc": `+string(docString)+`, "doc_as_upsert": true}`)
		}
	}
	if len(bulkBody) > 0 && !dryRun {
		parseResponse(client.Bulk(bulkBody))
	}
}

func parseResponse(response string) *gabs.Container {
	pageBits := []byte(response)
	jsonParsed, err := gabs.ParseJSON(pageBits)
	if err != nil {
		log.Fatal(err)
	}

	if exists := jsonParsed.Exists("error"); exists {
		log.Fatal(jsonParsed.StringIndent("", " "))
	}
	return jsonParsed
}
