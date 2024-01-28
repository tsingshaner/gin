package domain

import (
	"github.com/lab-online/internal/user/constant"
	"github.com/lab-online/internal/user/entity"
)

func (d *Domain) GetUser(userEntity entity.UserEntity) (entity.UserEntity, error) {
	user, err := d.repository.FindByUserID(userEntity.GetUserID())
	if err != nil {
		return user, newCustomError(constant.DB_ERROR)
	}

	return user, nil
}
