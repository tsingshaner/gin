package handler

import (
	"github.com/tsingshaner/gin/mod/user/constant/code"
	"github.com/tsingshaner/gin/mod/user/dto"
	"github.com/tsingshaner/gin/resp"
	"github.com/tsingshaner/gin/shared"
	v "github.com/tsingshaner/gin/validator"
)

type (
	getUserParams struct {
		ID shared.ID `uri:"id" binding:"required,min=1" example:"1"`
	}

	getUserResp struct {
		ID       uint64    `json:"id"`
		Role     *dto.Role `json:"role"`
		Username string    `json:"username"`
		Nickname string    `json:"nickname"`
	}
)

func (h *handler) getUserChain() Chain {
	return Chain{v.Params[getUserParams](), h.getUser}
}

// getUser
//
//	@Summary		查询用户
//	@Description	使用 id 查询用户信息
//	@Tags			用户模块
//	@Security		BearerToken
//	@Accept			json
//	@Produce		json
//	@Param			id	query		string							true	"用户ID"
//	@Success		200	{object}	getUserResp						"用户信息"
//	@Failure		401	{object}	resp.FailedBody[string, string]	"无权访问"
//	@Router			/users/{id} [get]
func (h *handler) getUser(c Ctx) {
	p := v.GetParams[getUserParams](c)

	if u, err := h.UserQuery.ByID(p.ID); err != nil {
		resp.Error(c, err)
	} else {
		resp.OK(c, code.GetUser, &getUserResp{
			ID:       u.ID,
			Role:     u.Role,
			Username: u.Username,
			Nickname: u.Nickname,
		})
	}
}
