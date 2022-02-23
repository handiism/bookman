package repo

import (
	"bookman/model"
	"gorm.io/gorm"
)

type Book interface {
	Insert(model.Book) (model.Book, error)
	FindById(id int) (model.Book, error)
	Update(model.Book) (model.Book, error)
	Delete(model.Book) error
	FindAll() ([]model.Book, error)
}

type book struct {
	db *gorm.DB
}

func NewBookImpl(db *gorm.DB) Book {
	return &book{db}
}

func (repo *book) Insert(book model.Book) (model.Book, error) {
	err := repo.db.Create(&book).Error

	if err != nil {
		return model.Book{}, nil
	}

	return book, nil
}

func (repo *book) FindById(id int) (model.Book, error) {
	var book model.Book

	err := repo.db.First(&book, id).Error

	if err != nil {
		return model.Book{}, err
	}

	return book, nil
}

func (repo *book) Update(book model.Book) (model.Book, error) {
	err := repo.db.Save(&book).Error

	if err != nil {
		return model.Book{}, err
	}

	return model.Book{}, nil
}

func (repo *book) Delete(book model.Book) error {
	err := repo.db.Delete(&book).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *book) FindAll() ([]model.Book, error) {
	var books []model.Book

	err := repo.db.Find(&books).Error

	if err != nil {
		return []model.Book{}, err
	}

	return books, nil
}
