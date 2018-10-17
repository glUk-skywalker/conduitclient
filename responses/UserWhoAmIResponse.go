package responses

// UserWhoAmIResponse is the response stricture for the reuqest `user.whoami`
type UserWhoAmIResponse struct {
	Phid         string   `json:"phid"`
	UserName     string   `json:"userName"`
	RealName     string   `json:"realName"`
	Image        string   `json:"image"`
	URI          string   `json:"uri"`
	Roles        []string `json:"roles"`
	PrimaryEmail string   `json:"primaryEmail"`
}
