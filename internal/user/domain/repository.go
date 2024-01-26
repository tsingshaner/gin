package domain

import "github.com/lab-online/internal/user/entity"

type UserRepository interface {
	CheckUserExists(string) (bool, error)
	InsertUser(userEntity entity.UserEntity) error
}
