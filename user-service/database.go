package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// CreateConnection with postgres using gorm
func CreateConnection() (*gorm.DB, error) {
	// Get database details from env var
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	return gorm.Open(
		"postgres",
		fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", host, user, dbName, password),
	)
}
