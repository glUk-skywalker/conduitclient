package responses

import "encoding/json"

// ConduitBasicReponse is the base response stricture for any conduit request
type ConduitBasicReponse struct {
	Result    json.RawMessage `json:"result"`
	ErrorCode string          `json:"error_code"`
	ErrorInfo string          `json:"error_info"`
}
