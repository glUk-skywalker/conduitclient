package parameters

import (
	"net/url"
	"strconv"
)

// ManiphestSearch is the structure for the params of `maniphest.search` query
type ManiphestSearch struct {
	Constraints struct {
		IDs []int
	}
}

// ToConduitParams turns the structure to url.Values
func (p ManiphestSearch) ToConduitParams() url.Values {
	params := url.Values{}
	for i, v := range p.Constraints.IDs {
		params.Add("constraints[ids]["+strconv.Itoa(i)+"]", strconv.Itoa(v))
	}
	return params
}
