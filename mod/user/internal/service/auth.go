package service

import (
	"errors"

	"github.com/tsingshaner/gin/mod/user/constant/errs"
	"github.com/tsingshaner/gin/mod/user/dto"
	"github.com/tsingshaner/gin/mod/user/interfaces"
	"github.com/tsingshaner/gin/mod/user/internal/entity"
	"github.com/tsingshaner/gin/shared"
	"github.com/tsingshaner/go-pkg/jwt"
)

type auth struct {
	tm *jwt.TokenMeta
	interfaces.UserQuery
	interfaces.UserModify
}

const Alg = jwt.EdDSA

func NewAuth(tm *jwt.TokenMeta, uq interfaces.UserQuery, um interfaces.UserModify) interfaces.Auth {
	return &auth{tm, uq, um}
}

func (a *auth) Login(username, password string) (*dto.Token, error) {
	user, err := a.UserQuery.ByUsername(username)
	if err != nil {
		if errors.Is(err, errs.Basic.DataNotFound) {
			return nil, errs.Unauthorized.Login
		}

		return nil, errors.Join(errs.InternalServerError.DatabaseQuery, err)
	}

	if ok, err := user.VerifyPassword(password); err != nil {
		return nil, errors.Join(errs.InternalServerError.PasswordHash, err)
	} else if !ok {
		return nil, errs.Unauthorized.Login
	}

	token, err := &dto.Token{}, error(nil)

	if token.Access, err = a.tm.SignedWithClaims(Alg,
		user.BuildAuthPayload(a.tm.NewRegisteredClaims(false)),
	); err != nil {
		return nil, errors.Join(errs.InternalServerError.GenToken, err)
	}

	if token.Refresh, err = a.tm.SignedWithClaims(Alg,
		user.BuildAuthPayload(a.tm.NewRegisteredClaims(true)),
	); err != nil {
		return nil, errors.Join(errs.InternalServerError.GenToken, err)
	}

	return token, nil
}

func (a *auth) Refresh(userID shared.ID) (*dto.Token, error) {
	return nil, errs.NotImplemented.None
}

func (a *auth) Register(username, password string) error {
	user := entity.New(entity.WithUser(&dto.User{Username: username, Password: password}))
	if err := user.HashPassword(); err != nil {
		return errors.Join(errs.InternalServerError.PasswordHash, err)
	}

	if _, err := a.UserModify.Create(user.User); err != nil {
		return err
	}

	return nil
}
