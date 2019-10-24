package parameters

import (
	"strconv"

	"github.com/gluk-skywalker/conduitclient/urlvalues"
)

// ProjectSearch is the structure for the params of `project.search` query
type ProjectSearch struct {
	QueryKey    []int
	Constraints struct {
		PHIDs []string
	}
}

// ToConduitParams turns the structure to urlvalues.URLValues
func (p ProjectSearch) ToConduitParams() urlvalues.URLValues {
	params := urlvalues.URLValues{}
	for i := 0; i < len(p.Constraints.PHIDs); i++ {
		params.Add("constraints[phids]["+strconv.Itoa(i)+"]", p.Constraints.PHIDs[i])
	}
	return params
}
