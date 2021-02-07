package api

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func buildURL(baseStr string, endpoint string) *url.URL {
	u, err := url.Parse(endpoint)
	if err != nil {
		log.Fatal(err)
	}

	base, err := url.Parse(baseStr)
	if err != nil {
		log.Fatal("Supplied URL is invalid.")
	}

	return base.ResolveReference(u)
}

// Request sends a HTTP request to the server.
func Request(base, endpoint, method, username, password string, hasBody bool, data interface{}) (interface{}, error) {
	client := &http.Client{}
	url := buildURL(base, endpoint)

	var reqBodyReader io.Reader
	reqBodyReader = nil

	if hasBody {
		reqBody, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}

		reqBodyReader = strings.NewReader(string(reqBody))
	}

	req, err := http.NewRequest(method, url.String(), reqBodyReader)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	req.SetBasicAuth(username, password)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Request failed with HTTP %s.", resp.Status)
	}

	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var obj interface{}

	err = json.Unmarshal(bodyText, &obj)
	if err != nil {
		log.Fatal(err)
	}

	return obj, nil
}

// Delete sends a HTTP DELETE request to the server.
func Delete(base, endpoint, username, password string) (interface{}, error) {
	return Request(base, endpoint, "DELETE", username, password, false, nil)
}

// Get sends a HTTP GET request to the server.
func Get(base, endpoint, username, password string) (interface{}, error) {
	return Request(base, endpoint, "GET", username, password, false, nil)
}

// Post sends a HTTP POST request to the server.
func Post(base, endpoint, username, password string, data interface{}) (interface{}, error) {
	return Request(base, endpoint, "POST", username, password, true, data)
}

// Put sends a HTTP PUT request to the server.
func Put(base, endpoint, username, password string, data interface{}) (interface{}, error) {
	return Request(base, endpoint, "PUT", username, password, true, data)
}
