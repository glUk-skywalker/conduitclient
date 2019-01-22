package serviceobjects

// PaginationCursor is the structure for data stored under `cursor` key in some responses
type PaginationCursor struct {
	Limit  int      `json:"limit"`
	After  string   `json:"after"`
	Before string   `json:"before"`
	Order  []string `json:"order"`
}
