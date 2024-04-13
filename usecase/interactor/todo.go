package interactor

import (
	"try-go-react-todo-back/domain/entity"
	"try-go-react-todo-back/usecase/port"
)

// TodoUsecase
// usecase interactor＝portで定義したメソッドを呼び出す
// portで定義したメソッド＝Todoメソッド
type TodoUsecase struct {
	TodoRepository port.TodoRepository
}

// TodoUsecaseのコンストラクタ
func NewTodoUsecase(todoRepository port.TodoRepository) *TodoUsecase {
	return &TodoUsecase{TodoRepository: todoRepository}
}

// Todoメソッド
func (t TodoUsecase) Todo() ([]entity.Todo, error) {
	// Todoメソッドを呼び出し
	return t.TodoRepository.Todo()
}

// AddTodoメソッド
func (t TodoUsecase) AddTodo(id string, title string) error {
	// entity.Todoに入れる
	todo := entity.Todo{
		ID:    id,
		Title: title,
	}
	// AddTodoメソッドを呼び出し
	return t.TodoRepository.AddTodo(todo)
}
