package service

import (
	"github.com/gin-gonic/gin"
	"github.com/tsingshaner/gin/mod/user/constant/errs"
	"github.com/tsingshaner/gin/mod/user/dto"
	"github.com/tsingshaner/gin/mod/user/interfaces"
	"github.com/tsingshaner/gin/resp"
	"github.com/tsingshaner/gin/validator"
	"github.com/tsingshaner/go-pkg/jwt"
	"github.com/tsingshaner/go-pkg/util/bitmask"
)

const AuthKey = "@@auth"

type verify struct {
	tm *jwt.TokenMeta
}

func NewVerify(tm *jwt.TokenMeta) interfaces.Verify {
	return &verify{tm}
}

type authHeader struct {
	Authorization string `header:"Authorization" binding:"required,jwt,startswith=Bearer "`
}

func (v *verify) Validate(role dto.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		if p, ok := c.Get(AuthKey); ok {
			if payload, ok := p.(*dto.AuthPayload); ok {
				if err := v.judgeRolePassed(*payload.Role, role); err != nil {
					resp.Error(c, err)
				}

				return
			}
		}

		header := &authHeader{}
		if err := c.ShouldBindHeader(header); err != nil {
			validator.ErrorHandler(c, err)
			return
		}

		payload := &dto.AuthPayload{}
		if _, err := v.tm.ParseWithClaims(header.Authorization[7:], payload); err != nil {
			_, err = jwt.BuildRESTError(err)
			resp.Error(c, err)
		} else if err := v.judgeRolePassed(*payload.Role, role); err != nil {
			resp.Error(c, err)
		} else {
			c.Set(AuthKey, payload)
		}
	}
}

func (v *verify) judgeRolePassed(cur, target dto.Role) error {
	if target == 0 || bitmask.Has(target, cur) {
		return nil
	}

	return errs.Forbidden.RoleNotMatch
}

func (v *verify) Payload(c *gin.Context) *dto.AuthPayload {
	return c.MustGet(AuthKey).(*dto.AuthPayload)
}
