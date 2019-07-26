package gaia

import "net/http"

// DocumentIndex indexes a document
func (c *Client) DocumentIndex(index string, id string, body string) string {
	r := c.newRequest()
	return r.put(index+"/"+index+"/"+id, body)
}

// DocumentGet get json document
func (c *Client) DocumentGet(index string, id string, source bool, fields string) string {
	r := c.newRequest()
	u := index + "/_all/" + id
	if source {
		u += "/_source"
	}
	if fields != "" {
		r.params["_source_include"] = fields
	}
	return r.get(u)
}

// DocumentDelete deletes a document
func (c *Client) DocumentDelete(index string, id string) string {
	r := c.newRequest()
	return r.delete(index + "/" + index + "/" + id)
}

// DocumentDeleteByQuery performs a query then deletes
// the results
func (c *Client) DocumentDeleteByQuery(index string, query string) string {
	r := c.newRequest()
	return r.post(index+"/_delete_by_query", query)
}

// DocumentUpdate performs update operation on a document
func (c *Client) DocumentUpdate(index string, id string, body string) string {
	r := c.newRequest()
	return r.post(index+"/"+index+"/"+id+"/_update", body)
}

// DocumentMultiGet allows to get multiple documents
// Annoyngly the mget api requires a GET request
// AND a request body so we need to call directly
// against the request instead of using the `get` wrapper
func (c *Client) DocumentMultiGet(index string, body string, fields string) string {
	r := c.newRequest()
	if fields != "" {
		r.params["_source_include"] = fields
	}
	r.method = http.MethodGet
	r.endpoint = index + "/_mget"
	r.body = body
	return r.do()
}

// DocumentReindex copies documents from one index to another
// !!! Reindex does not attempt to set up the destination index.
// It does not copy the settings of the source index. You should
// set up the destination index prior to running a _reindex action,
// including setting up mappings, shard counts, replicas, etc.
func (c *Client) DocumentReindex(body string) string {
	r := c.newRequest()
	return r.post("_reindex", body)
}
