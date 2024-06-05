package export

import "github.com/tsingshaner/gin-starter/internal/user/entity"

type UserExport interface {
	GetUser(id uint) (entity.UserEntity, error)
	GetManyUser(id []uint) ([]entity.UserEntity, error)
}
