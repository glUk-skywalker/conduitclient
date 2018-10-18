package responses

import "github.com/gluk-skywalker/conduitclient/objects"

// ProjectSearch is the response stricture for the reuqest `project.search`
type ProjectSearch struct {
	Data []objects.Project `json:"data"`
}
