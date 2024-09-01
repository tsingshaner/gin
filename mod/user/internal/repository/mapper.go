package repository

import (
	"github.com/tsingshaner/gin/mod/user/dto"
	"github.com/tsingshaner/gin/mod/user/internal/entity"
	"github.com/tsingshaner/gin/mod/user/model"
	"github.com/tsingshaner/gin/shared"
)

func convertUserModelToEntity(user *model.User) *entity.User {
	return entity.New(entity.WithPasswordHashed(), entity.WithUser(&dto.User{
		ID:       user.ID,
		Role:     user.Role,
		Username: user.Username,
		Nickname: user.Nickname,
		Password: user.Password,
	}))
}

func convertUserEntityToModel(user *entity.User) *model.User {
	return &model.User{
		Model:    *shared.NewModel(shared.WithModelID(user.ID)),
		Role:     user.Role,
		Username: user.Username,
		Nickname: user.Nickname,
		Password: user.Password,
	}
}
