package model

import (
	baseModel "github.com/lab-online/pkg/model"
)

type User struct {
	baseModel.BaseModel
	UserID   string `gorm:"unique"`
	Password string
	Role     int8
}
