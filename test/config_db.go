package test

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func OpenDBConn() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/gorm_api_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Discard.LogMode(logger.Silent),
	})

	if err != nil {
		panic(err)
	}

	return db
}

func CloseDBConn(db *gorm.DB) {
	conn, err := db.DB()

	if err != nil {
		panic(err)
	}

	conn.Close()
}