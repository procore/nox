package elastic

import (
	"log"
	"net/http"
	"strings"
)

// KibanaQuery allows for running kibana-like commands
// it takes in a command string [METHOD] [ENDPOINT]
// and a json body and returns the results from ES
func KibanaQuery(command string, body string) string {
	cmdArray := strings.Split(command, " ")
	if len(cmdArray) != 2 {
		log.Fatal("Malformed query")
	}
	method := strings.ToUpper(cmdArray[0])
	u := cmdArray[1]

	switch method {
	case "GET":
		return request(http.MethodGet, u, body)
	case "PUT":
		return request(http.MethodPut, u, body)
	case "POST":
		return request(http.MethodPost, u, body)
	case "DELETE":
		return request(http.MethodDelete, u, body)
	default:
		return "Unknown method keyword " + method
	}
}
