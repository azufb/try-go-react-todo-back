package main

import (
	"fmt"
	"os"
	"try-go-react-todo-back/api/handler"
	"try-go-react-todo-back/api/router"
	"try-go-react-todo-back/database"
	"try-go-react-todo-back/interface/repository"
	"try-go-react-todo-back/usecase/interactor"
)

// DBつないでToDoアプリを実装したい
// DBからのデータ取得・DBにデータ挿入・DBのデータ編集・DBのデータ削除をやる

func main() {
	// データベース接続
	db, err := database.NewDB()
	if err != nil {
		fmt.Errorf("failed to connect to database: %w", err)
		return
	}

	if err := database.Migrate(db); err != nil {
		fmt.Errorf("failed to migrate: %w", err)
		return
	}

	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			fmt.Errorf("failed to get sql.db: %w", err)
		}
		err = sqlDB.Close()
		if err != nil {
			fmt.Errorf("failed to close database connection: %w", err)
		}
	}()

	// portを持つrepository初期化
	todoRepository := repository.NewTodoRepository(db)
	// portを引数に受け取ってusecaseを初期化
	todoUC := interactor.NewTodoUsecase(todoRepository)
	// usecaseを引数に受け取ってhandlerを初期化
	todoHandler := handler.NewTodoHandler(todoUC)

	server := router.NewServer(todoHandler)
	if err := server.Start(":8080"); err != nil {
		fmt.Errorf("failed to start server: %w", err)
		os.Exit(1)
	}
}
