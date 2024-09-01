package app

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"time"

	"github.com/tsingshaner/go-pkg/log/console"
)

func (a *app) Shutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	console.Info("shutdown server ......")
	defer a.logger.Info("server has been shut down")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := a.server.Shutdown(ctx); err != nil {
		a.logger.Error("error occurred while the server was shutting down",
			slog.String("err", err.Error()))
	}

	console.Info("cancel context ......")
	for _, fn := range a.effects {
		fn()
	}
}
