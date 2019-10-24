package parameters

import (
	"strconv"

	"github.com/gluk-skywalker/conduitclient/urlvalues"
)

// ProjectColumnSearch is the structure for the params of `project.column.search` query
type ProjectColumnSearch struct {
	Constraints struct {
		Projects []string
	}
}

// ToConduitParams turns the structure to urlvalues.URLValues
func (p ProjectColumnSearch) ToConduitParams() urlvalues.URLValues {
	params := urlvalues.URLValues{}
	for i, v := range p.Constraints.Projects {
		params.Add("constraints[projects]["+strconv.Itoa(i)+"]", v)
	}
	return params
}
