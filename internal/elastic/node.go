package elastic

// NodeList lists nodes in cluster
func NodeList() string {
	return Get("_nodes")
}

// NodeStats for this cluster
func NodeStats() string {
	return Get("_nodes/stats")
}

// NodeShow info about a specfic node
func NodeShow(node string) string {
	return Get("_nodes/" + node)
}

// NodeSearchStats returns search stats for a node
func NodeSearchStats() string {
	return Get("_nodes/stats/indices/search")
}
