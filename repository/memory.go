package repository

import (
	"fmt"
	"slices"
	"time"

	"github.com/SharathKumarK06/todo-server/models"
)

type InMemoryRepo struct {
	todos  []models.Todo
	nextID int
}

func NewInMemoryRepo() *InMemoryRepo {
	return &InMemoryRepo{
		todos:  []models.Todo{},
		nextID: 0,
	}
}

func (r *InMemoryRepo) Create(todo models.Todo) (models.Todo, error) {
	if todo.Title == "" {
		return models.Todo{}, fmt.Errorf("Title is empty!")
	}
	if len(todo.Title) < 1 || len(todo.Title) > 100 {
		return models.Todo{}, fmt.Errorf("Title should have minimum 1 and maximum 100 characters!")
	}

	now := time.Now()

	r.nextID++
	todo.ID = r.nextID
	todo.CreatedAt = now
	todo.UpdatedAt = now
	r.todos = append(r.todos, todo)

	return todo, nil
}

func (r *InMemoryRepo) List() ([]models.Todo, error) {
	if r.todos == nil {
		return []models.Todo{}, fmt.Errorf("No repository exist!")
	}
	return r.todos, nil
}

func (r *InMemoryRepo) GetByID(id int) (models.Todo, error) {
	index, err := r.findIndexByID(id)
	if err != nil {
		return models.Todo{}, err
	}
	if r.todos[index].ID == id {
		return r.todos[index], nil
	}

	return models.Todo{}, fmt.Errorf("Couldn't find todo with ID %d!", id)
}

func (r *InMemoryRepo) Update(id int, todo models.Todo) (models.Todo, error) {
	oldTodo, err := r.GetByID(id)
	updated := false
	if err != nil {
		return models.Todo{}, err
	}
	index, err := r.findIndexByID(id)
	if err != nil {
		return models.Todo{}, err
	}

	if todo.Title == "" {
		return models.Todo{}, fmt.Errorf("Title cannot be empty!")
	}

	if oldTodo.Title != todo.Title {
		oldTodo.Title = todo.Title
		updated = true
	}
	if oldTodo.Description != todo.Description {
		oldTodo.Description = todo.Description
		updated = true
	}
	if oldTodo.Done != todo.Done {
		oldTodo.Done = todo.Done
		updated = true
	}
	if updated {
		oldTodo.UpdatedAt = time.Now()
	}

	r.todos[index] = oldTodo
	return oldTodo, nil
}

func (r *InMemoryRepo) Delete(id int) error {
	index, err := r.findIndexByID(id)
	if err != nil {
		return err
	}
	r.todos = slices.Delete(r.todos, index, index+1)
	return nil
}

func (r *InMemoryRepo) findIndexByID(id int) (int, error) {
	if id <= 0 {
		return 0, fmt.Errorf("Bad ID %d!", id)
	}

	for i := 0; i < len(r.todos); i++ {
		if id == r.todos[i].ID {
			return i, nil
		}
	}
	return 0, fmt.Errorf("Todo with ID %d not found!", id)
}
