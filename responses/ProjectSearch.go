package responses

import (
	"github.com/gluk-skywalker/conduitclient/objects"
	"github.com/gluk-skywalker/conduitclient/serviceobjects"
)

// ProjectSearch is the response stricture for the reuqest `project.search`
type ProjectSearch struct {
	Data []objects.Project `json:"data"`
	Maps struct {
		SlugMap struct{} `json:"slugMap"`
	} `json:"maps"`
	Query struct {
		QueryKey struct{} `json:"queryKey"`
	} `json:"query"`
	Cursor serviceobjects.PaginationCursor `json:"cursor"`
}
