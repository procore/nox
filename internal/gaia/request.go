package gaia

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type request struct {
	client   *Client
	method   string
	endpoint string
	body     string
	headers  map[string]string
	params   map[string]string
}

// Get requests to elasticsearch cluster
func (r *request) get(endpoint string) string {
	r.method = http.MethodGet
	r.endpoint = endpoint
	return r.do()
}

// Put requestss to elasticsearch cluster
func (r *request) put(endpoint string, body string) string {
	r.method = http.MethodPut
	r.endpoint = endpoint
	r.body = body
	return r.do()
}

func (r *request) post(endpoint string, body string) string {
	r.method = http.MethodPost
	r.endpoint = endpoint
	r.body = body
	return r.do()
}

func (r *request) delete(endpoint string) string {
	r.method = http.MethodDelete
	r.endpoint = endpoint
	return r.do()
}

func (r *request) protocol() string {
	if r.client.Config.Net.TLS.Enable == true {
		return "https://"
	}
	return "http://"
}

func (r *request) authheader() (string, bool) {
	if r.client.Config.User.Name != "" && r.client.Config.User.Password != "" {
		return "Basic " + r.client.Config.User.Name + ":" + r.client.Config.User.Password, true
	}
	return "", false
}

func (r *request) do() string {

	c := &http.Client{}

	url := r.protocol() + r.client.Config.Net.Host + ":" + r.client.Config.Net.Port

	req, err := http.NewRequest(r.method, url+"/"+r.endpoint, strings.NewReader(r.body))
	if err != nil {
		log.Fatal(err)
	}

	if t, ok := r.authheader(); ok {
		r.headers["Authorization"] = t
	}

	if _, ok := r.headers["Content-Type"]; !ok {
		r.headers["Content-Type"] = "application/json; charset=utf-8"
	}

	if r.client.Config.Pretty {
		r.params["pretty"] = "true"
	}

	q := req.URL.Query()
	if r.client.Config.Pretty {
		r.params["pretty"] = "true"
	}
	for k, v := range r.params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	for k, v := range r.headers {
		req.Header.Set(k, v)
	}

	if r.client.Config.Debug {
		httpdump(req)
	}

	resp, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	d, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(d)
}
