package parameters

import (
	"net/url"
)

// UserWhoAmI is the structure for the params of `project.search` query
type UserWhoAmI struct{}

// ToConduitParams turns the structure to url.Values
func (p UserWhoAmI) ToConduitParams() url.Values {
	return url.Values{}
}
