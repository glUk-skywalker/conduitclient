package ConduitClient

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"./responses"
)

// Client is object for interaction with the P_h_a_bricator conduit API
type Client struct {
	url   string
	token string
}

func (c Client) generateURL(conduitMethod string, params url.Values) string {
	return c.url + "/api/" + conduitMethod + "?" + params.Encode()
}

func (c Client) request(path string, params url.Values) (json.RawMessage, error) {
	params.Set("api.token", c.token)

	resp, err := http.Get(c.generateURL(path, params))
	if err != nil {
		return []byte{}, errors.New("request error: " + err.Error())
	}

	if resp.StatusCode != 200 {
		return []byte{}, errors.New("bad response code: " + strconv.Itoa(resp.StatusCode) + " - " + http.StatusText(resp.StatusCode))
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("reading response error: " + err.Error())
	}

	var conduitResp responses.ConduitBasicReponse
	err = json.Unmarshal(content, &conduitResp)
	if err != nil {
		return []byte{}, errors.New("response parsing error: " + err.Error())
	}

	if conduitResp.ErrorCode != "" {
		return []byte{}, errors.New("conduit error: " + conduitResp.ErrorInfo)
	}

	return conduitResp.Result, nil
}

// UserWhoAmI performs the `service.whoami` request
func (c Client) UserWhoAmI() (responses.UserWhoAmIResponse, error) {
	var userData responses.UserWhoAmIResponse

	basicResp, err := c.request("user.whoami", url.Values{})
	if err != nil {
		return userData, errors.New("whoami rquest error: " + err.Error())
	}

	err = json.Unmarshal(basicResp, &userData)
	if err != nil {
		return userData, errors.New("response parsing error: " + err.Error())
	}
	return userData, nil
}
