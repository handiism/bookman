package model

import "time"

type Book struct {
	ID        int       `faker:"-"`
	Title     string    `gorm:"type:varchar(100);not null" faker:"word"`
	Author    string    `gorm:"type:varchar(100);not null" faker:"name"`
	Year      int       `gorm:"type:int(4);not null"`
	Publisher string    `gorm:"type:varchar(100);not null" faker:"word"`
	PageCount int       `gorm:"type:int(10);not null"`
	CreatedAt time.Time `gorm:"not null" faker:"-"`
	UpdatedAt time.Time `gorm:"not null" faker:"-"`
}

func (Book) TableName() string {
	return "book"
}
