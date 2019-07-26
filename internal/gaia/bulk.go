package gaia

import (
	"strings"
)

// Bulk performs a bulk API request
// Parameters:
// 	- body []string
//		a slice of bulk operations
// Example:
//  var body []string
//  body = append("{'update': {...}}", body)
//  body = append("{'doc': {...}}", body)
//  client.Bulk(body)
func (c *Client) Bulk(body []string) string {
	r := c.newRequest()
	r.headers["Content-Type"] = "application/x-ndjson"
	bulkBody := strings.Join(body, "\n")
	return r.post("_bulk", bulkBody)
}
