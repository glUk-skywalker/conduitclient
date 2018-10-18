package responses

type project struct {
	ID int `json:"id"`
}

// ProjectSearch is the response stricture for the reuqest `project.search`
type ProjectSearch struct {
	Data []project `json:"data"`
}
