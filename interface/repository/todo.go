package repository

// modelを定義した後、そのデータを永続化するための処理を実装する
// 永続化する処理を担うのが、interface層のrepository

import (
	"try-go-react-todo-back/database/model"
	"try-go-react-todo-back/domain/entity"
	"try-go-react-todo-back/usecase/port"

	"gorm.io/gorm"
)

type TodoRepository struct {
	// gorm.DBインスタンスを保持
	db *gorm.DB
}

// TodoRepositoryのコンストラクタ
func NewTodoRepository(db *gorm.DB) port.TodoRepository {
	return &TodoRepository{db: db}
}

func (t TodoRepository) Todo() ([]entity.Todo, error) {
	var todoModels []model.Todo
	if err := t.db.Find(&todoModels).Error; err != nil {
		return nil, err
	}

	// sliceを用意（この時点では空）
	todos := make([]entity.Todo, len(todoModels))
	for i, todoModel := range todoModels {
		// slice「todos」に値を入れていく
		todos[i] = todoModel.Entity()
	}

	return todos, nil
}

func (t TodoRepository) AddTodo(todo entity.Todo) error {
	todoModel := &model.Todo{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		Tag:         todo.Tag,
		Level:       todo.Level,
	}

	return t.db.Create(todoModel).Error
}

func (t TodoRepository) DeleteTodo(id string) error {
	todoModel := &model.Todo{}
	return t.db.Delete(todoModel, id).Error
}

func (t TodoRepository) DeleteTodos(ids []string) error {
	todoModel := &model.Todo{}
	return t.db.Delete(todoModel, ids).Error
}

func (t TodoRepository) FindTodo(id string) (entity.Todo, error) {
	todoModel := &model.Todo{}

	if err := t.db.Where("id = ?", id).First(&todoModel).Error; err != nil {
		return entity.Todo{}, err
	}

	return todoModel.Entity(), nil
}

func (t TodoRepository) UpdateTodo(todo entity.Todo) error {
	todoModel := &model.Todo{}
	updateData := model.Todo{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		Status:      todo.Status,
		Level:       todo.Level,
		Tag:         todo.Tag,
	}

	return t.db.Model(todoModel).Where("id = ?", todo.ID).Updates(updateData).Error
}
