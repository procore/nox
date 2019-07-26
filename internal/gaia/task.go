package gaia

// GetTask retrieves currently executing tasks on a node
// Args
//	- taskID string
//		id of the task to get
//  - nodes string
//		comma seperated list of nodes to query
func (c *Client) GetTask(taskID string, nodes string) string {
	r := c.newRequest()
	endpoint := "_tasks"
	if nodes != "" {
		r.params["nodes"] = nodes
	}

	if taskID != "" {
		endpoint = endpoint + "/" + taskID + ":1"
	}
	return r.get(endpoint)
}
