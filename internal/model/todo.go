package model

type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Desc string `json:"desc"`
	Priority *int `json:"priority,omitempty"`
	DueDate string `json:"due_date"`
	CompletedAt string `json:"completed_at,omitempty"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
