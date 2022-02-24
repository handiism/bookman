package model

import "time"

type UserBook struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	UserID    int       `gorm:"foreignKey"`
	BookID    int       `gorm:"foreignKey"`
	PageRead  int       `gorm:"type:int(10);not null;default:0"`
	IsFinish  bool      `gorm:"default:false;not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

func (UserBook) TableName() string {
	return "user_book"
}
