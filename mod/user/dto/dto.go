package dto

import (
	"github.com/tsingshaner/gin/shared"
	"github.com/tsingshaner/go-pkg/jwt"
)

type (
	Role uint8

	User struct {
		ID       shared.ID `json:"id"`
		Role     *Role     `json:"role"`
		Username string    `json:"username"`
		Nickname string    `json:"nickname"`
		Password string    `json:"password"`
	}

	AuthPayload struct {
		Role *Role     `json:"rol" example:"2"`
		UID  shared.ID `json:"uid" example:"1"`
		*jwt.RegisteredClaims
	}

	Token struct {
		Access  string `json:"accessToken" binding:"required"`
		Refresh string `json:"refreshToken,omitempty"`
	}
)

const (
	RoleUser  Role = 1 << iota
	RoleAdmin Role = 1 << iota
	RoleSuper Role = 1<<iota - 1
	RoleNone  Role = 0
)

var _ jwt.Claims = &AuthPayload{}
