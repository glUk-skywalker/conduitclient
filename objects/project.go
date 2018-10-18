package objects

// Project is the struct for the project data
type Project struct {
	ID     int    `json:"id"`
	Type   string `json:"type"`
	PHID   string `json:"phid"`
	Fields struct {
		Name      string `json:"name"`
		Slug      string `json:"slug"`
		Milestone string `json:"milestone"`
		Depth     int64  `json:"depth"`
		Parent    string `json:"parent"`
		Icon      struct {
			Key  string `json:"key"`
			Name string `json:"name"`
			Icon string `json:"icon"`
		} `json:"icon"`
		Color struct {
			Key  string `json:"color"`
			Name string `json:"name"`
		} `json:"color"`
		SpacePHID    string `json:"spacePHID"`
		DateCreated  int64  `json:"dateCreated"`
		DateModified int64  `json:"dateModified"`
		Policy       struct {
			View string `json:"view"`
			Edit string `json:"edit"`
			Join string `json:"join"`
		} `json:"policy"`
		Description string `json:"description"`
	} `json:"fields"`
	Attachments Attachment `json:"attachments"`
}
