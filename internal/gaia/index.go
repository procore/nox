package gaia

// IndexToggleOpen can toggle the
// open/close status of an index
// Args:
// - s string
// 	the state of the index (open, close)
// - i string
//  the name of the index to open/close
func (c *Client) IndexToggleOpen(s string, i string) string {
	r := c.newRequest()
	return r.post(i+"/_"+s, "")
}

// IndexDelete deletes an index
// Use with caution
// Args:
// - i string
//  the name of the index to delete
func (c *Client) IndexDelete(i string) string {
	r := c.newRequest()
	return r.delete(i)
}
