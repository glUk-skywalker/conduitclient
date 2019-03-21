package responses

import (
	"github.com/gluk-skywalker/conduitclient/objects"
)

// ManiphestGetTaskTransactions is the response stricture for the reuqest `maniphest.gettasktransactions`
type ManiphestGetTaskTransactions map[string][]objects.Transaction
