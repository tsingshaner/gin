package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID       int64  `gorm:"index"`
	UserID   string `gorm:"unique"`
	Username string
	Password string
	Role     int8
	Created  time.Time
	Updated  time.Time
	Deleted  sql.NullTime
}
