package domain

import (
	"github.com/tsingshaner/gin-starter/internal/user/constant"
	"github.com/tsingshaner/gin-starter/internal/user/entity"
)

func (d *Domain) GetUser(userEntity entity.UserEntity) (entity.UserEntity, error) {
	user, err := d.repository.FindByUserID(userEntity.GetUserID())
	if err != nil {
		return user, newCustomError(constant.DBError)
	}

	return user, nil
}
