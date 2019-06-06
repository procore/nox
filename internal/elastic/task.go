package elastic

// GetTask retrieves currently executing tasks on a node
func GetTask(taskID string, nodes string) string {
	endpoint := "_tasks"
	if nodes != "" {
		urlParams["nodes"] = nodes
	}

	if taskID != "" {
		endpoint = endpoint + "/" + taskID + ":1"
	}
	r := Get(endpoint)
	delete(urlParams, "nodes")
	return r
}
