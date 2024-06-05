package model

import (
	"github.com/tsingshaner/gin-starter/pkg/model"
)

type User struct {
	model.BaseModel
	UserID   string `gorm:"unique"`
	Username string
	Password string
	Role     uint8
}
