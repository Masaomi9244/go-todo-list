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

// IDに基づいて Todo を取得するメソッド。
func (r *TodoRepository) FindById(id uint) (*domain.Todo, error) {
	var todo domain.Todo
	if err := DB.First(&todo, id).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

// 既存の Todo を更新するメソッド。
func (r *TodoRepository) Update(todo *domain.Todo) error {
	return DB.Save(todo).Error
}
