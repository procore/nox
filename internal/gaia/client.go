package gaia

// Client wraps config for making requests to the Elasticsearch cluster
type Client struct {
	Config *Config
}

// NewClient returns a client that is ready to make requests to the
// elasticsearch cluster
func NewClient(config *Config) *Client {
	return &Client{Config: config}
}

func (c *Client) newRequest() *request {
	r := &request{client: c}
	r.params = make(map[string]string)
	r.headers = make(map[string]string)
	return r
}
