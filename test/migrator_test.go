package test

import (
	"bookman/model"
	"bookman/repo"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/bxcodec/faker/v3"
)

func TestMigrateTable(*testing.T) {
	db := OpenDBConn()
	defer CloseDBConn(db)

	db.AutoMigrate(model.Book{}, model.User{}, model.UserBook{})
}

func TestMigrateData(*testing.T) {
	db := OpenDBConn()
	defer CloseDBConn(db)

	for i := 0; i < 100; i++ {
		user := model.User{}
		faker.FakeData(&user)
		user.CreatedAt = time.Now()
		user.UpdatedAt = user.CreatedAt
		repo.NewUserImpl(db).Insert(user)

		book := model.Book{}
		faker.FakeData(&book)
		book.Year, _ = strconv.Atoi(faker.YEAR)
		book.CreatedAt = time.Now()
		book.UpdatedAt = book.CreatedAt
		repo.NewBookImpl(db).Insert(book)
	}

	for i := 0; i < 100; i++ {
		userBook := model.UserBook{
			UserID:    rand.Intn(99) + 1,
			BookID:    rand.Intn(99) + 1,
			PageRead:  0,
			IsFinish:  false,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		repo.NewUserBookImpl(db).Insert(userBook)
	}
}
