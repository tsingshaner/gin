package model

import (
	"fmt"
	"time"

	"github.com/tsingshaner/gin/mod/user/dto"
	"github.com/tsingshaner/gin/shared"
	"gorm.io/gorm"
)

type User struct {
	shared.Model
	Role     *dto.Role `json:"role" gorm:"default:0"`
	Username string    `json:"username" gorm:"type:varchar(100);uniqueIndex:uniq_username"`
	Nickname string    `json:"nickname" gorm:"type:varchar(64)"`
	Password string    `json:"password" gorm:"type:varchar(60);not null"`
}

func (*User) TableName() string {
	return "users"
}

// BeforeCreate 软删除前添加后缀, 解决再次创建时唯一索引冲突问题
func (u *User) BeforeDelete(tx *gorm.DB) error {
	return tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(u, u.ID).Error; err != nil {
			return err
		}

		uniqueUsername := fmt.Sprintf("%s&%d", u.Username, time.Now().UnixMilli())
		return tx.Model(u).Update("username", uniqueUsername).Error
	})
}
