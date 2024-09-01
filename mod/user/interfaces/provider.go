package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/tsingshaner/gin/mod/user/dto"
	"github.com/tsingshaner/gin/mod/user/internal/entity"
	"github.com/tsingshaner/gin/shared"
)

type (
	UserQuery interface {
		ByID(shared.ID) (*entity.User, error)
		ByIDs([]shared.ID) ([]*entity.User, error)
		ByUsername(string) (*entity.User, error)
	}

	UserModify interface {
		Create(user *dto.User) (*entity.User, error)
		Update(user *dto.User) (*entity.User, error)
		UpdatePassword(id shared.ID, oldPassword, newPassword string) error
		Delete(shared.ID) error
	}

	Auth interface {
		Login(username, password string) (*dto.Token, error)
		Refresh(shared.ID) (*dto.Token, error)
		Register(username, password string) error
	}

	Verify interface {
		Validate(role dto.Role) gin.HandlerFunc
		Payload(*gin.Context) *dto.AuthPayload
	}
)
