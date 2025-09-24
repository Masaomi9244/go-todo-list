package main

import (
	"go-todo-list/internal/infra"
	"go-todo-list/internal/interface/handler"
)

func main() {
	// DB初期化
	infra.InitDB()

	// Ginルーター初期化
	r := handler.SetupRouter()

	// サーバー起動
	r.Run(":8080")
}
