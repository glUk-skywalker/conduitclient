package responses

import (
	"github.com/glUk-skywalker/conduitclient/objects"
	"github.com/glUk-skywalker/conduitclient/serviceobjects"
)

// ManiphestSearch is the response stricture for the reuqest `maniphest.search`
type ManiphestSearch struct {
	Data  []objects.Task `json:"data"`
	Maps  struct{}       `json:"maps"`
	Query struct {
		QueryKey struct{} `json:"queryKey"`
	} `json:"query"`
	Cursor serviceobjects.PaginationCursor `json:"cursor"`
}
