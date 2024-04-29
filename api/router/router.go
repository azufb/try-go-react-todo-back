package router

import (
	"net/http"
	"try-go-react-todo-back/api/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewServer(todoHandler handler.TodoHandler) *echo.Echo {
	e := echo.New()

	// middleware
	e.Use(middleware.CORS())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet},
	}))

	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello!")
	})

	e.GET("/todo", todoHandler.Todo)

	e.POST("/addTodo", todoHandler.AddTodo)

	e.DELETE("/deleteTodo", todoHandler.DeleteTodo)

	e.DELETE("/deleteTodos", todoHandler.DeleteTodos)

	e.POST("/findTodo", todoHandler.FindTodo)

	e.POST("/updateTodo", todoHandler.UpdateTodo)

	return e
}
