package model

type Milestone struct {
	Id          int    `json:"id"`
	Iid         int    `json:"iid"`
	ProjectId   int    `json:"project_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	State       string `json:"state"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DueDate     string `json:"due_date"`
}
