package user

import (
	"github.com/tsingshaner/gin-starter/internal/user/entity"
	"github.com/tsingshaner/gin-starter/internal/user/infra/model"
)

func (u *User) InsertUser(userEntity entity.UserEntity) error {
	user := &model.User{
		UserID:   userEntity.GetUserID(),
		Username: userEntity.GetUsername(),
		Password: userEntity.GetPassword(),
		Role:     userEntity.GetRole(),
	}

	u.table().Save(user)
	return nil
}
