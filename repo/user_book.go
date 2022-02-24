package repo

import (
	"bookman/model"
	"gorm.io/gorm"
)

type UserBook interface {
	Insert(model.UserBook) (model.UserBook, error)
	FindById(int) (map[string]interface{}, error)
	Update(model.UserBook) (model.UserBook, error)
	Delete(model.UserBook) error
	FindAll() ([]model.UserBook, error)
}

type userBook struct {
	db *gorm.DB
}

func NewUserBookImpl(db *gorm.DB) UserBook {
	return &userBook{db}
}

func (repo *userBook) Insert(userBook model.UserBook) (model.UserBook, error) {
	err := repo.db.Create(&userBook).Error

	if err != nil {
		return model.UserBook{}, err
	}

	return userBook, nil
}

func (repo *userBook) FindById(id int) (map[string]interface{}, error) {
	var userBook map[string]interface{}
	err := repo.db.Raw(
		`SELECT user_book.id,
				user_book.user_id,
				user_book.book_id,
				user_book.page_read,
				book.page_count,
				user_book.is_finish,
				user_book.created_at,
				user_book.updated_at
 		 FROM user_book
		 			 JOIN book ON user_book.book_id = book.id
 		 WHERE user_book.id = ?`, id,
	).Scan(&userBook).Error

	if err != nil {
		return map[string]interface{}{}, err
	}

	return userBook, nil
}

func (repo *userBook) Update(userBook model.UserBook) (model.UserBook, error) {
	err := repo.db.Save(&userBook).Error

	if err != nil {
		return model.UserBook{}, err
	}

	return userBook, nil
}

func (repo *userBook) Delete(book model.UserBook) error {
	err := repo.db.Delete(book).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *userBook) FindAll() ([]model.UserBook, error) {
	var userBooks []model.UserBook

	err := repo.db.Find(&userBooks).Error

	if err != nil {
		return nil, err
	}

	return userBooks, nil
}
