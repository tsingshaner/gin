package repository

import (
	"log/slog"

	"github.com/lab-online/internal/user/entity"
	_ "github.com/lab-online/pkg/logger"
)

func (r *Repository) InsertUser(user *entity.User) error {
	return nil
}

func (r *Repository) CheckUserExists(userID string) (bool, error) {
	var count int64

	if err := r.db.Where("user_id = ?", userID).Count(&count).Error; err != nil {
		slog.Error("CheckUserExists", err)
		return false, err
	}

	return count > 0, nil
}
