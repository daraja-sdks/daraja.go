package daraja

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type httpHeaders = map[string]string

func getRequest[model interface{}](app *mpesa, url string, headers httpHeaders) (*model, error) {
	if app.Environment == PRODUCTION {
		url = routes["production"] + url
	} else {
		url = routes["sandbox"] + url
	}

	req, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer([]byte{}))
	if err != nil {
		return nil, fmt.Errorf("failed to create a new request.%v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{Timeout: 60 * time.Second}
	res, err := client.Do(req)
	if res != nil {
		defer res.Body.Close()
	}
	if err != nil {
		return nil, fmt.Errorf("failed to send the request.%v", err)
	}

	fmt.Printf("response: %+v\n", res)
	var response model
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to decode json response: %v", err)
	}

	return &response, nil
}

func postRequest[resModel interface{}](app *mpesa, url string, body []byte, headers httpHeaders) (*resModel, error) {
	if app.Environment == PRODUCTION {
		url = routes["production"] + url
	} else {
		url = routes["sandbox"] + url
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create a new request.%v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{Timeout: 60 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send the request.%v", err)
	}

	defer res.Body.Close()
	var response resModel
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to decode json response: %v", err)
	}

	return &response, nil
}
