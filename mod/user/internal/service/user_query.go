package service

import (
	"github.com/tsingshaner/gin/mod/user/constant/errs"
	"github.com/tsingshaner/gin/mod/user/interfaces"
	"github.com/tsingshaner/gin/mod/user/internal/entity"
	"github.com/tsingshaner/gin/mod/user/internal/repository"
	"github.com/tsingshaner/gin/shared"
)

type userQuery struct {
	repo *repository.User
}

func NewUserQuery(repo *repository.User) interfaces.UserQuery {
	return &userQuery{repo}
}

func (uq *userQuery) ByID(id shared.ID) (*entity.User, error) {
	return nil, errs.NotImplemented.None

}
func (uq *userQuery) ByIDs(ids []shared.ID) ([]*entity.User, error) {
	return nil, errs.NotImplemented.None
}

func (uq *userQuery) ByUsername(username string) (*entity.User, error) {
	return uq.repo.GetUserByUsername(username)
}
