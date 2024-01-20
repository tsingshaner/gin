package handler

import (
	"github.com/lab-online/internal/user/domain"
	"github.com/lab-online/internal/user/interface/http"
)

type Handler struct {
	domain domain.UserDomain
}

func NewHandler(domain domain.UserDomain) http.UserHandler {
	return &Handler{
		domain: domain,
	}
}
