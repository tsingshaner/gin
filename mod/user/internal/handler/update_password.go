package handler

import (
	"github.com/tsingshaner/gin/mod/user/constant/code"
	"github.com/tsingshaner/gin/mod/user/constant/errs"
	"github.com/tsingshaner/gin/resp"
	"github.com/tsingshaner/gin/shared"
	v "github.com/tsingshaner/gin/validator"
)

type (
	updatePasswordParams struct {
		ID shared.ID `uri:"id" binding:"required,min=1" example:"1"` // 用户 ID
	}

	updatePasswordBody struct {
		OldPassword string `json:"oldPassword" binding:"required,min=6,max=20" example:"123456"` // 旧密码
		NewPassword string `json:"newPassword" binding:"required,min=6,max=20" example:"123456"` // 新密码
	}
)

func (h *handler) updatePasswordChain() Chain {
	return Chain{v.Params[updatePasswordParams](), v.Body[updatePasswordBody](), h.updatePassword}
}

// updatePassword
//
//	@Summary		修改密码
//	@Description	仅可更新当前用户的密码
//	@Tags			用户模块
//	@Security		BearerToken
//	@Accept			json
//	@Produce		json
//	@Param			id	path		integer							true	"用户ID"
//	@Success		200	{object}	resp.SuccessBody[any, string]	"修改成功"
//	@Failure		401	{object}	resp.FailedBody[string, string]	"无权限"
//	@Router			/users/{id}/password [patch]
func (h *handler) updatePassword(c Ctx) {
	auth := h.Verify.Payload(c)
	p, b := v.GetParams[updatePasswordParams](c), v.GetBody[updatePasswordBody](c)

	if auth.UID != p.ID {
		resp.Error(c, errs.Forbidden.RoleNotMatch)
		return
	}

	if err := h.UserModify.UpdatePassword(p.ID, b.OldPassword, b.NewPassword); err != nil {
		resp.Error(c, err)
	} else {
		resp.OK[string, any](c, code.UpdatePassword)
	}
}
