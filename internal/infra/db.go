package infra

import (
	"go-todo-list/internal/domain"
	"go-todo-list/internal/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=todo_user password=password dbname=todo_db port=5432 sslmode=disable"

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	utils.FatalIfErr(err, "DB接続失敗")

	// マイグレーション（todoテーブル作成）
	err = DB.AutoMigrate(&domain.Todo{})
	utils.FatalIfErr(err, "マイグレーション失敗")
}
