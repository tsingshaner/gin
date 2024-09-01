package repository

import (
	"errors"

	"github.com/tsingshaner/gin/mod/user/constant/errs"
	"github.com/tsingshaner/gin/mod/user/internal/entity"
	"github.com/tsingshaner/gin/mod/user/internal/repository/query"
	"github.com/tsingshaner/gin/mod/user/model"
	"github.com/tsingshaner/gin/shared"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	query.SetDefault(db)
	return &User{db}
}

func (u *User) GetUserByID(id shared.ID) (*entity.User, error) {
	userModel, err := query.User.Where(query.User.ID.Eq(id)).First()

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Join(errs.Basic.DataNotFound, err)
		}

		return nil, errors.Join(errs.InternalServerError.DatabaseQuery, err)
	}

	return convertUserModelToEntity(userModel), nil
}

func (u *User) GetUserByUsername(username string) (*entity.User, error) {
	userModel, err := query.User.Where(query.User.Username.Eq(username)).First()

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Join(errs.Basic.DataNotFound, err)
		}

		return nil, errors.Join(errs.Basic.DataQuery, err)
	}

	return convertUserModelToEntity(userModel), nil
}

func (u *User) Save(user *entity.User) (*entity.User, error) {
	userModel := convertUserEntityToModel(user)

	if userModel.ID == 0 {
		if err := query.User.Create(userModel); err != nil {
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				return nil, errors.Join(errs.Basic.Duplicate, err)
			}

			return nil, errors.Join(errs.Basic.DataInsert, err)
		}

		user.ID = userModel.ID
		return user, nil
	}

	if _, err := query.User.Where(query.User.ID.Eq(userModel.ID)).Updates(userModel); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, errors.Join(errs.Basic.Duplicate, err)
		}

		return nil, errors.Join(errs.Basic.DataInsert, err)
	}

	return user, nil
}

func (u *User) Remove(id shared.ID) error {
	if _, err := query.User.Delete(&model.User{
		Model: *shared.NewModel(shared.WithModelID(id)),
	}); err != nil {
		return errors.Join(errs.Basic.Delete, err)
	}

	return nil
}
