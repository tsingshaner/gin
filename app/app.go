//	@description.markdown
//	@termsOfService	https://swagger.io/terms/

//	@securityDefinitions.apikey	bearer
//	@in							header
//	@name						Authorization

//	@contact.name	Issues
//	@contact.url	https://github.com/tsingshaner/gin/issues

//	@license.name	ISC
//	@license.url	https://github.com/tsingshaner/gin/blob/main/LICENSE

// @externalDocs.description	ApiFox
// @externalDocs.url			https://apifox.com/apidoc/shared-3e844af7-e01f-4a3a-a44d-9b395189d4d5
package app

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tsingshaner/go-pkg/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(opts *Options) *app {
	zapCore, _ := log.NewZapCore(
		log.NewZapJSONEncoder(),
		zapcore.AddSync(os.Stdout),
		int8(opts.Logger.Level),
	)

	a := &app{
		Options: opts,
		logger: log.NewZapLog(zapCore, zap.AddCaller(), zap.AddCallerSkip(3)).
			Named("gin-template").Child(slog.Int("pid", os.Getpid())),
	}

	return a
}

type app struct {
	Options *Options

	effects []func()
	engine  *gin.Engine
	logger  log.Slog
	server  *http.Server
}
