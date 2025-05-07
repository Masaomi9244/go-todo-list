package infra

import (
	"go-todo-list/internal/domain"
)

type TodoRepository struct{}

func (r *TodoRepository) Save(todo *domain.Todo) error {
	return DB.Create(todo).Error
}
