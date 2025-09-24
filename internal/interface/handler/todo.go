package handler

import (
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
