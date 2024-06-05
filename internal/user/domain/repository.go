package domain

import "github.com/tsingshaner/gin-starter/internal/user/entity"

type UserRepository interface {
	CheckUserExists(string) (bool, error)
	InsertUser(entity.UserEntity) error
	FindByUserID(string) (entity.UserEntity, error)
}
