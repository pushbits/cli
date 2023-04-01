// Package api provides low-level functionality to interact with the PushBits API.
package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/pushbits/cli/internal/handling"
	log "github.com/sirupsen/logrus"
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
func Request(base, endpoint, method, proxy, username, password string, hasBody bool, data interface{}) (interface{}, error) {
	transport := http.DefaultTransport

	if len(proxy) > 0 {
		log.Printf("Using proxy '%s'", proxy)

		proxyURL, err := url.Parse(proxy)
		if err != nil {
			return nil, err
		}

		transport = &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
	}

	client := &http.Client{
		Transport: transport,
	}

	url := buildURL(base, endpoint)

	var reqBodyReader io.Reader
	reqBodyReader = nil

	if hasBody {
		reqBody, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}

		reqBodyReader = strings.NewReader(string(reqBody))
	}

	req, err := http.NewRequest(method, url.String(), reqBodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	req.SetBasicAuth(username, password)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer handling.Close(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with HTTP %s", resp.Status)
	}

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var obj interface{}

	err = json.Unmarshal(bodyText, &obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

// Delete sends a HTTP DELETE request to the server.
func Delete(base, endpoint, proxy, username, password string) (interface{}, error) {
	return Request(base, endpoint, "DELETE", proxy, username, password, false, nil)
}

// Get sends a HTTP GET request to the server.
func Get(base, endpoint, proxy, username, password string) (interface{}, error) {
	return Request(base, endpoint, "GET", proxy, username, password, false, nil)
}

// Post sends a HTTP POST request to the server.
func Post(base, endpoint, proxy, username, password string, data interface{}) (interface{}, error) {
	return Request(base, endpoint, "POST", proxy, username, password, true, data)
}

// Put sends a HTTP PUT request to the server.
func Put(base, endpoint, proxy, username, password string, data interface{}) (interface{}, error) {
	return Request(base, endpoint, "PUT", proxy, username, password, true, data)
}
