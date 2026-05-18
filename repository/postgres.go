package repository

import (
	"github.com/SharathKumarK06/todo-server/models"
	"gorm.io/gorm"
)

type PostgresRepo struct {
	db *gorm.DB
}

func NewPostgresRepo(db *gorm.DB) *PostgresRepo {
	return &PostgresRepo{
		db: db,
	}

}

func (r *PostgresRepo) Create(todo models.Todo) (models.Todo, error) {
	err := r.db.Create(&todo).Error
	if err != nil {
		return models.Todo{}, err
	}

	return todo, nil
}

func (r *PostgresRepo) List() ([]models.Todo, error) {
	var todos []models.Todo
	err := r.db.Find(&todos).Error
	if err != nil {
		return []models.Todo{}, err
	}

	return todos, err
}

func (r *PostgresRepo) GetByID(id int) (models.Todo, error) {
	var todo models.Todo
	err := r.db.First(&todo, id).Error
	if err != nil {
		return models.Todo{}, err
	}

	return todo, nil
}

func (r *PostgresRepo) Update(id int, todo models.Todo) (models.Todo, error) {
	err := r.db.Save(&todo).Error
	if err != nil {
		return models.Todo{}, err
	}

	return todo, nil
}

func (r *PostgresRepo) Delete(id int) error {
	err := r.db.Delete(&models.Todo{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
