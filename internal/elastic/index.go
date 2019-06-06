package elastic

// IndexToggleOpen can toggle the
// open/close status of an index
func IndexToggleOpen(s string, i string) string {
	return Post(i+"/_"+s, "")
}

// IndexDelete deletes an index
// Use with caution
func IndexDelete(i string) string {
	return Delete(i)
}
