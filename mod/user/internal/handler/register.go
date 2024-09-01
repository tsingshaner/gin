package handler

import (
	"github.com/tsingshaner/gin/mod/user/constant/code"
	"github.com/tsingshaner/gin/resp"
	v "github.com/tsingshaner/gin/validator"
)

type (
	registerBody struct {
		Username string `json:"username" binding:"required,min=5,max=20" example:"admin"`  // 用户名
		Password string `json:"password" binding:"required,min=6,max=20" example:"123456"` // 密码
	}
)

func (h *handler) registerChain() Chain {
	return Chain{v.Body[registerBody](), h.register}
}

// register
//
//	@Summary	用户注册
//	@Tags		用户模块
//	@Param		Body	body		registerBody	true	"JSON 请求体"
//	@Success	201		{object}	resp.SuccessBody[any, string]
//	@Failure	400		{object}	resp.FailedBody[string, string]
//	@Router		/users/register [post]
func (h *handler) register(c Ctx) {
	b := v.GetBody[registerBody](c)

	if err := h.Auth.Register(b.Username, b.Password); err != nil {
		resp.Error(c, err)
	} else {
		resp.Created[string, any](c, code.Register)
	}
}
