package model

type Commit struct {
	Id          string `json:"id"`
	ShortId     string `json:"short_id"`
	Title       string `json:"title"`
	AuthorName  string `json:"author_name"`
	AuthorEmail string `json:"author_email"`
	CreatedAt   string `json:"created_at"`
	Message     string `json:"message"`
}
