package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// Client is object for interaction with the Phabricator conduit API
type Client struct {
	url   string
	token string
}

type response struct {
	Result    json.RawMessage `json:"result"`
	ErrorCode string          `json:"error_code"`
	ErrorInfo string          `json:"error_info"`
}

func (c Client) request(path string, params url.Values) (json.RawMessage, error) {
	params.Set("api.token", c.token)
	fullURL := c.url + "/api" + path + "?" + params.Encode()
	var (
		err  error
		resp *http.Response
	)

	resp, err = http.Get(fullURL)
	if err != nil {
		return []byte{}, errors.New("request error: " + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return []byte{}, errors.New("bad response code: " + strconv.Itoa(resp.StatusCode) + " - " + http.StatusText(resp.StatusCode))
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("reading response error: " + err.Error())
	}

	var phabResp response
	err = json.Unmarshal(content, &phabResp)
	if err != nil {
		return []byte{}, errors.New("response parsing error: " + err.Error())
	}

	if phabResp.ErrorCode != "" {
		return []byte{}, errors.New("phabricator error: " + phabResp.ErrorInfo)
	}

	return phabResp.Result, nil
}
