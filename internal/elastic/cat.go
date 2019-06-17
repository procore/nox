package elastic

// Cat endpoint
func Cat(action string) string {
	urlParams["v"] = "true"
	r := Get("_cat/" + action)
	delete(urlParams, "v")
	return r
}

// CatCountIndex count documents for a specific index
func CatCountIndex(index string) string {
	return Cat("count/" + index)
}

// CatNodeType Cat nodes of a certain type
func CatNodeType(nodetype string) []string {
	resp := Cat("nodes")
	return grep(resp, nodetype)
}

// CatSnapshots list snapshots in a repo
func CatSnapshots(repo string) string {
	_, repo, _ = snapshotdefaultargs("", repo, "")
	return Cat("snapshots/" + repo)
}
