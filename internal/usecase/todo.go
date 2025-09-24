package usecase

import "go-todo-list/internal/domain"

// usecase 層で利用するリポジトリのインターフェースを定義する。
// infra 層でこのインターフェースを実装し、usecase 層に注入することで、usecase 層は具体的なデータアクセスの実装に依存しなくなる。
type TodoRepository interface {
	FindAll() ([]domain.Todo, error)
}

// Todo に関するユースケースを実装する構造体。
type TodoUsecase struct {
	repo TodoRepository
}

// TodoUsecase のコンストラクタ
// TodoRepository インターフェースを実装したリポジトリを受け取る。
func NewTodoUsecase(repo TodoRepository) *TodoUsecase {
	return &TodoUsecase{repo: repo}
}

// すべての Todo を取得するユースケース。
func (u *TodoUsecase) GetAllTodos() ([]domain.Todo, error) {
	return u.repo.FindAll()
}
