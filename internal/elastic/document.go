package elastic

import "net/http"

// DocumentIndex indexes a document
func DocumentIndex(index string, id string, body string) string {
	return Put(index+"/"+index+"/"+id, body)
}

// DocumentGet get json document
func DocumentGet(index string, id string, source bool, fields string) string {
	u := index + "/_all/" + id
	if source {
		u += "/_source"
	}
	if fields != "" {
		urlParams["_source_include"] = fields
	}
	r := Get(u)
	delete(urlParams, "_source_include")
	return r
}

// DocumentDelete deletes a document
func DocumentDelete(index string, id string) string {
	return Delete(index + "/" + index + "/" + id)
}

// DocumentDeleteByQuery performs a query then deletes
// the results
func DocumentDeleteByQuery(index string, query string) string {
	return Post(index+"/_delete_by_query", query)
}

// DocumentUpdate performs update operation on a document
func DocumentUpdate(index string, id string, body string) string {
	return Post(index+"/"+index+"/"+id+"/_update", body)
}

// DocumentMultiGet allows to get multiple documents
// Annoyngly the mget api requires a GET request
// AND a request body so we need to call directly
// against `request` instead of using the `Get` wrapper
func DocumentMultiGet(index string, body string, fields string) string {
	u := index + "/_mget"

	if fields != "" {
		urlParams["_source_include"] = fields
	}
	r := request(http.MethodGet, u, body)
	delete(urlParams, "_source_include")
	return r
}

// DocumentReindex copies documents from one index to another
// !!! Reindex does not attempt to set up the destination index.
// It does not copy the settings of the source index. You should
// set up the destination index prior to running a _reindex action,
// including setting up mappings, shard counts, replicas, etc.
func DocumentReindex(body string) string {
	return Post("_reindex", body)
}
