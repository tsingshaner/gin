package export

import "github.com/lab-online/internal/user/entity"

type UserExport interface {
	GetUser(string) entity.UserEntity
}
