package http

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func Get(baseURL string, param map[string]string, timeout time.Duration) ([]byte, error) {
	query := url.Values{}
	for k, v := range param {
		query.Add(k, v)
	}

	request, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		return nil, err
	}

	request.URL.RawQuery = query.Encode()

	response, err := (&http.Client{Timeout: timeout}).Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)
}

func PostForm(baseURL string, param map[string]string, timeout time.Duration) ([]byte, error) {
	formData := url.Values{}
	for k, v := range param {
		formData.Add(k, v)
	}

	response, err := (&http.Client{Timeout: timeout}).Post(baseURL, "application/x-www-form-urlencoded", bytes.NewBuffer([]byte(formData.Encode())))
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)
}

func PostJsonWithHeader(baseURL string, param map[string]string, header map[string]string, timeout time.Duration) ([]byte, error) {
	client := &http.Client{Timeout: timeout}

	bs, err := json.Marshal(&param)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", baseURL, bytes.NewBuffer(bs))
	if err != nil {
		return nil, err
	}

	for k, v := range header {
		request.Header.Add(k, v)
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)
}

func PostJson(baseURL string, param map[string]string, timeout time.Duration) ([]byte, error) {
	bs, err := json.Marshal(&param)
	if err != nil {
		return nil, err
	}

	response, err := (&http.Client{Timeout: timeout}).Post(baseURL, "application/json", bytes.NewBuffer(bs))
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)
}
