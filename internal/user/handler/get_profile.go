package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tsingshaner/gin-starter/internal/user/constant"
	"github.com/tsingshaner/gin-starter/internal/user/entity"
	"github.com/tsingshaner/gin-starter/pkg/auth"
	"github.com/tsingshaner/gin-starter/pkg/resp"
)

type GetUserProfileRespData struct {
	UserID   string `json:"userID" binding:"required" example:"20240126"` // 学号或工号
	Username string `json:"username" binding:"required" example:"杏鸣"`     // 姓名
	Role     uint8  `json:"role" binding:"required" example:"1"`          // 角色
}
type GetUserProfileRespBody struct {
	resp.BaseRespBody
	Data GetUserProfileRespData `json:"data" binding:"required"`
}

// GetProfile
//
//	@Summary				用户信息
//	@Description.markdown	GetProfile
//	@Tags					user
//	@Accept					json
//	@Produce				json
//	@Param					Authorization	header		string	true	"Bearer Token"
//	@Success				200				{object}	GetUserProfileRespBody
//	@Failure				400				{object}	resp.BaseRespBody
//	@Failure				401				{object}	resp.BaseRespBody
//	@Security				bearer
//	@x-apifox-status		"testing"
//	@Router					/v1/user/profile [get]
func (h *Handler) GetProfile(c *gin.Context) {
	userInfo := c.MustGet(auth.KeyAuth).(*auth.AuthInfo)

	userEntity := entity.NewUserEntity(entity.WithUserID(userInfo.UserID))

	if user, err := h.domain.GetUser(userEntity); err != nil {
		handleError(c, err)
	} else {
		resp.Success(c, constant.GetProfileSuccess, &GetUserProfileRespData{
			UserID:   user.GetUserID(),
			Username: user.GetUsername(),
			Role:     user.GetRole(),
		})
	}
}
