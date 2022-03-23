package parameters

import (
	"net/url"
	"strconv"
)

// ManiphestSearch is the structure for the params of `maniphest.search` query
type ManiphestSearch struct {
	Constraints struct {
		IDs           []int
		Statuses      []string
		Projects      []string
		Priorities    []int
		ModifiedStart int64
		ModifiedEnd   int64
	}
	Order       []string
	After       string
	Attachments map[string]bool
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

	for i, v := range p.Constraints.Priorities {
		params.Add("constraints[priorities]["+strconv.Itoa(i)+"]", strconv.Itoa(v))
	}

	if p.Constraints.ModifiedStart != 0 {
		params.Add("constraints[modifiedStart]", strconv.FormatInt(p.Constraints.ModifiedStart, 10))
	}

	if p.Constraints.ModifiedEnd != 0 {
		params.Add("constraints[modifiedEnd]", strconv.FormatInt(p.Constraints.ModifiedStart, 10))
	}

	for i, v := range p.Order {
		params.Add("order["+strconv.Itoa(i)+"]", v)
	}

	if p.After != "" {
		params.Add("after", p.After)
	}

	for key, v := range p.Attachments {
		params.Add("attachments["+key+"]", btos(v))
	}

	return params
}

func btos(b bool) string {
	if b {
		return "1"
	}

	return "0"
}
