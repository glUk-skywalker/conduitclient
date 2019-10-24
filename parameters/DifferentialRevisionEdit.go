package parameters

import (
	"strconv"

	"github.com/gluk-skywalker/conduitclient/objects"
	"github.com/gluk-skywalker/conduitclient/urlvalues"
)

// DifferentialRevisionEdit is the structure for the params of `differential.revision.edit` query
type DifferentialRevisionEdit struct {
	Transactions     []objects.URLParamAppendable
	ObjectIdentifier string
}

// ToConduitParams turns the structure to urlvalues.URLValues
func (p DifferentialRevisionEdit) ToConduitParams() urlvalues.URLValues {
	params := urlvalues.URLValues{}

	params.Set("objectIdentifier", p.ObjectIdentifier)

	for i, transaction := range p.Transactions {
		transactionPrefix := "transactions[" + strconv.Itoa(i) + "]"
		transaction.AppendTo(&params, transactionPrefix)
	}

	return params
}
