package helper

import (
	"bookman/model"
	"gorm.io/gorm"
)

func DatabaseMigration(db *gorm.DB) {
	err := db.AutoMigrate(
		model.User{},
		model.Book{},
		model.UserBook{},
	)

	if err != nil {
		panic(err)
	}
}
