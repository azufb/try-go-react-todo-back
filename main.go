package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// DBつないでToDoアプリを実装したい
// DBからのデータ取得・DBにデータ挿入・DBのデータ編集・DBのデータ削除をやる

func todos() error {
	// 取得
	return nil
}

func create() error {
	// データ挿入
	return nil
}

func update() error {
	// データ編集
	return nil
}

func delete() error {
	// データ削除
	return nil
}

func main() {
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

	if err := e.Start(":8080"); err != nil {
		os.Exit(1)
	}
}
