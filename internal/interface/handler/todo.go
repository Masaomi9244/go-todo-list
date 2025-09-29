package handler

import (
	"fmt"
	"go-todo-list/internal/domain"
	usecase "go-todo-list/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HTTP ハンドラの集まり。
// usecase を保持して、ハンドラからビジネスロジックを呼び出す。
type TodoHandler struct {
	uc *usecase.TodoUsecase
}

// TodoHandler のコンストラクタ。
func NewTodoHandler(uc *usecase.TodoUsecase) *TodoHandler {
	return &TodoHandler{uc: uc}
}

// すべての Todo を取得して JSON で返すハンドラ関数。
func (h *TodoHandler) GetTodos(c *gin.Context) {
	todos, err := h.uc.GetAllTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "データの取得に失敗しました"})
		return
	}
	c.JSON(http.StatusOK, todos)
}

// 入力値のDTO（データ転送用の入れ物）
type CreateTodoRequest struct {
	Title       string `json:"title" binding:"required,min=1,max=100"`
	Description string `json:"description" binding:"max=1000"`
	Status      *int   `json:"status,omitempty"`
}

// CreateTodoRequest を domain.Todo に変換するメソッド。
func (req *CreateTodoRequest) ToDomain() (domain.Todo, error) {
	status := domain.NotStarted
	if req.Status != nil {
		if *req.Status < 0 || *req.Status > 2 {
			return domain.Todo{}, fmt.Errorf("ステータスは0から2の整数で指定してください")
		}
		status = domain.Status(*req.Status)
	}

	return domain.Todo{
		Title:       req.Title,
		Description: req.Description,
		Status:      status,
	}, nil
}

// 新しい Todo を作成するハンドラ関数。
func (h *TodoHandler) PostTodo(c *gin.Context) {
	var req CreateTodoRequest

	// jsonのパースとバリデーション
	if err := c.ShouldBindJSON(&req); err != nil {
		// バリデーションエラーの場合はstatus400を返す
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効な入力です"})
		return
	}

	todo, err := req.ToDomain()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ユースケースを呼び出して Todo を作成
	created, err := h.uc.CreateTodo(&todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Todoの作成に失敗しました"})
		return
	}
	c.JSON(http.StatusCreated, created)
}
