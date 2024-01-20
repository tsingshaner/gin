package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary		用户登录
// @Description	登录
// @Tags			user
// @Accept			json
// @Produce		json
// @Param			q	query		string	false	"name search by q"	Format(email)
// @Success		200	{string}	string	"success"
// @Failure		400	{string}	string	"bad request"
// @Router			/user/login [post]
func (*Handler) Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
