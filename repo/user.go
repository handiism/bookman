package repo

import (
	"bookman/model"
	"gorm.io/gorm"
)

type User interface {
	Insert(model.User) (model.User, error)
	FindById(id int) (model.User, error)
	IsEmailDuplicate(email string) (bool, error)
	Update(user model.User) (model.User, error)
	FindAll() ([]model.User, error)
}

type user struct {
	db *gorm.DB
}

func NewUserImpl(db *gorm.DB) User {
	return &user{db}
}

func (repo *user) Insert(user model.User) (model.User, error) {
	err := repo.db.Create(&user).Error

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (repo *user) FindById(id int) (model.User, error) {
	var user model.User

	err := repo.db.First(&user, id).Error

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (repo *user) IsEmailDuplicate(email string) (bool, error) {
	err := repo.db.Where(&model.User{Email: email}).First(&model.User{}).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

func (repo *user) Update(user model.User) (model.User, error) {
	err := repo.db.Save(&user).Error

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (repo *user) FindAll() ([]model.User, error) {
	var users []model.User

	err := repo.db.Find(&users).Error

	if err != nil {
		return []model.User{}, err
	}

	return users, nil
}
