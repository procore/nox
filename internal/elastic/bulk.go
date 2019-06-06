package elastic

import "strings"

// Bulk performs a bulk API request
func Bulk(body []string) string {
	headers["Content-Type"] = "application/x-ndjson"
	bulkBody := strings.Join(body, "\n")
	return Post("_bulk", bulkBody+"\n")
}
