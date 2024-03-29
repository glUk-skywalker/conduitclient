package responses

import (
	"github.com/glUk-skywalker/conduitclient/objects"
	"github.com/glUk-skywalker/conduitclient/serviceobjects"
)

// ProjectColumnSearch is the response stricture for the reuqest `project.column.search`
type ProjectColumnSearch struct {
	Data  []objects.Column `json:"data"`
	Maps  struct{}         `json:"maps"`
	Query struct {
		QueryKey struct{} `json:"queryKey"`
	} `json:"query"`
	Cursor serviceobjects.PaginationCursor `json:"cursor"`
}
