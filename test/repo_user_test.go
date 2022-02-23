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
	user   model.User
	userID int
)

func TestRepoUserInsertSuccess(t *testing.T) {
	db := OpenDBConn()
	defer CloseDBConn(db)

	res, err := repo.NewUserImpl(db).Insert(model.User{
		Name:      "Muhammmad Handi Rachmawan",
		Email:     "example@mail.com",
		Password:  "katasandi",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	assert.Nil(t, err)
	assert.NotNil(t, res)

	user = res
	userID = res.ID
}

func TestRepoUserFindByIdSuccess(t *testing.T) {
	db := OpenDBConn()
	defer CloseDBConn(db)

	res, err := repo.NewUserImpl(db).FindById(userID)

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestRepoUserFindByIdFailed(t *testing.T) {
	db := OpenDBConn()
	defer CloseDBConn(db)

	res, err := repo.NewUserImpl(db).FindById(userID + 10)
	assert.NotNil(t, err)

	err = validator.New().Struct(res)
	assert.NotNil(t, err)
}

func TestRepoUserIsEmailDuplicate(t *testing.T) {
	db := OpenDBConn()
	defer CloseDBConn(db)

	res, err := repo.NewUserImpl(db).IsEmailDuplicate("example@mail.com")

	assert.Nil(t, err)
	assert.True(t, res)
}

func TestRepoUserIsEmailNotDuplicate(t *testing.T) {
	db := OpenDBConn()
	defer CloseDBConn(db)

	res, err := repo.NewUserImpl(db).IsEmailDuplicate("another@mail.com")

	assert.NotNil(t, err)
	assert.False(t, res)
}

func TestRepoUserUpdate(t *testing.T) {
	db := OpenDBConn()
	defer CloseDBConn(db)

	user = model.User{
		ID:        user.ID,
		Name:      "Name Update",
		Email:     "Email Update",
		Password:  "Password Update",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	res, err := repo.NewUserImpl(db).Update(user)

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestRepoUserFindAll(t *testing.T) {
	db := OpenDBConn()
	defer CloseDBConn(db)

	res, err := repo.NewUserImpl(db).FindAll()

	assert.Nil(t, err)
	assert.NotNil(t, res)
}
