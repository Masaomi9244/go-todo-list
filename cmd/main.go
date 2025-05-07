package main

import (
	"fmt"
	"go-todo-list/internal/domain"
	"go-todo-list/internal/infra"
)

func main() {
	infra.InitDB()

	todo := domain.Todo{
		Title:       "ステータス付きタグ",
		Description: "ステータス付きタグの説明",
		Status:      domain.NotStarted,
	}

	repo := infra.TodoRepository{}
	err := repo.Save(&todo)
	if err != nil {
		panic(err)
	}
	fmt.Println("✅ Todo 登録完了！")
}
