package gaia

import "strings"

// CountNodes returns number of nodes
func (c *Client) CountNodes() int {
	r := c.newRequest()
	nodes := r.get("_cat/nodes")
	narray := strings.Split(nodes, "\n")
	count := len(narray) - 1
	return count
}

// CountNodeType Cat nodes of a certain type
func (c *Client) CountNodeType(nodetype string) int {
	matches := c.CatNodeType(nodetype)
	return len(matches)
}
