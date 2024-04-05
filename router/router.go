package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware" // バージョン不一致でCORSエラー吐くのでv4指定
)

// Helo Worldを返すAPI
// https://pkg.go.dev/github.com/labstack/echo
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!")
}

func NewServer() *echo.Echo {
	e := echo.New()

	// middleware
	e.Use(middleware.CORS())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet},
	}))

	e.GET("/hello", hello)

	return e
}
