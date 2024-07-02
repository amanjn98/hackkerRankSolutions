package Trial

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func hitAPIWithBody(url string, method string, body interface{}) ([]byte, error) {
	// Create a new HTTP client
	client := &http.Client{}

	// Marshal the request body (if applicable)
	var jsonData []byte
	if body != nil {
		jsonData, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	// Create a new HTTP request
	req, err := http.NewRequest(method, url, bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}

	// Set content type header if request body exists
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	// Do the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
