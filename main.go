package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/namsral/flag"
)

var (
	vmalertHost   = flag.String("host", "localhost", "Host where VMAlert responds")
	vmalertPort   = flag.Int("port", 8880, "VMAlert port")
	vmalertAction = flag.String("action", "groups", "VMAlert action to take {groups|alerts|metrics|reload}")
)

func init() {
	flag.Parse()
}

func getJsonData(apiBase string, apiEndpoint string) []byte {
	response, err := http.Get(apiBase + apiEndpoint)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	return data
}

func main() {
	host := *vmalertHost
	action := *vmalertAction
	vmalertBase := "http://" + host + ":" + strconv.Itoa(*vmalertPort)

	switch takeAction := action; takeAction {
	case "groups":
		endpoint := "/api/v1/groups"
		fmt.Println(string(getJsonData(vmalertBase, endpoint)))
	case "alerts":
		endpoint := "/api/v1/alerts"
		fmt.Println(string(getJsonData(vmalertBase, endpoint)))
	case "metrics":
		endpoint := "/metrics"
		fmt.Println(string(getJsonData(vmalertBase, endpoint)))
	case "reload":
		endpoint := "/-/reload"
		fmt.Println(string(getJsonData(vmalertBase, endpoint)))
	}
}
