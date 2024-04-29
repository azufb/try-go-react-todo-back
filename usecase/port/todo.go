package port

import "try-go-react-todo-back/domain/entity"

// interface
// メソッド型宣言の集合・どのようなメソッドを実装するかを示す仕様書
type TodoRepository interface {
	Todo() ([]entity.Todo, error)
	AddTodo(todo entity.Todo) error
	DeleteTodo(id string) error
	DeleteTodos(ids []string) error
	FindTodo(id string) (entity.Todo, error)
	UpdateTodo(id string) error
}
