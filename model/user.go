package model

import "time"

const (
	UserAdmin  = "admin"
	UserMember = "member"
)

type User struct {
	ID        int       `faker:"-"`
	Name      string    `gorm:"type:varchar(100);not null" validate:"required" faker:"name"`
	Email     string    `gorm:"type:varchar(100);not null:unique" validate:"required" faker:"email"`
	Password  string    `gorm:"type:varchar(100);not null" validate:"required" faker:"password"`
	Role      string    `gorm:"type:varchar(6);not null" validate:"required" faker:"oneof: admin, member"`
	CreatedAt time.Time `gorm:"not null" faker:"-"`
	UpdatedAt time.Time `gorm:"not null" faker:"-"`
	Books     []Book    `gorm:"many2many:user_book;" faker:"-"`
}

func (User) TableName() string {
	return "user"
}
