package responses

import (
	"github.com/gluk-skywalker/conduitclient/objects"
	"github.com/gluk-skywalker/conduitclient/serviceobjects"
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
