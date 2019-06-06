package elastic

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// Get requests to elasticsearch cluster
func Get(endpoint string) string {
	return request(http.MethodGet, endpoint, "")
}

// Put requestss to elasticsearch cluster
func Put(endpoint string, body string) string {
	return request(http.MethodPut, endpoint, body)

}

// Post requests
func Post(endpoint string, body string) string {
	return request(http.MethodPost, endpoint, body)
}

// Delete requests
func Delete(endpoint string) string {
	return request(http.MethodDelete, endpoint, "")
}

func request(method, endpoint string, body string) string {
	// var mux sync.Mutex
	// mux.Lock()
	client := &http.Client{}

	if url == "" {
		log.Fatal("you need to call InitConfig() before making requests")
	}

	req, err := http.NewRequest(method, url+"/"+endpoint, strings.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}

	if t, ok := authheader(); ok {
		headers["Authorization"] = t
	}

	if _, ok := headers["Content-Type"]; !ok {
		headers["Content-Type"] = "application/json; charset=utf-8"
	}

	q := req.URL.Query()
	for k, v := range urlParams {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	if conf.Debug {
		httpdump(req)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	d, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// mux.Unlock()
	return string(d)
}
