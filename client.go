package conduitclient

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gluk-skywalker/conduitclient/parameters"
	"github.com/gluk-skywalker/conduitclient/responses"
)

// New creates and instance of Client
func New(path string, token string) Client {
	return Client{url: path, token: token}
}

// Client is object for interaction with the Phabricator conduit API
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

	var conduitResp responses.ConduitBasic
	err = json.Unmarshal(content, &conduitResp)
	if err != nil {
		return []byte{}, errors.New("response parsing error: " + err.Error())
	}

	if conduitResp.ErrorCode != "" {
		return []byte{}, errors.New("conduit error: [" + conduitResp.ErrorCode + "] " + conduitResp.ErrorInfo)
	}

	return conduitResp.Result, nil
}

// UserWhoAmI performs the `service.whoami` request
func (c Client) UserWhoAmI() (responses.UserWhoAmI, error) {
	var userData responses.UserWhoAmI

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

// ProjectSearch performs the `project.search` request
func (c Client) ProjectSearch(params parameters.ProjectSearch) (responses.ProjectSearch, error) {
	var projectData responses.ProjectSearch

	urlParams := url.Values{}
	// TODO: SET THE PARAMETERS
	// if len(params.QueryKey) > 0
	// end

	basicResp, err := c.request("project.search", urlParams)
	if err != nil {
		return projectData, errors.New("project.search rquest error: " + err.Error())
	}

	err = json.Unmarshal(basicResp, &projectData)
	if err != nil {
		return projectData, errors.New("response parsing error: " + err.Error())
	}

	return projectData, nil
}
