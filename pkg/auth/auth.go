package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tsingshaner/gin-starter/pkg/errors"
	"github.com/tsingshaner/gin-starter/pkg/jwt"
	"github.com/tsingshaner/gin-starter/pkg/resp"
)

type Role uint8

const (
	RoleNone Role = iota
	RoleStudent
	RoleTeacher
	RoleAdmin
)
const KeyAuth = "auth"

type AuthInfo struct {
	UserID string
	Role   Role
}

type TokenParser interface {
	ParseToken(token string) (*jwt.Claims, error)
}

type Middleware func(enabledRoles []Role) gin.HandlerFunc

func New(parser TokenParser) Middleware {
	return func(enabledRoles []Role) gin.HandlerFunc {
		isPublic := isEnabledRole(RoleNone, enabledRoles)

		return func(c *gin.Context) {
			authorization := c.GetHeader("Authorization")
			if len(authorization) <= 7 {
				err := errors.New(resp.CodeAuthError, "Authorization header is invalid")
				c.Error(err)
				resp.BadRequest(c, err.Code, err.Message)
				return
			}

			if claims, err := parser.ParseToken(authorization[7:]); err != nil {
				c.Error(errors.New(resp.CodeAuthError, "Token is invalid"))
				c.AbortWithStatusJSON(
					http.StatusUnauthorized,
					resp.NewBaseRespBody(resp.CodeAuthError),
				)
				return
			} else {
				if isPublic || isEnabledRole(Role(claims.Role), enabledRoles) {
					c.Set(KeyAuth, &AuthInfo{claims.UserID, Role(claims.Role)})
					c.Next()
				} else {
					c.Error(errors.New(resp.CodeForbidden, "Permission denied"))
					c.AbortWithStatusJSON(
						http.StatusForbidden,
						resp.NewBaseRespBody(resp.CodeForbidden),
					)
					return
				}
			}
		}
	}
}

func isEnabledRole(role Role, enabledRoles []Role) bool {
	for _, r := range enabledRoles {
		if r == role {
			return true
		}
	}

	return false
}
