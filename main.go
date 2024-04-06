package main

import (
	"fmt"
	"os"
	"try-go-react-todo-back/api"
	"try-go-react-todo-back/api/router"
	"try-go-react-todo-back/database"
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
	// データベース接続
	db, err := database.NewDB()
	if err != nil {
		fmt.Errorf("failed to connect to database: %w", err)
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

	todoApi := api.NewTodoAPI(db)

	server := router.NewServer(*todoApi)
	if err := server.Start(":8080"); err != nil {
		fmt.Errorf("failed to start server: %w", err)
		os.Exit(1)
	}
}
