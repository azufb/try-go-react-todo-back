package database

import (
	"fmt"
	"math"
	"time"
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

	// Check connection
	const retryMax = 10
	for i := 0; i < retryMax; i++ {
		err = sqlDB.Ping()
		if err == nil {
			break
		}
		if i == retryMax-1 {
			return nil, fmt.Errorf("failed to connect to database: %w", err)
		}
		duration := time.Millisecond * time.Duration(math.Pow(1.5, float64(i))*1000)
		time.Sleep(duration)
	}

	return db, nil
}

// マイグレーション
func Migrate(db *gorm.DB) error {
	if err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Todo{}); err != nil {
		return err
	}

	return nil
}
