package model

type Author struct {
	Name      string `json:"name"`
	Username  string `json:"username"`
	Id        int    `json:"id"`
	State     string `json:"state"`
	AvatarUrl string `json:"avatar_url"`
	WebUrl    string `json:"web_url"`
}
