package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary		用户信息
// @Description	登录
// @Tags			user
// @Accept			json
// @Produce		json
// @Param			q	query		string	false	"name search by q"	Format(email)
// @Success		200	{string}	string	"success"
// @Failure		400	{string}	string	"bad request"
// @Router			/user/profile [get]
func (*Handler) GetUserProfile(c *gin.Context) {
	time.Sleep(5 * time.Second)
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
