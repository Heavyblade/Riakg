package riakapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var Host string = "localhost"
var Port string = "18098"

type BucketResponse struct {
	Bukckets []string `json:"buckets"`
}

func SetHost(host string) {
	Host = host
}

func SetPort(port string) {
	Port = port
}

func GetUrl() string {
	return fmt.Sprintf("http://%s:%s", Host, Port)
}

func Get(url string, params, headers map[string]string) (error, []byte) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error on creating request %v", err)
		return err, []byte{}
	}

	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("[HTTP ERROR]: %v", err)
		return err, []byte{}
	}

	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("[READ BODY]: %v", err)
		return err, []byte{}
	}

	return nil, respByte
}

func GetBuckets() BucketResponse {
	targetUrl := GetUrl() + "/buckets"
	params := map[string]string{"buckets": "true"}
	headers := map[string]string{"Accept": "application/json"}

	err, respByte := Get(targetUrl, params, headers)
	if err != nil {
		panic(err)
	}

	buckets := BucketResponse{}
	json.Unmarshal(respByte, &buckets)

	return buckets
}

func GetBucketKeys(bucketKey string) []string {
	escapedBucket := url.QueryEscape(bucketKey)
	targetUrl := GetUrl() + "/buckets/" + escapedBucket + "/keys"
	params := map[string]string{"keys": "true"}
	headers := map[string]string{"Accept": "application/json"}

	err, respByte := Get(targetUrl, params, headers)
	if err != nil {
		panic(err)
	}

	jsonStruct := struct {
		Keys []string `json:"keys"`
	}{}

	json.Unmarshal(respByte, &jsonStruct)

	return jsonStruct.Keys
}

func GetKeyValue(bucket, key string) string {
	escapedBucket := url.QueryEscape(bucket)
	escapedKey := url.QueryEscape(key)
	targetUrl := GetUrl() + "/buckets/" + escapedBucket + "/keys/" + escapedKey
	headers := map[string]string{"Accept": "application/json"}

	err, respByte := Get(targetUrl, map[string]string{}, headers)
	if err != nil {
		panic(err)
	}

	return string(respByte)
}
