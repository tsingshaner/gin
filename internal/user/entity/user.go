package entity

import (
	"database/sql"
	"time"
)

type User struct {
	ID       int64        `json:"id"`
	UserID   string       `json:"userId"`
	Username string       `json:"username"`
	Password string       `json:"password"`
	Role     uint8        `json:"role"`
	Created  time.Time    `json:"created"`
	Updated  time.Time    `json:"updated"`
	Deleted  sql.NullTime `json:"deleted"`
}
