package container

import (
	"context"
	"sync"

	"github.com/tsingshaner/gin/app"
	"github.com/tsingshaner/gin/config"
)

func App(ctx context.Context, wg *sync.WaitGroup, opts *app.Options) app.Server {
	Postgres(ctx, wg, opts.Postgres)

	server := app.New(opts)
	server.Ready()
	server.Migrate()
	server.BindRoutes()

	return server
}

func NewAppWithCleanup() (server app.Server, cleanup func()) {
	wg := &sync.WaitGroup{}
	ctx, done := context.WithCancel(context.Background())

	return App(ctx, wg, config.NewTestConf()), func() {
		done()
		wg.Wait()
	}
}
