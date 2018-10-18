package responses

import "encoding/json"

// ConduitBasic is the base response stricture for any conduit request
type ConduitBasic struct {
	Result    json.RawMessage `json:"result"`
	ErrorCode string          `json:"error_code"`
	ErrorInfo string          `json:"error_info"`
}
