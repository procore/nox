package elastic

import (
	"fmt"
	"strconv"
)

// ClusterInfo returns general data about your cluster
func ClusterInfo() string {
	return Get("")
}

// ClusterHealth returns the general health of the cluster
func ClusterHealth() string {
	return Get("_cluster/health")
}

// ClusterState returns the general health of the cluster
func ClusterState() string {
	return Get("_cluster/state")
}

// ClusterStats Get some sweet sweet stats
func ClusterStats() string {
	urlParams["human"] = "true"
	r := Get("_cluster/stats")
	delete(urlParams, "human")
	return r
}

// ClusterPendingTasks returns list of cluster level
// changes that have not been executedut
func ClusterPendingTasks() string {
	return Get("_cluster/pending_tasks")
}

// ClusterReroute explicitly executes a
// cluster reroute allocation command
func ClusterReroute(body string) string {
	return Post("_cluster/reroute", body)
}

// ClusterSettings returns current cluster level settings
func ClusterSettings(flatSettings bool) string {
	urlParams["flat_settings"] = strconv.FormatBool(flatSettings)
	r := Get("_cluster/settings")
	delete(urlParams, "flat_settings")
	return r
}

// ClusterUpdateSettings allows updating cluster wide settings
func ClusterUpdateSettings(body string, flatSettings bool) string {
	urlParams["flat_settings"] = strconv.FormatBool(flatSettings)
	r := Put("_cluster/settings", body)
	delete(urlParams, "flat_settings")
	return r
}

// ToggleRouting turns off dynaamic allocation
func ToggleRouting(setting string, flatSettings bool) string {
	urlParams["flat_settings"] = strconv.FormatBool(flatSettings)
	t := fmt.Sprintf("{\"transient\": {\"cluster.routing.allocation.enable\": \"%s\"},"+
		"\"persistent\": {\"cluster.routing.allocation.enable\": \"%s\"}}", setting, setting)
	r := Put("_cluster/settings", t)
	delete(urlParams, "flat_settings")
	return r
}
