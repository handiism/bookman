package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func OpenDatabaseConnection() *gorm.DB {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	address := os.Getenv("DB_ADDRESS")
	name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		address,
		name,
	)
	
	db, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic(err)
	}

	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	conn, err := db.DB()
	if err != nil {
		panic(err)
	}

	conn.Close()
}
