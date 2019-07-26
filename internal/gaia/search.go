package gaia

import (
	"encoding/json"
	"log"
	"net/http"
)

type countResult struct {
	Count int
}

// Search runs the search provided on the index provided
func (c *Client) Search(index string, body string, options map[string]string) string {
	r := c.newRequest()
	if options["scroll"] != "" {
		r.params["scroll"] = options["scroll"]
	}
	return r.post(index+"/_search", body)
}

// Scroll gets the next page of results for a search
func (c *Client) Scroll(scrollID string, scroll string) string {
	r := c.newRequest()
	body := `
{
	"scroll": "` + scroll + `",
	"scroll_id": "` + scrollID + `"
}
`
	return r.post("_search/scroll", body)
}

// DeleteScroll clears the scroll context when
// you no longer need it
func (c *Client) DeleteScroll(scrollID string) string {
	r := c.newRequest()
	return r.delete("_search/scroll/" + scrollID)
}

// Count returns the number of results from a query
func (c *Client) Count(index string, body string) int {
	r := request{
		client:   c,
		method:   http.MethodGet,
		endpoint: index + "/_count",
		body:     body,
	}
	var result countResult
	resp := []byte(r.do())
	err := json.Unmarshal(resp, &result)
	if err != nil {
		log.Fatal(err)
	}
	return result.Count
}
