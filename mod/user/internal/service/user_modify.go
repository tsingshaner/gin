package service

import (
	"errors"

	"github.com/tsingshaner/gin/mod/user/constant/errs"
	"github.com/tsingshaner/gin/mod/user/dto"
	"github.com/tsingshaner/gin/mod/user/interfaces"
	"github.com/tsingshaner/gin/mod/user/internal/entity"
	"github.com/tsingshaner/gin/mod/user/internal/repository"
	"github.com/tsingshaner/gin/shared"
)

type userModify struct {
	interfaces.UserQuery
	repo *repository.User
}

func NewUserModify(
	repo *repository.User, userQuery interfaces.UserQuery,
) interfaces.UserModify {
	return &userModify{userQuery, repo}
}

func (um *userModify) Create(user *dto.User) (*entity.User, error) {
	if u, err := um.repo.Save(&entity.User{User: user}); err != nil {
		if errors.Is(err, errs.Basic.Duplicate) {
			return nil, errors.Join(errs.Basic.Duplicate, err)
		}
		return nil, errors.Join(errs.InternalServerError.DatabaseInsert, err)
	} else {
		return u, nil
	}
}
func (um *userModify) Update(user *dto.User) (*entity.User, error) {
	return nil, errs.NotImplemented.None
}
func (um *userModify) Delete(id shared.ID) error {
	return errs.NotImplemented.None
}

func (um *userModify) UpdatePassword(id shared.ID, oldPassword, newPassword string) error {
	u, err := um.UserQuery.ByID(id)
	if err != nil {
		return err
	}

	if ok, err := u.VerifyPassword(oldPassword); err != nil {
		return errors.Join(errs.InternalServerError.PasswordHash, err)
	} else if !ok {
		return errs.Forbidden.RoleNotMatch
	}

	u.Password = newPassword
	if err := u.HashPassword(); err != nil {
		return errors.Join(errs.InternalServerError.PasswordHash, err)
	}

	if _, err := um.repo.Save(u); err != nil {
		return errors.Join(errs.InternalServerError.DatabaseInsert, err)
	}
	return nil
}
