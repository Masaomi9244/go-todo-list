package handler

import "github.com/gin-gonic/gin"

// SetupRouter は Gin のルーターを初期化して返す関数。
// このルーターにはログ出力とエラー回復の機能（ミドルウェア）が最初から含まれている。
func SetupRouter() *gin.Engine {
	// Gin のデフォルトルーターを作成する。
	// ロガー（リクエストログを自動出力）とリカバリー（エラー時でもサーバーが落ちない）が有効になっている。
	r := gin.Default()

	// GET /todos にアクセスがあったとき、GetTodos 関数が呼び出されるようにルーティングを設定する。
	r.GET("/todos", GetTodos)

	// 初期化したルーターを返す。
	return r
}
