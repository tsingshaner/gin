package export

import "github.com/tsingshaner/gin-starter/internal/user/entity"

type Repository interface {
	FindUserByID(id uint) (entity.UserEntity, error)
	FindManyUserByID(id []uint) ([]entity.UserEntity, error)
}
