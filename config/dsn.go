package config

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/rds/auth"
	_ "github.com/go-sql-driver/mysql"
)

func DSN() string {
	userName := os.Getenv("DB_MASTER_USER_NAME")
	// password := os.Getenv("DB_MASTER_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	// dbUrl := os.Getenv("DB_URL")
	dbEndpoint := fmt.Sprintf("%s:%s", host, port)
	region := os.Getenv("DB_REGION")

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error: " + err.Error())
	}

	authenticationToken, err := auth.BuildAuthToken(
		context.TODO(), dbEndpoint, region, userName, cfg.Credentials)
	if err != nil {
		panic("failed to create authentication token:" + err.Error())
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&tls=true&allowCleartextPasswords=true&parseTime=True&loc=Local",
		userName, authenticationToken, dbEndpoint, dbName,
	)

	return dsn
}
