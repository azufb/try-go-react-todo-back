package entity

// Todo
// システム内で使う構造体
// model＝テーブル定義
// entity＝システム内で使うデータの型
type Todo struct {
	ID          string
	Title       string
	Status      string
	Description string
	Level       string
	Tag         string
}
