package port

import "try-go-react-todo-back/domain/entity"

// interface
// メソッド型宣言の集合・どのようなメソッドを実装するかを示す仕様書
type TodoRepository interface {
	Todo() ([]entity.Todo, error)
}
