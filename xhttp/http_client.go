package xhttp

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

// PostRequest ...
func PostRequest(url string, param interface{}) (v []byte, err error) {
	v, err = PostRequestWithHeader(url, param, map[string]string{})
	return
}

// PostRequestWithHeader ...
func PostRequestWithHeader(url string, param interface{}, headers map[string]string) (v []byte, err error) {
	paramBytes, err := json.Marshal(param)
	if err != nil {
		return
	}
	client := &http.Client{}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(paramBytes))
	if err != nil {
		return
	}
	for field, value := range headers {
		request.Header.Set(field, value)
	}
	request.Header.Set("Content-type", "application/json")
	resp, err := client.Do(request)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	v, err = ioutil.ReadAll(resp.Body)
	return
}

// GetRequestAndParam Get http get method
func GetRequestAndParam(url string, params map[string]string, headers map[string]string) (*http.Response, error) {
	//new request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.New("new request is fail ")
	}
	//add params
	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}
	//if headers != nil {
	for key, val := range headers {
		req.Header.Add(key, val)
	}
	//http client
	client := &http.Client{Timeout: 30 * time.Second}
	return client.Do(req)
}
