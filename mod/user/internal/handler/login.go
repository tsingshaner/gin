package handler

import (
	"github.com/tsingshaner/gin/mod/user/constant/code"
	"github.com/tsingshaner/gin/mod/user/constant/errs"
	"github.com/tsingshaner/gin/resp"
	v "github.com/tsingshaner/gin/validator"
)

type (
	loginQuery struct {
		Grant string `form:"grant" binding:"required,min=1,max=10,oneof=pwd oauth2" example:"pwd"` // 验证策略
	}

	loginBody struct {
		Username string `json:"username" binding:"required,min=5,max=20" example:"admin"`  // 用户名
		Password string `json:"password" binding:"required,min=6,max=20" example:"123456"` // 用户密码
	}

	loginResp struct {
		Access  string `json:"accessToken" binding:"required" example:"xxxxx"` // 身份令牌
		Refresh string `json:"refreshToken" binding:"required" example:"xxxx"` // 刷新令牌
	}
)

func (h *handler) loginChain() Chain {
	return Chain{v.Query[loginQuery](), v.Body[loginBody](), h.login}
}

// login
//
//	@Summary			用户登录
//	@Tags				用户模块
//	@x-apifox-status	"designing"
//	@Accept				json
//	@Produce			json
//	@Param				grant	query		loginQuery	true	"验证策略"
//	@Param				JSON	body		loginBody	true	"JSON payload"
//	@Success			200		{object}	loginResp
//	@Failure			401		{object}	resp.FailedBody[string, string]
//	@router				/users/login [post]
func (h *handler) login(c Ctx) {
	q, b := v.GetQuery[loginQuery](c), v.GetBody[loginBody](c)

	if q.Grant == "oauth2" {
		resp.Error(c, errs.NotImplemented.None)
	} else if token, err := h.Auth.Login(b.Username, b.Password); err != nil {
		resp.Error(c, err)
	} else {
		resp.OK(c, code.Login, &loginResp{
			Access:  token.Access,
			Refresh: token.Refresh,
		})
	}
}
