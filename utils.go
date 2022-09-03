package daraja

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type httpHeaders = map[string]string

func getRequest[model interface{}](url string, headers httpHeaders) (*model, error) {
	res := _makeRequest(url, http.MethodGet, headers, []byte{})
	if res != nil {
		defer res.Body.Close()
	}

	var response model
	err := json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to decode json response: %v", err)
	}

	return &response, nil
}

func postRequest[resModel interface{}](url string, body []byte, headers httpHeaders) (*resModel, error) {
	res := _makeRequest(url, http.MethodPost, headers, body)
	defer res.Body.Close()

	var response resModel
	err := json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to decode json response: %v", err)
	}

	return &response, nil
}

func _makeRequest(url string, method string, headers httpHeaders, body []byte) *http.Response {
	req, e := http.NewRequest(method, url, bytes.NewBuffer(body))
	if e != nil {
		panic("An internal error occurred when creating the request.")
	}

	req.Header.Set("Content-Type", "application/json")
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{Timeout: 60 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		panic("An internal error occurred when trying to send a request.")
	}

	return res
}
