package handler_test

import "github.com/tsingshaner/gin/e2e/h3test"

var api = struct {
	authRegister, authLogin *h3test.Request
}{
	authRegister: h3test.New("/api/auth/register").POST(),
	authLogin:    h3test.New("/api/auth/login").AddQuery("grant", "pwd").POST(),
}
