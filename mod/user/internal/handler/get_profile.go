package handler

import (
	"github.com/tsingshaner/gin/mod/user/constant/code"
	"github.com/tsingshaner/gin/mod/user/dto"
	"github.com/tsingshaner/gin/resp"
)

type (
	getProfileResp struct {
		Username string    `json:"username" example:"admin"` // 用户名
		Nickname string    `json:"nickname" example:"name"`  // 昵称
		Role     *dto.Role `json:"role" example:"3"`         // 角色
	}
)

func (h *handler) getProfileChain() Chain {
	return Chain{h.getProfile}
}

// getProfile
//
//	@x-apifox-status	"released"
//	@Summary			用户信息
//	@Description		获取当前 token 的用户信息
//	@Tags				用户信息
//	@Security			BearerToken
//	@Accept				json
//	@Produce			json
//	@Param				Authorization	header		string							true	"Access Token"
//	@Success			200				{object}	getProfileResp					"用户信息"
//	@Failure			400				{object}	resp.FailedBody[string, string]	"token 无效"
//	@Router				/users [get]
func (h *handler) getProfile(c Ctx) {
	auth := h.Verify.Payload(c)
	if user, err := h.UserQuery.ByID(auth.UID); err != nil {
		resp.Error(c, err)
	} else {
		resp.OK(c, code.GetProfile, &getProfileResp{
			Username: user.Username,
			Nickname: user.Nickname,
			Role:     user.Role,
		})
	}
}
