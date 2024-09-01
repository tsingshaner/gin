package handler

import "github.com/tsingshaner/gin/mod/user/dto"

func (h *handler) authGuard() HF {
	return h.Validate(dto.RoleNone)
}
