package user

import (
	"github.com/lab-online/internal/user/infra/model"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func New(db *gorm.DB) *User {
	return &User{db}
}

func (u *User) table() *gorm.DB {
	return u.db.Model(&model.User{})
}
