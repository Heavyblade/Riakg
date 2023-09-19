package riakapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"

	"github.com/tidwall/pretty"
)

var Host string = "localhost"
var Port string = "18098"
var Username string
var Password string

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

func Request(verb, url string, params, headers map[string]string, body io.Reader) (error, []byte) {
	client := &http.Client{}

	req, err := http.NewRequest(verb, url, body)
	if err != nil {
		fmt.Printf("Error on creating request %v", err)
		return err, []byte{}
	}
	if Username != "" && Password != "" {
		req.SetBasicAuth(Username, Password)
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

func Get(url string, params, headers map[string]string) (error, []byte) {
	return Request("GET", url, params, headers, nil)
}

func Delete(url string, params, headers map[string]string) (error, []byte) {
	return Request("DELETE", url, params, headers, nil)
}

func Put(url string, params, headers map[string]string, body string) (error, []byte) {
	return Request("PUT", url, params, headers, bytes.NewBuffer([]byte(body)))
}

func Post(url string, params, headers map[string]string, body string) (error, []byte) {
	return Request("Post", url, params, headers, bytes.NewBuffer([]byte(body)))
}

func GetBuckets() BucketResponse {
	targetUrl := GetUrl() + "/buckets"
	params := map[string]string{"buckets": "true"}
	headers := map[string]string{"Accept": "application/json"}

	err, respByte := Get(targetUrl, params, headers)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
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
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	jsonStruct := struct {
		Keys []string `json:"keys"`
	}{}

	json.Unmarshal(respByte, &jsonStruct)
	sort.Strings(jsonStruct.Keys)

	return jsonStruct.Keys
}

func GetKeyValue(bucket, key string) string {
	escapedBucket := url.QueryEscape(bucket)
	escapedKey := url.QueryEscape(key)
	targetUrl := GetUrl() + "/buckets/" + escapedBucket + "/keys/" + escapedKey
	headers := map[string]string{"Accept": "application/json"}

	err, respByte := Get(targetUrl, map[string]string{}, headers)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	prettified := pretty.Pretty(respByte)
	highlighted := pretty.Color([]byte(prettified), nil)
	return string(highlighted)
}

func DeleteKey(bucket, key string) bool {
	escapedBucket := url.QueryEscape(bucket)
	escapedKey := url.QueryEscape(key)
	targetUrl := GetUrl() + "/buckets/" + escapedBucket + "/keys/" + escapedKey
	headers := map[string]string{"Accept": "application/json"}

	err, _ := Delete(targetUrl, map[string]string{}, headers)
	return err == nil
}

func UpdateKeyValue(bucket, key, value string) bool {
	escapedBucket := url.QueryEscape(bucket)
	escapedKey := url.QueryEscape(key)

	targetUrl := GetUrl() + "/buckets/" + escapedBucket + "/keys/" + escapedKey
	headers := map[string]string{"Accept": "application/json", "Content-Type": "application/json"}

	err, body := Put(targetUrl, map[string]string{}, headers, value)
	log.Print(string(body))
	return err == nil
}
