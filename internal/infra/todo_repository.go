package infra

import "go-todo-list/internal/domain"

// infra 層で usecase 層の TodoRepository インターフェースを実装する。
type TodoRepository struct{}

// NewTodoRepository は TodoRepository のコンストラクタ。
func NewTodoRepository() *TodoRepository {
	return &TodoRepository{}
}

// FindAll はすべての Todo をデータベースから取得する。
func (r *TodoRepository) FindAll() ([]domain.Todo, error) {
	var todos []domain.Todo
	if err := DB.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}
