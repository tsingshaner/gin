package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lab-online/internal/user/entity"
	"github.com/lab-online/pkg/middleware"
	"github.com/lab-online/pkg/resp"
)

type RegisterReqBody struct {
	UserID   string `json:"userID" binding:"required" example:"20240126"`                // 学号或者工号
	Username string `json:"username" binding:"required" example:"杏鸣"`                    // 姓名
	Password string `json:"password" binding:"required,max=64,min=8" example:"12345678"` // 密码
}

type RegisterRespBody = resp.BaseRespBody

var RegisterReq = &middleware.ValidatorOptions{
	Body: &RegisterReqBody{},
}

// Register
//
//	@Summary	注册
//	@Description.markdown
//	@Tags		user
//	@Accept		json
//	@Produce	json
//	@Param		Body	body		RegisterReqBody	true	"JSON 请求体"
//	@Success	201		{object}	RegisterRespBody
//	@Failure	400		{object}	string
//	@Router		/v1/user/ [post]
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

	resp.Created[any](c, 20100)
}
