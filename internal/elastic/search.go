package elastic

import (
	"encoding/json"
	"log"
	"net/http"
)

type countResult struct {
	Count int
}

// Search runs the search provided on the index provided
func Search(index string, body string, options map[string]string) string {
	if options["scroll"] != "" {
		urlParams["scroll"] = options["scroll"]
	}
	r := Post(index+"/_search", body)
	delete(urlParams, "scroll")
	return r
}

// Scroll gets the next page of results for a search
func Scroll(scrollID string, scroll string) string {
	body := `
{
	"scroll": "` + scroll + `",
	"scroll_id": "` + scrollID + `"
}
`
	return Post("_search/scroll", body)
}

// DeleteScroll clears the scroll context when
// you no longer need it
func DeleteScroll(scrollID string) string {
	return Delete("_search/scroll/" + scrollID)
}

// Count returns the number of results from a query
func Count(index string, body string) int {
	var r countResult
	resp := []byte(request(http.MethodGet, index+"/_count", body))
	err := json.Unmarshal(resp, &r)
	if err != nil {
		log.Fatal(err)
	}
	return r.Count
}
