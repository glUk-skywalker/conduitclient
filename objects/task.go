package objects

// Task is the struct for the task data
type Task struct {
	ID     int    `json:"id"`
	Type   string `json:"type"`
	PHID   string `json:"phid"`
	Fields struct {
		Name        string `json:"name"`
		DateClosed  int64  `json:"dateClosed"`
		DateCreated int64  `json:"dateCreated"`
	} `json:"fields"`
}
