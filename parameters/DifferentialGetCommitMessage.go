package parameters

import "net/url"

// DifferentialGetCommitMessage is the structure for the params of `differential.getcommitmessage` query
type DifferentialGetCommitMessage struct {
	RevisionID string
}

// ToConduitParams turns the structure to url.Values
func (d DifferentialGetCommitMessage) ToConduitParams() url.Values {
	params := url.Values{}

	params.Set("revision_id", d.RevisionID)

	return params
}
