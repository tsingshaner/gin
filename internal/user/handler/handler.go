package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/tsingshaner/gin-starter/pkg/errors"

	"github.com/tsingshaner/gin-starter/internal/user/constant"
	"github.com/tsingshaner/gin-starter/internal/user/domain"
	"github.com/tsingshaner/gin-starter/internal/user/interface/rest"
)

type Handler struct {
	domain domain.UserDomain
}

func NewHandler(domain domain.UserDomain) rest.UserHandler {
	return &Handler{
		domain: domain,
	}
}

func handleError(c *gin.Context, err error) {
	errors.HandleError(c, err, constant.StatusMap)
}
