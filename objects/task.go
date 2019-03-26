package objects

// Task is the struct for the task data
type Task struct {
	ID     int    `json:"id"`
	Type   string `json:"type"`
	PHID   string `json:"phid"`
	Fields struct {
		Name         string `json:"name"`
		DateClosed   int64  `json:"dateClosed"`
		DateCreated  int64  `json:"dateCreated"`
		DateModified int    `json:"dateModified"`
		Description  struct {
			Raw string `json:"raw"`
		} `json:"description"`
		AuthorPHID string `json:"authorPHID"`
		OwnerPHID  string `json:"ownerPHID"`
		СloserPHID string `json:"closerPHID"`
		Status     struct {
			Value string `json:"value"`
			Name  string `json:"name"`
		} `json:"status"`
		Priority struct {
			Value int    `json:"value"`
			Name  string `json:"name"`
			Color string `json:"color"`
		} `json:"priority"`
	} `json:"fields"`
	Attachments struct {
		Projects struct {
			ProjectPHIDs []string `json:"projectPHIDs"`
		} `json:"projects"`
		Columns struct {
			Boards map[string]struct {
				Columns []struct {
					ID   int    `json:"id"`
					PHID string `json:"phid"`
					Name string `json:"name"`
				}
			} `json:"boards"`
		} `json:"columns"`
	} `json:"attachments"`
}
