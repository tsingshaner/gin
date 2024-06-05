package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/tsingshaner/gin-starter/pkg/middleware"
	"github.com/tsingshaner/gin-starter/pkg/resp"

	"github.com/tsingshaner/gin-starter/internal/user/constant"
	"github.com/tsingshaner/gin-starter/internal/user/entity"
)

type LoginReqBody struct {
	UserID   string `json:"userID" binding:"required,max=10,min=3" example:"20240126"`   // 学号或者工号
	Password string `json:"password" binding:"required,max=64,min=8" example:"12345678"` // 密码
}

var LoginValidator = middleware.Validator(&middleware.ValidatorOptions{
	Body: &LoginReqBody{},
})

type AccessToken struct {
	AccessToken string `json:"accessToken" binding:"required" example:"token"` // 访问令牌
}
type LoginRespBody struct {
	resp.BaseRespBody
	Data *AccessToken `json:"data" binding:"required"`
}
type LoginFailedRespBody struct {
	resp.BaseRespBody
	Err string `json:"err"  binding:"required" example:"错误信息"`
}

// Login
//
//	@Summary				用户登录
//	@Description.markdown	Login
//	@Tags					user
//	@Accept					json
//	@Produce				json
//	@Param					Body	body		LoginReqBody	true	"JSON 请求体"
//	@Success				200		{object}	LoginRespBody
//	@Failure				400		{object}	LoginFailedRespBody
//	@Router					/v1/user/login [post]
func (h *Handler) Login(c *gin.Context) {
	user := c.MustGet(middleware.KeyBody).(*LoginReqBody)

	// TODO 限制登录次数

	userEntity := entity.NewUserEntity(
		entity.WithUserID(user.UserID),
		entity.WithPassword(user.Password),
	)

	if token, err := h.domain.Login(userEntity); err != nil {
		handleError(c, err)
		return
	} else {
		resp.Success(c, constant.SuccessLogin, &AccessToken{token})
	}
}
