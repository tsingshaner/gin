package user

import (
	"database/sql"
	"log/slog"

	"github.com/lab-online/internal/user/entity"
	"github.com/lab-online/internal/user/infra/model"
)

func (u *User) CheckUserExists(userID string) (bool, error) {
	var count int64

	if err := u.table().Where("user_id = ?", userID).Count(&count).Error; err != nil {
		slog.Error("CheckUserExists", err)
		return false, err
	}

	return count > 0, nil
}

func (u *User) FindByUserID(userID string) (entity.UserEntity, error) {
	var user model.User

	if err := u.table().Where("user_id = ?", userID).First(&user).Error; err != nil {
		slog.Error("FindByUserID", err)
		return nil, err
	}

	return &entity.User{
		ID:        user.ID,
		UserID:    user.UserID,
		Username:  user.Username,
		Password:  user.Password,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: sql.NullTime(user.DeletedAt),
	}, nil
}
