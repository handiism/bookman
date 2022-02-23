package repo

import (
	"bookman/model"
	"gorm.io/gorm"
)

type UserBook interface {
	FindAll() (map[string]interface{}, error)
	Insert(model.UserBook) (model.UserBook, error)
}

type userBook struct {
	db *gorm.DB
}

func NewUserBookImpl(db *gorm.DB) UserBook {
	return &userBook{db}
}

func (repo *userBook) FindAll() (map[string]interface{}, error) {
	var userBooks map[string]interface{}

	err := repo.db.Model(model.UserBook{}).Take(&userBooks).Error

	if err != nil {
		return nil, err
	}

	return userBooks, nil
}

func (repo *userBook) Insert(userBook model.UserBook) (model.UserBook, error) {
	err := repo.db.Create(&userBook).Error

	if err != nil {
		return model.UserBook{}, err
	}

	return userBook, nil
}
