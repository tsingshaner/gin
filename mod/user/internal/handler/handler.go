//go:generate go run github.com/tsingshaner/gin/gen/cmd/gen-handler --config handler.yml
package handler

import "github.com/tsingshaner/gin/mod/user/interfaces"

type Provider struct {
	interfaces.Auth
	interfaces.Verify
	interfaces.UserModify
	interfaces.UserQuery
}
