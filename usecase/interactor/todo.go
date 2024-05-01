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
func (t TodoUsecase) AddTodo(id string, title string, description string, tag string) error {
	// entity.Todoに入れる
	todo := entity.Todo{
		ID:          id,
		Title:       title,
		Description: description,
		Tag:         tag,
	}
	// AddTodoメソッドを呼び出し
	return t.TodoRepository.AddTodo(todo)
}

// DeleteTodoメソッド
func (t TodoUsecase) DeleteTodo(id string) error {
	return t.TodoRepository.DeleteTodo(id)
}

func (t TodoUsecase) DeleteTodos(ids []string) error {
	return t.TodoRepository.DeleteTodos(ids)
}

func (t TodoUsecase) FindTodo(id string) (entity.Todo, error) {
	return t.TodoRepository.FindTodo(id)
}

func (t TodoUsecase) UpdateTodo(id string, title string, description string, status string, tag string, level string) error {
	todo := entity.Todo{
		ID:          id,
		Title:       title,
		Description: description,
		Status:      status,
		Tag:         tag,
		Level:       level,
	}
	return t.TodoRepository.UpdateTodo(todo)
}
