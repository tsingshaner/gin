package model

import (
	"github.com/lab-online/pkg/model"
)

type User struct {
	model.BaseModel
	UserID   string `gorm:"unique"`
	Username string
	Password string
	Role     uint8
}
