package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/tsingshaner/gin-starter/pkg/middleware"
	"github.com/tsingshaner/gin-starter/pkg/resp"

	"github.com/tsingshaner/gin-starter/internal/user/constant"
	"github.com/tsingshaner/gin-starter/internal/user/entity"
)

type RegisterReqBody struct {
	UserID   string `json:"userID" binding:"required" example:"20240126"`                // 学号或者工号
	Username string `json:"username" binding:"required" example:"杏鸣"`                    // 姓名
	Password string `json:"password" binding:"required,max=64,min=8" example:"12345678"` // 密码
}

var RegisterValidator = middleware.Validator(&middleware.ValidatorOptions{
	Body: &RegisterReqBody{},
})

type RegisterFailedRespBody struct {
	resp.BaseRespBody
	Err string `json:"err"  binding:"required" example:"错误信息"`
}

// Register
//
//	@Summary				注册
//	@Description.markdown	Register
//	@Tags					user
//	@Accept					json
//	@Produce				json
//	@Param					Body	body		RegisterReqBody	true	"JSON 请求体"
//	@Success				201		{object}	resp.BaseRespBody
//	@Failure				400		{object}	RegisterFailedRespBody
//	@Router					/v1/user/ [post]
func (h *Handler) Register(c *gin.Context) {
	user := c.MustGet(middleware.KeyBody).(*RegisterReqBody)

	userEntity := entity.NewUserEntity(
		entity.WithPassword(user.Password),
		entity.WithUserID(user.UserID),
		entity.WithUsername(user.Username),
		entity.WithRole(1),
	)

	if err := h.domain.AddUser(userEntity); err != nil {
		handleError(c, err)
		return
	}

	resp.Created[any](c, constant.SuccessRegister)
}
