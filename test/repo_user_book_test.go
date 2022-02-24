package test

import (
	"testing"
	"time"

	"bookman/model"
	"bookman/repo"

	"github.com/stretchr/testify/assert"
)

func TestRepoUserBookInsert(t *testing.T) {
	db := OpenDBConn()
	defer CloseDBConn(db)

	res, err := repo.NewUserBookImpl(db).Insert(model.UserBook{
		UserID:    100,
		BookID:    100,
		PageRead:  100,
		IsFinish:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestRepoUserBookFindById(t *testing.T) {
	db := OpenDBConn()
	defer CloseDBConn(db)

	res, err := repo.NewUserBookImpl(db).FindById(100)

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestRepoUserBookUpdate(t *testing.T) {
	db := OpenDBConn()
	defer CloseDBConn(db)

	res, err := repo.NewUserBookImpl(db).Update(model.UserBook{
		ID:        102,
		UserID:    70,
		BookID:    70,
		PageRead:  70,
		IsFinish:  false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestRepoUserBookDelete(t *testing.T) {
	db := OpenDBConn()
	defer CloseDBConn(db)

	err := repo.NewUserBookImpl(db).Delete(model.UserBook{
		ID:        101,
		UserID:    0,
		BookID:    0,
		PageRead:  0,
		IsFinish:  false,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	})

	assert.Nil(t, err)
}

func TestRepoUserBookFindAll(t *testing.T) {
	db := OpenDBConn()
	defer CloseDBConn(db)

	res, err := repo.NewUserBookImpl(db).FindAll()

	assert.Nil(t, err)
	assert.NotNil(t, res)
}
