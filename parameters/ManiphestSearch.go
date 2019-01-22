package parameters

import (
	"net/url"
	"strconv"
)

// ManiphestSearch is the structure for the params of `maniphest.search` query
type ManiphestSearch struct {
	Constraints struct {
		IDs      []int
		Statuses []string
		Projects []string
	}
	Order []string
	After string
}

// ToConduitParams turns the structure to url.Values
func (p ManiphestSearch) ToConduitParams() url.Values {
	params := url.Values{}
	for i, v := range p.Constraints.IDs {
		params.Add("constraints[ids]["+strconv.Itoa(i)+"]", strconv.Itoa(v))
	}
	for i, v := range p.Constraints.Statuses {
		params.Add("constraints[statuses]["+strconv.Itoa(i)+"]", v)
	}
	for i, v := range p.Constraints.Projects {
		params.Add("constraints[projects]["+strconv.Itoa(i)+"]", v)
	}
	for i, v := range p.Order {
		params.Add("order["+strconv.Itoa(i)+"]", v)
	}
	if p.After != "" {
		params.Add("after", p.After)
	}

	return params
}
