package test

import (
	"testing"
	"time"

	"bookman/model"
	"bookman/repo"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

var (
	book   model.Book
	bookID int
)

func TestRepoBookInsertSuccess(t *testing.T) {
	db := OpenDBConn()
	defer CloseDBConn(db)

	res, err := repo.NewBookImpl(db).Insert(model.Book{
		Title:     "Laskar Pelangi",
		Author:    "Andrea Hirata",
		Year:      2006,
		Publisher: "Gramedia",
		PageCount: 480,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	assert.Nil(t, err)
	assert.NotNil(t, res)

	book = res
	bookID = res.ID
}

func TestRepoBookFindByIdSuccess(t *testing.T) {
	db := OpenDBConn()
	defer CloseDBConn(db)

	res, err := repo.NewBookImpl(db).FindById(bookID)

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestRepoBookFindByIdFailed(t *testing.T) {
	db := OpenDBConn()
	defer CloseDBConn(db)

	res, err := repo.NewBookImpl(db).FindById(bookID + 10)
	assert.NotNil(t, err)

	err = validator.New().Struct(res)
	assert.Nil(t, err)
}

func TestRepoBookUpdate(t *testing.T) {
	db := OpenDBConn()
	defer CloseDBConn(db)

	book = model.Book{
		ID:        book.ID,
		Title:     "Buku Baru",
		Author:    "Buku Baru",
		Year:      2022,
		Publisher: "Buku Baru",
		PageCount: 2022,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	res, err := repo.NewBookImpl(db).Update(book)
	assert.Nil(t, err)

	err = validator.New().Struct(res)
	assert.Nil(t, err)
}

func TestRepoBookDeleteSuccess(t *testing.T) {
	db := OpenDBConn()
	defer CloseDBConn(db)

	err := repo.NewBookImpl(db).Delete(book)

	assert.Nil(t, err)
}

func TestRepoBookDeleteFailed(t *testing.T) {
	db := OpenDBConn()
	defer CloseDBConn(db)

	err := repo.NewBookImpl(db).Delete(model.Book{})

	assert.NotNil(t, err)
}

func TestRepoBookFindAll(t *testing.T) {
	db := OpenDBConn()
	defer CloseDBConn(db)

	res, err := repo.NewBookImpl(db).FindAll()

	assert.Nil(t, err)
	assert.NotNil(t, res)
}
