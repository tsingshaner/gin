package export

import "github.com/lab-online/internal/user/entity"

type UserExport interface {
	GetUser(id uint) (entity.UserEntity, error)
	GetManyUser(id []uint) ([]entity.UserEntity, error)
}
