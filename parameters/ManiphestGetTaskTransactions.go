package parameters

import (
	"net/url"
	"strconv"
)

// ManiphestGetTaskTransactions is the structure for the params of `maniphest.gettasktransactions` query
type ManiphestGetTaskTransactions []int

// ToConduitParams turns the structure to url.Values
func (p ManiphestGetTaskTransactions) ToConduitParams() url.Values {
	params := url.Values{}
	for i, v := range p {
		params.Add("ids["+strconv.Itoa(i)+"]", strconv.Itoa(v))
	}
	return params
}
