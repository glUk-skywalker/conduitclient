package serviceobjects

// PaginationCursor is the structure for data stored under `cursor` key in some responses
type PaginationCursor struct {
	Limit  int    `json:"limit"`
	After  int    `json:"after"`
	Before int    `json:"before"`
	Order  string `json:"order"`
}
