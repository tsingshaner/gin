package handler_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tsingshaner/gin/e2e/container"
	"github.com/tsingshaner/gin/e2e/h3test"
	"github.com/tsingshaner/gin/mod/user/constant/code"
)

func TestAuthRegister(t *testing.T) {
	t.Parallel()

	server, cleanup := container.NewAppWithCleanup()
	t.Cleanup(cleanup)

	t.Run("Validator", func(t *testing.T) {
		t.Parallel()
		cases := []struct {
			rawBody map[string]string
			msg     string
		}{
			{map[string]string{}, "without body"},
			{map[string]string{"password": "123456"}, "without username field"},
			{map[string]string{"username": "user", "password": "123456"}, "username is too short"},
		}

		for _, c := range cases {
			t.Run(c.msg, func(t *testing.T) {
				t.Parallel()
				res := api.authRegister.JSON(c.rawBody).Send(server.Engine())
				assert.Equal(t, http.StatusBadRequest, res.Code, c.msg)
			})
		}
	})

	t.Run("Success", func(t *testing.T) {
		t.Parallel()
		res := api.authRegister.JSON(map[string]string{
			"username": "admin",
			"password": "123456",
		}).Send(server.Engine())

		assert.Equal(t, http.StatusCreated, res.Code)

		data := h3test.ExtractJSON[struct {
			Code string `json:"code"`
		}](res)

		assert.Equal(t, code.Register, data.Code)
	})
}
