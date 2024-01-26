package repository

import (
	"log/slog"

	"github.com/lab-online/internal/user/entity"
	"github.com/lab-online/pkg/model"
)

func (r *Repository) InsertUser(userEntity entity.UserEntity) error {
	user := &model.User{
		UserID:   userEntity.GetUserID(),
		Username: userEntity.GetUsername(),
		Password: userEntity.GetPassword(),
		Role:     userEntity.GetRole(),
	}

	r.user.Save(user)
	return nil
}

func (r *Repository) CheckUserExists(userID string) (bool, error) {
	var count int64

	if err := r.user.Where("user_id = ?", userID).Count(&count).Error; err != nil {
		slog.Error("CheckUserExists", err)
		return false, err
	}

	return count > 0, nil
}
