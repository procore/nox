package elastic

import (
	"log"
	"time"

	"github.com/Jeffail/gabs"
)

// package variables
var url string
var conf Config
var urlParams map[string]string
var headers map[string]string

// Config stores configuration for the elastic client
type Config struct {
	TLS         bool
	Username    string
	Password    string
	Debug       bool
	Host        string
	Port        string
	Pretty      bool
	clustername string
	timestamp   string
}

// InitConfig sets up config struct for client
// must be called before requests are made
func InitConfig(config *Config) {
	conf = *config
	urlParams = make(map[string]string)
	headers = make(map[string]string)

	if conf.Pretty {
		urlParams["pretty"] = "true"
	}

	url = protocol() + conf.Host + ":" + conf.Port

	r, err := gabs.ParseJSON([]byte(Get("")))
	if err != nil {
		log.Fatal(err)
	}
	if clusterName, ok := r.Path("cluster_name").Data().(string); ok {
		conf.clustername = clusterName
	}

	t := time.Now().UTC()
	conf.timestamp = t.Format("20060102")
}
