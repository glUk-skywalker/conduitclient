package parameters

import "github.com/gluk-skywalker/conduitclient/urlvalues"

// DifferentialGetCommitMessage is the structure for the params of `differential.getcommitmessage` query
type DifferentialGetCommitMessage struct {
	RevisionID string
}

// ToConduitParams turns the structure to urlvalues.URLValues
func (d DifferentialGetCommitMessage) ToConduitParams() urlvalues.URLValues {
	params := urlvalues.URLValues{}

	params.Set("revision_id", d.RevisionID)

	return params
}
