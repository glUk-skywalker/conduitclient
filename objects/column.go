package objects

// Column is the struct for the clumn data
type Column struct {
	ID     int    `json:"id"`
	Type   string `json:"type"`
	PHID   string `json:"phid"`
	Fields struct {
		Name string `json:"name"`
	} `json:"fields"`
}
