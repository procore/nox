package gaia

import (
	"fmt"
	"strconv"
)

// ClusterInfo returns general data about your cluster
func (c *Client) ClusterInfo() string {
	r := c.newRequest()
	return r.get("")
}

// ClusterHealth returns the general health of the cluster
func (c *Client) ClusterHealth() string {
	r := c.newRequest()
	return r.get("_cluster/health")
}

// ClusterState returns the general health of the cluster
func (c *Client) ClusterState() string {
	r := c.newRequest()
	return r.get("_cluster/state")
}

// ClusterStats Get some sweet sweet stats about your cluster
func (c *Client) ClusterStats() string {
	r := c.newRequest()
	r.params["human"] = "true"
	return r.get("_cluster/stats")
}

// ClusterPendingTasks returns list of cluster level
// changes that have not been executedut
func (c *Client) ClusterPendingTasks() string {
	r := c.newRequest()
	return r.get("_cluster/pending_tasks")
}

// ClusterReroute explicitly executes a
// cluster reroute allocation command
func (c *Client) ClusterReroute(body string) string {
	r := c.newRequest()
	return r.post("_cluster/reroute", body)
}

// ClusterSettings returns current cluster level settings
func (c *Client) ClusterSettings(flatSettings bool) string {
	r := c.newRequest()
	r.params["flat_settings"] = strconv.FormatBool(flatSettings)
	return r.get("_cluster/settings")
}

// ClusterUpdateSettings allows updating cluster wide settings
func (c *Client) ClusterUpdateSettings(body string, flatSettings bool) string {
	r := c.newRequest()
	r.params["flat_settings"] = strconv.FormatBool(flatSettings)
	return r.put("_cluster/settings", body)
}

// ToggleRouting turns off dynaamic allocation
// Parameters:
//	- setting string
//		string setting for cluster allocation
func (c *Client) ToggleRouting(setting string, flatSettings bool) string {
	r := c.newRequest()
	r.params["flat_settings"] = strconv.FormatBool(flatSettings)
	t := fmt.Sprintf("{\"transient\": {\"cluster.routing.allocation.enable\": \"%s\"},"+
		"\"persistent\": {\"cluster.routing.allocation.enable\": \"%s\"}}", setting, setting)
	return r.put("_cluster/settings", t)
}
