package database

import (
	"fmt"
	"try-go-react-todo-back/config"
	"try-go-react-todo-back/database/model"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql" // これがBroken Importになるので上のやつ入れる
	"gorm.io/gorm"
)

// DB接続
func NewDB() (*gorm.DB, error) {
	dsn := config.DSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open MySQL: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql.DB: %w", err)
	}

	sqlDB.SetConnMaxIdleTime(100)
	sqlDB.SetMaxOpenConns(100)

	return db, nil
}

// マイグレーション
func Migrate(db *gorm.DB) error {
	if err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Todo{}); err != nil {
		return err
	}

	return nil
}
