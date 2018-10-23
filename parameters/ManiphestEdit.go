package parameters

import (
	"net/url"
	"strconv"
)

// ManiphestEdit is the structure for the params of `maniphest.edit` query
type ManiphestEdit struct {
	ObjectIdentifier string
	Transactions     []struct {
		Type    string
		Columns []string
	}
}

// ToConduitParams turns the structure to url.Values
func (p ManiphestEdit) ToConduitParams() url.Values {
	params := url.Values{}

	params.Set("objectIdentifier", p.ObjectIdentifier)

	for iTransaction, vTransaction := range p.Transactions {
		transactionPrefix := "transactions[" + strconv.Itoa(iTransaction) + "]"
		params.Add(transactionPrefix+"[type]", vTransaction.Type)
		for i, v := range vTransaction.Columns {
			params.Add(transactionPrefix+"[value]["+strconv.Itoa(i)+"]", v)
		}
	}

	return params
}
