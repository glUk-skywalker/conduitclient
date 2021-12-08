package parameters

import (
	"net/url"
	"strconv"
)

// ProjectSearch is the structure for the params of `project.search` query
type ProjectSearch struct {
	QueryKey    []int
	Constraints struct {
		PHIDs   []string
		Members []string
	}
}

// ToConduitParams turns the structure to url.Values
func (p ProjectSearch) ToConduitParams() url.Values {
	params := url.Values{}
	for i, v := range p.Constraints.PHIDs {
		params.Add("constraints[phids]["+strconv.Itoa(i)+"]", v)
	}

	for i, v := range p.Constraints.Members {
		params.Add("constraints[members]["+strconv.Itoa(i)+"]", v)
	}

	return params
}
