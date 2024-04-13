package handler

import (
	"net/http"
	"try-go-react-todo-back/api/schema"
	"try-go-react-todo-back/usecase/interactor"

	"github.com/labstack/echo/v4"
)

// TodoUsecaseを持つ
type TodoHandler struct {
	TodoUC *interactor.TodoUsecase
}

// TodoHandlerのコンストラクタ
func NewTodoHandler(todoUC *interactor.TodoUsecase) TodoHandler {
	return TodoHandler{TodoUC: todoUC}
}

// ここでリクエストを受け取って、usecaseを実行する
// ここでは、Todoメソッドを実行し、結果を返す
func (t *TodoHandler) Todo(c echo.Context) error {
	todos, err := t.TodoUC.Todo()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, schema.TodosResFromEntity(todos))
}

// タスクを追加する
func (t *TodoHandler) AddTodo(c echo.Context) error {
	// request
	req := &schema.TodoReq{}
	// requestをreq変数に格納
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err := t.TodoUC.AddTodo(req.ID, req.Title)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, "Success!!")
}
