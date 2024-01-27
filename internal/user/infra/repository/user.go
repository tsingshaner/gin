package repository

import (
	"database/sql"
	"log/slog"

	"gorm.io/gorm"

	"github.com/lab-online/internal/user/entity"
	"github.com/lab-online/internal/user/infra/model"
)

func (r *Repository) user() *gorm.DB {
	return r.db.Model(&model.User{})
}

func (r *Repository) InsertUser(userEntity entity.UserEntity) error {
	user := &model.User{
		UserID:   userEntity.GetUserID(),
		Username: userEntity.GetUsername(),
		Password: userEntity.GetPassword(),
		Role:     userEntity.GetRole(),
	}

	r.user().Save(user)
	return nil
}

func (r *Repository) CheckUserExists(userID string) (bool, error) {
	var count int64

	if err := r.user().Where("user_id = ?", userID).Count(&count).Error; err != nil {
		slog.Error("CheckUserExists", err)
		return false, err
	}

	return count > 0, nil
}

func (r *Repository) FindByUserID(userID string) (entity.UserEntity, error) {
	var user model.User

	if err := r.user().Where("user_id = ?", userID).First(&user).Error; err != nil {
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
