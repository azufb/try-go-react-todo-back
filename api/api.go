package api

import "gorm.io/gorm"

type TodoAPI struct {
	db *gorm.DB
}

func NewTodoAPI(db *gorm.DB) *TodoAPI {
	return &TodoAPI{db: db}
}
