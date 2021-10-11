package riakapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var Host string = "localhost"
var Port string = "18098"

func SetHost(host string) {
	Host = host
}

func SetPort(port string) {
	Port = port
}

type BucketResponse struct {
	Bukckets []string `json:"buckets"`
}

func GetBuckets() BucketResponse {
	client := &http.Client{}
	url := fmt.Sprintf("http://%s:%s/buckets", Host, Port)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error on creating request %v", err)
	}

	q := req.URL.Query()
	q.Add("buckets", "true")
	req.URL.RawQuery = q.Encode()

	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("[HTTP ERROR]: %v", err)

	}

	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("[READ BODY]: %v", err)
	}

	buckets := BucketResponse{}
	json.Unmarshal(respByte, &buckets)

	return buckets
}

func GetUrl() string {
	return fmt.Sprintf("http://%s:%s", Host, Port)
}

func GetBucketKeys(bucketKey string) []string {
	client := &http.Client{}

	escapedBucket := url.QueryEscape(bucketKey)
	targetUrl := GetUrl() + "/buckets/" + escapedBucket + "/keys?keys=true"

	log.Println(targetUrl)

	req, err := http.NewRequest("GET", targetUrl, nil)
	if err != nil {
		fmt.Printf("Error on creating request %v", err)
	}

	//q := req.URL.Query()
	//q.Add("keys", "true")
	//req.URL.RawQuery = q.Encode()

	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("[HTTP ERROR]: %v", err)
	}

	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("[READ BODY]: %v", err)
	}

	jsonStruct := struct {
		Keys []string `json:"keys"`
	}{}

	json.Unmarshal(respByte, &jsonStruct)

	return jsonStruct.Keys
}