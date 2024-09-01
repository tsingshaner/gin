package handler_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tsingshaner/gin/e2e/container"
	"github.com/tsingshaner/gin/e2e/h3test"
	"github.com/tsingshaner/gin/mod/user/constant/code"
)

func TestAuthLogin(t *testing.T) {
	t.Parallel()

	server, cleanup := container.NewAppWithCleanup()
	t.Cleanup(cleanup)

	t.Run("PasswordOrUsernameNotMatch", func(t *testing.T) {
		t.Parallel()

		res := api.authLogin.JSON(map[string]string{
			"username": "no_admin",
			"password": "123456",
		}).Send(server.Engine())

		assert.Equal(t, http.StatusUnauthorized, res.Code)
	})

	t.Run("Success", func(t *testing.T) {
		t.Parallel()

		assert.NoError(t, server.Providers().Auth.Register("admin", "123456"))

		res := api.authLogin.JSON(map[string]string{
			"username": "admin",
			"password": "123456",
		}).Send(server.Engine())

		assert.Equal(t, http.StatusOK, res.Code)

		data := h3test.ExtractJSON[struct {
			Code string `json:"code"`
			Data struct {
				RefreshToken string `json:"refreshToken"`
				AccessToken  string `json:"accessToken"`
			} `json:"data"`
		}](res)

		assert.Equal(t, code.Login, data.Code)
		assert.Greater(t, len(data.Data.RefreshToken), 10)
		assert.Greater(t, len(data.Data.AccessToken), 10)
	})
}
