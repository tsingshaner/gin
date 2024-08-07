package shared

import "github.com/tsingshaner/gin/mod/user/dto"

const (
	RoleUser  dto.Role = 1 << iota
	RoleAdmin dto.Role = 1 << iota
	RoleSuper dto.Role = 1<<iota - 1
	RoleNone  dto.Role = 0
)
