package parameters

import (
	"net/url"
	"strconv"

	"github.com/gluk-skywalker/conduitclient/objects"
)

// DifferentialRevisionEdit is the structure for the params of `differential.revision.edit` query
type DifferentialRevisionEdit struct {
	Transactions     []objects.URLParamAppendable
	ObjectIdentifier string
}

// ToConduitParams turns the structure to url.Values
func (p DifferentialRevisionEdit) ToConduitParams() url.Values {
	params := url.Values{}

	params.Set("objectIdentifier", p.ObjectIdentifier)

	for i, transaction := range p.Transactions {
		transactionPrefix := "transactions[" + strconv.Itoa(i) + "]"
		transaction.AppendTo(&params, transactionPrefix)
	}

	return params
}
