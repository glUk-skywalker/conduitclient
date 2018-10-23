package parameters

import (
	"net/url"
	"strconv"
)

// ProjectColumnSearch is the structure for the params of `project.column.search` query
type ProjectColumnSearch struct {
	Constraints struct {
		Projects []string
	}
}

// ToConduitParams turns the structure to url.Values
func (p ProjectColumnSearch) ToConduitParams() url.Values {
	params := url.Values{}
	for i, v := range p.Constraints.Projects {
		params.Add("constraints[projects]["+strconv.Itoa(i)+"]", v)
	}
	return params
}
