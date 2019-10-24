package parameters

import (
	"strconv"

	"github.com/gluk-skywalker/conduitclient/urlvalues"
)

// ManiphestGetTaskTransactions is the structure for the params of `maniphest.gettasktransactions` query
type ManiphestGetTaskTransactions []int

// ToConduitParams turns the structure to urlvalues.URLValues
func (p ManiphestGetTaskTransactions) ToConduitParams() urlvalues.URLValues {
	params := urlvalues.URLValues{}
	for _, v := range p {
		params.Add("ids", strconv.Itoa(v))
	}
	return params
}
