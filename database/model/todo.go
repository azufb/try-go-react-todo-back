package model

// テーブルの構成を定義
// テーブル名：todos
// idカラム・titleカラム・statusカラムがある
// primary keyはid
type Todo struct {
	ID     string `gorm:"primaryKey;size:26;not null"`
	Title  string `gorm:"size:255;not null"`
	Status string `gorm:"size:255;not null"`
}
