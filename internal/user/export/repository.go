package export

import "github.com/lab-online/internal/user/entity"

type Repository interface {
	FindUserByID(id uint) (entity.UserEntity, error)
	FindManyUserByID(id []uint) ([]entity.UserEntity, error)
}
