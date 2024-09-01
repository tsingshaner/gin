package app

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/zh"
	zhTranslator "github.com/go-playground/validator/v10/translations/zh"
	"github.com/tsingshaner/gin/middleware"
	"github.com/tsingshaner/gin/middleware/requestid"
	"github.com/tsingshaner/gin/validator"
	"github.com/tsingshaner/go-pkg/jwt"
	"github.com/tsingshaner/go-pkg/log"
	"github.com/tsingshaner/go-pkg/log/console"
	"github.com/tsingshaner/go-pkg/log/helper"
)

// Ready init app for starting
func (a *app) Ready() {
	if a.engine != nil {
		return
	}

	a.buildEngine()
	a.buildServer()

	if j, err := jwt.New(a.Options.JWT); err != nil {
		console.Fatal("jwt init failed %s", err)
	} else {
		a.jwtMeta = j
	}

	a.connectDatabase()

	console.Info("app is ready to start")
}

// buildEngine init gin engine with mode
func (a *app) buildEngine() {
	gin.SetMode(a.Options.Server.Mode)

	gin.DebugPrintRouteFunc = helper.NewGinRouterLogger(a.logger)

	ginDebugLogger := a.logger.WithOptions(&log.ChildLoggerOptions{
		AddSource:  true,
		SkipCaller: 2,
	})
	gin.DebugPrintFunc = func(format string, values ...any) {
		if ginDebugLogger.Enabled(slog.Level(log.LevelDebug)) {
			ginDebugLogger.Debug(strings.ReplaceAll(
				strings.TrimRight(fmt.Sprintf(format, values...), "\n"),
				"\n", "\n      ",
			))
		}
	}

	_ = validator.ApplyTranslator(zh.New(), true, zhTranslator.RegisterDefaultTranslations)

	a.engine = gin.New()
	a.engine.NoRoute(middleware.NotFoundHandler)
	a.engine.Use(
		requestid.New(&requestid.Options{HeaderKey: a.Options.Server.RequestIdHeader}),
		helper.New(&helper.Options{Logger: a.logger, TraceIDExtractor: requestid.Get}),
		middleware.NewErrorHandler(a.logger),
		middleware.Cors(a.Options.Cors),
	)
}

func (a *app) buildServer() {
	a.server = BuildServer(a.engine, a.Options.Server)
}

// BuildServer init http server with gin engine
func BuildServer(engine *gin.Engine, opts *ServerOptions) *http.Server {
	return &http.Server{
		Addr:           fmt.Sprintf("%s:%d", opts.Host, opts.Port),
		Handler:        engine,
		ReadTimeout:    opts.ReadTimeout,
		WriteTimeout:   opts.WriteTimeout,
		MaxHeaderBytes: opts.MaxHeaderBytes,
	}
}
