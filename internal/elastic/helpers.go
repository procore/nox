package elastic

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

func protocol() string {
	if conf.TLS == true {
		return "https://"
	}
	return "http://"
}

func authheader() (string, bool) {
	if conf.Username != "" && conf.Password != "" {
		return "Basic " + conf.Username + ":" + conf.Password, true
	}
	return "", false
}

func httpdump(req *http.Request) {
	d, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", d)
}

func snapshotdefaultargs(n string, r string, f string) (string, string, string) {
	if r == "" {
		r = conf.clustername
	}

	if f != "" {
		r = r + "_" + f
	}

	if n == "" {
		n = conf.clustername + "_" + conf.timestamp
	}

	return n, r, f

}
