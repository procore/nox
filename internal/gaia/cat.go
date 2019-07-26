package gaia

// Cat endpoint
func (c *Client) Cat(action string) string {
	r := c.newRequest()
	r.params["v"] = "true"
	r.endpoint = "_cat/" + action
	return r.do()
}

// CatCountIndex count documents for a specific index
func (c *Client) CatCountIndex(index string) string {
	return c.Cat("count/" + index)
}

// CatNodeType Cat nodes of a certain type
func (c *Client) CatNodeType(nodetype string) []string {
	resp := c.Cat("nodes")
	return grep(resp, nodetype)
}

// CatSnapshots list snapshots in a repo
func (c *Client) CatSnapshots(repo string) string {
	return c.Cat("snapshots/" + repo)
}
