package elastic

// RemoteInfo returns info on remote clusters
func RemoteInfo() string {
	return Get("_remote/info")
}
