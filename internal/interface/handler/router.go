package handler

import (
	"go-todo-list/internal/infra"

	"go-todo-list/internal/usecase"

	"github.com/gin-gonic/gin"
)

// SetupRouter は Gin のルーターを初期化して返す関数。
// このルーターにはログ出力とエラー回復の機能（ミドルウェア）が最初から含まれている。
func SetupRouter() *gin.Engine {
	// Gin のデフォルトルーターを作成する。
	// ロガー（リクエストログを自動出力）とリカバリー（エラー時でもサーバーが落ちない）が有効になっている。
	r := gin.Default()

	// 依存関係の注入
	// ここでリポジトリ、ユースケース、ハンドラを順に初期化し、依存関係を注入していく。
	// これにより各層が疎結合になり、テストや保守が容易になる。
	// 具体的には、infra 層のリポジトリを usecase 層に渡し、usecase 層のユースケースを handler 層に渡す。
	// こうすることで、handler 層は usecase 層のインターフェースにのみ依存し、具体的な実装には依存しなくなる。
	repo := infra.NewTodoRepository()
	uc := usecase.NewTodoUsecase(repo)
	h := NewTodoHandler(uc)

	// GET /todos にアクセスがあったとき、GetTodos 関数が呼び出されるようにルーティングを設定する。
	r.GET("/todos", h.GetTodos)
	// POST /todos にアクセスがあったとき、CreateTodo 関数が呼び出されるようにルーティングを設定する。
	r.POST("/todos", h.CreateTodo)

	// 初期化したルーターを返す。
	return r
}
