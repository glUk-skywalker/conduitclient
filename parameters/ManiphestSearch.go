package parameters

import (
	"strconv"

	"github.com/gluk-skywalker/conduitclient/urlvalues"
)

// ManiphestSearch is the structure for the params of `maniphest.search` query
type ManiphestSearch struct {
	Constraints struct {
		IDs      []int
		Statuses []string
		Projects []string
	}
	Order       []string
	After       string
	Attachments map[string]bool
}

// ToConduitParams turns the structure to urlvalues.URLValues
func (p ManiphestSearch) ToConduitParams() urlvalues.URLValues {
	params := urlvalues.URLValues{}
	for _, v := range p.Constraints.IDs {
		params.Add("constraints[ids]", strconv.Itoa(v))
	}
	for _, v := range p.Constraints.Statuses {
		params.Add("constraints[statuses]", v)
	}
	for _, v := range p.Constraints.Projects {
		params.Add("constraints[projects]", v)
	}
	for _, v := range p.Order {
		params.Add("order", v)
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
