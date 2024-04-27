package model

import "try-go-react-todo-back/domain/entity"

// テーブルの構成を定義
// テーブル名：todos
// idカラム・titleカラム・statusカラムがある
// primary keyはid
type Todo struct {
	ID          string `gorm:"primaryKey;size:255;not null"`
	Title       string `gorm:"size:255;not null"`
	Status      string `gorm:"type: enum('NOT_STARTED', 'IN_PROGRESS', 'DONE'); default: 'NOT_STARTED'; not null"`
	Description string `gorm:"size:255; not null"`
	Tag         string `gorm:"size:255; not null"`
	Level       string `gorm:"type: enum('urgent', 'normal', 'low'); default: 'normal'; not null"`
}

// Entityメソッド
// modelの型からentityの型（システム内での型）に変換する
func (t *Todo) Entity() entity.Todo {
	return entity.Todo{
		ID:          t.ID,
		Title:       t.Title,
		Status:      t.Status,
		Description: t.Description,
		Tag:         t.Tag,
		Level:       t.Level,
	}
}
