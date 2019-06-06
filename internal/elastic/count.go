package elastic

import "strings"

// CountNodes returns number of nodes
func CountNodes() int {
	nodes := Get("_cat/nodes")
	narray := strings.Split(nodes, "\n")
	count := len(narray) - 1
	return count
}

// CountNodeType Cat nodes of a certain type
func CountNodeType(nodetype string) int {
	matches := CatNodeType(nodetype)
	return len(matches)
}
