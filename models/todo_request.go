package models

type CreateTodoRequest struct {
	Title       *string `json:"title" binding:"required,min=1,max=100"`
	Description *string `json:"description"`
}

type UpdateTodoRequest struct {
	Title       *string `json:"title" binding:"omitempty,min=1,max=100"`
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
}
