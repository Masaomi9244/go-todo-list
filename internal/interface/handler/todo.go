package handler

import (
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

func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var req CreateTodoRequest

	// jsonのパースとバリデーション
	if err := c.ShouldBindJSON(&req); err != nil {
		// バリデーションエラーの場合はstatus400を返す
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効な入力です"})
		return
	}

	// ステータスのデフォルト値の設定とバリデーションチェック
	status := domain.NotStarted
	if req.Status != nil {
		if *req.Status < 0 || *req.Status > 2 {
			// ステータスが不正な場合はstatus400を返す
			c.JSON(http.StatusBadRequest, gin.H{"error": "ステータスは0から2の整数で指定してください"})
			return
		}
		status = domain.Status(*req.Status)
	}

	// ドメインモデルにリクエストの内容を詰める
	todo := &domain.Todo{
		Title:       req.Title,
		Description: req.Description,
		Status:      status,
	}

	// ユースケースを呼び出して Todo を作成
	created, err := h.uc.CreateTodo(todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Todoの作成に失敗しました"})
		return
	}
	c.JSON(http.StatusCreated, created)
}
