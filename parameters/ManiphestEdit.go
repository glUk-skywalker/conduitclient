package parameters

import (
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

	for _, vTransaction := range p.Transactions {
		vTransaction.AppendTo(&params, "transactions")
	}

	return params
}
