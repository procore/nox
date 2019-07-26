package gaia

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"regexp"
)

func grep(text string, match string) []string {
	re := regexp.MustCompile(`.*` + match + `.*`)
	return re.FindAllString(text, -1)
}

func httpdump(req *http.Request) {
	d, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", d)
}
