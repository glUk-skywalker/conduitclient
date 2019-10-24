package parameters

import (
	"strconv"

	"github.com/gluk-skywalker/conduitclient/objects"
	"github.com/gluk-skywalker/conduitclient/urlvalues"
)

// ManiphestEdit is the structure for the params of `maniphest.edit` query
type ManiphestEdit struct {
	ObjectIdentifier string
	Transactions     []objects.URLParamAppendable
}

// ToConduitParams turns the structure to urlvalues.URLValues
func (p ManiphestEdit) ToConduitParams() urlvalues.URLValues {
	params := urlvalues.URLValues{}

	params.Set("objectIdentifier", p.ObjectIdentifier)

	for iTransaction, vTransaction := range p.Transactions {
		transactionPrefix := "transactions[" + strconv.Itoa(iTransaction) + "]"
		vTransaction.AppendTo(&params, transactionPrefix)
	}

	return params
}
