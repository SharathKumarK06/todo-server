package repository

import (
	"github.com/SharathKumarK06/todo-server/models"
)

type TodoRepository interface {
	Create(todo models.Todo) (models.Todo, error)
	List() ([]models.Todo, error)
	GetByID(id int) (models.Todo, error)
	Update(id int, todo models.Todo) (models.Todo, error)
	Delete(id int) error
}
