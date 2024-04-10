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
