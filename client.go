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

// UserWhoAmI performs the `service.whoami` request
func (c Client) UserWhoAmI() (responses.UserWhoAmI, error) {
	var userData responses.UserWhoAmI

	basicResp, err := c.request("user.whoami", parameters.UserWhoAmI{})
	if err != nil {
		return userData, errors.New("whoami request error: " + err.Error())
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

	basicResp, err := c.request("project.search", params)
	if err != nil {
		return projectData, errors.New("project.search request error: " + err.Error())
	}

	err = json.Unmarshal(basicResp, &projectData)
	if err != nil {
		return projectData, errors.New("response parsing error: " + err.Error())
	}

	return projectData, nil
}

func (c Client) ManiphestSearch(params parameters.ManiphestSearch) (responses.ManiphestSearch, error) {
	var tasksData responses.ManiphestSearch

	basicResp, err := c.request("maniphest.search", params)
	if err != nil {
		return tasksData, errors.New("maniphest.search request error: " + err.Error())
	}

	err = json.Unmarshal(basicResp, &tasksData)
	if err != nil {
		return tasksData, errors.New("response parsing error: " + err.Error())
	}

	return tasksData, nil
}

// ProjectColumnSearch performs the `project.column.search` request
func (c Client) ProjectColumnSearch(params parameters.ProjectColumnSearch) (responses.ProjectColumnSearch, error) {
	var columnsData responses.ProjectColumnSearch

	basicResp, err := c.request("project.column.search", params)
	if err != nil {
		return columnsData, errors.New("project.column.search request error: " + err.Error())
	}

	err = json.Unmarshal(basicResp, &columnsData)
	if err != nil {
		return columnsData, errors.New("response parsing error: " + err.Error())
	}

	return columnsData, nil
}

// ManiphestEdit performs the `maniphest.edit` request
func (c Client) ManiphestEdit(params parameters.ManiphestEdit) (responses.ManiphestEdit, error) {
	var editData responses.ManiphestEdit

	basicResp, err := c.request("maniphest.edit", params)
	if err != nil {
		return editData, errors.New("maniphest.edit request error: " + err.Error())
	}

	err = json.Unmarshal(basicResp, &editData)
	if err != nil {
		return editData, errors.New("response parsing error: " + err.Error())
	}

	return editData, nil
}

// Client is object for interaction with the Phabricator conduit API
type Client struct {
	url   string
	token string
}

func (c Client) generateURL(conduitMethod string, params url.Values) string {
	return c.url + "/api/" + conduitMethod + "?" + params.Encode()
}

func (c Client) request(path string, params interface{ ToConduitParams() url.Values }) (json.RawMessage, error) {
	urlParams := params.ToConduitParams()
	urlParams.Set("api.token", c.token)

	resp, err := http.Get(c.generateURL(path, urlParams))
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
