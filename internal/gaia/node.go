package gaia

// NodeList lists nodes in cluster
func (c *Client) NodeList() string {
	r := c.newRequest()
	return r.get("_nodes")
}

// NodeStats for this cluster
func (c *Client) NodeStats() string {
	r := c.newRequest()
	return r.get("_nodes/stats")
}

// NodeShow info about a specific node
func (c *Client) NodeShow(node string) string {
	r := c.newRequest()
	return r.get("_nodes/" + node)
}

// NodeSearchStats returns search stats for a node
func (c *Client) NodeSearchStats() string {
	r := c.newRequest()
	return r.get("_nodes/stats/indices/search")
}
