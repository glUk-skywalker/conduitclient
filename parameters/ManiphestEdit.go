package parameters

import (
	"net/url"
	"strconv"

	"github.com/gluk-skywalker/conduitclient/objects"
)

// ManiphestEdit is the structure for the params of `maniphest.edit` query
type ManiphestEdit struct {
	ObjectIdentifier string
	Transactions     []objects.URLParamAppendable
}

// ToConduitParams turns the structure to url.Values
func (p ManiphestEdit) ToConduitParams() url.Values {
	params := url.Values{}

	params.Set("objectIdentifier", p.ObjectIdentifier)

	for iTransaction, vTransaction := range p.Transactions {
		transactionPrefix := "transactions[" + strconv.Itoa(iTransaction) + "]"
		vTransaction.AppendTo(&params, transactionPrefix)
	}

	return params
}
