package parameters

import "github.com/gluk-skywalker/conduitclient/urlvalues"

// UserWhoAmI is the structure for the params of `project.search` query
type UserWhoAmI struct{}

// ToConduitParams turns the structure to urlvalues.URLValues
func (p UserWhoAmI) ToConduitParams() urlvalues.URLValues {
	return urlvalues.URLValues{}
}
