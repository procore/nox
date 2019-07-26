package gaia

// RemoteInfo returns info on remote clusters
func (c *Client) RemoteInfo() string {
	r := c.newRequest()
	return r.get("_remote/info")
}
