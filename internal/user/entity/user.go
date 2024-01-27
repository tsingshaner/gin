package entity

import (
	"database/sql"
	"time"
)

type User struct {
	ID        uint         `json:"id"`
	UserID    string       `json:"userId"`
	Username  string       `json:"username"`
	Password  string       `json:"password"`
	Role      uint8        `json:"role"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt"`
	DeletedAt sql.NullTime `json:"deletedAt"`
}

type UserEntity interface {
	GetUserID() string
	GetUsername() string
	GetPassword() string
	GetRole() uint8

	HashPassword() error
	ComparePassword(password string) bool
}
