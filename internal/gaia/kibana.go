package gaia

import (
	"log"
	"net/http"
	"strings"
)

// KibanaQuery allows for running kibana-like commands
// it takes in a command string [METHOD] [ENDPOINT]
// and a json body and returns the results from ES
func (c *Client) KibanaQuery(command string, body string) string {
	r := c.newRequest()
	cmdArray := strings.Split(command, " ")
	if len(cmdArray) != 2 {
		log.Fatal("Malformed query")
	}
	method := strings.ToUpper(cmdArray[0])
	u := cmdArray[1]

	switch method {
	case "GET":
		r.method = http.MethodGet
	case "PUT":
		r.method = http.MethodPut
	case "POST":
		r.method = http.MethodPost
	case "DELETE":
		r.method = http.MethodDelete
	default:
		return "Unknown method keyword " + method
	}
	r.endpoint = u
	r.body = body
	return r.do()
}
