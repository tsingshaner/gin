package domain

import "github.com/lab-online/internal/user/entity"

type UserRepository interface {
	CheckUserExists(string) (bool, error)
	InsertUser(*entity.User) error
}
