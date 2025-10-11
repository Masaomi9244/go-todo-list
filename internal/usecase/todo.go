package usecase

import "go-todo-list/internal/domain"

// usecase 層で利用するリポジトリのインターフェースを定義する。
// infra 層でこのインターフェースを実装し、usecase 層に注入することで、usecase 層は具体的なデータアクセスの実装に依存しなくなる。
type TodoRepository interface {
	// すべての Todo を取得するメソッド。
	FindAll() ([]domain.Todo, error)
	// 新しい Todo を1件保存するメソッド。
	Create(todo *domain.Todo) error
	// IDに基づいて Todo を取得するメソッド。
	FindById(int uint) (*domain.Todo, error)
	// 既存の Todo を更新するメソッド。
	Update(todo *domain.Todo) error
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

// 新しい Todo を作成するユースケース。
func (u *TodoUsecase) CreateTodo(todo *domain.Todo) (*domain.Todo, error) {
	if err := u.repo.Create(todo); err != nil {
		return nil, err
	}
	return todo, nil
}

// ID に基づいて既存の Todo を更新するユースケース。
func (u *TodoUsecase) UpdateTodo(id uint, updated *domain.Todo) (*domain.Todo, error) {
	// infra層のリポジトリ経由でTodoを取得
	todo, err := u.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	// 更新内容を反映
	todo.Title = updated.Title
	todo.Description = updated.Description
	todo.Status = updated.Status

	// DBを更新
	if err := u.repo.Update(todo); err != nil {
		return nil, err
	}

	return todo, nil
}
