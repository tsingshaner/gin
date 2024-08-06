package container

import (
	"context"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"github.com/tsingshaner/gin/app"
	"github.com/tsingshaner/go-pkg/log/console"
)

type logger struct{}

func (logger) Printf(format string, v ...interface{}) {
	console.Trace(format, v...)
}

func Postgres(ctx context.Context, wg *sync.WaitGroup,
	opts *app.Postgres) *postgres.PostgresContainer {
	wg.Add(1)

	postgresContainer, err := postgres.Run(ctx, "docker.io/postgres:16-alpine",
		postgres.WithDatabase(opts.Database),
		postgres.WithUsername(opts.User),
		postgres.WithPassword(opts.Password),
		testcontainers.WithLogger(logger{}),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		console.Fatal("failed to start container: %s", err)
	}

	if ep, err := postgresContainer.Endpoint(ctx, ""); err != nil {
		console.Fatal("failed to get container endpoint: %s", err)
	} else {
		endpoint := strings.Split(ep, ":")
		port, _ := strconv.Atoi(endpoint[1])

		opts.Host = endpoint[0]
		opts.Port = int64(port)
	}

	go func() {
		defer wg.Done()
		<-ctx.Done()
		if err := postgresContainer.Terminate(context.TODO()); err != nil {
			console.Fatal("failed to terminate container: %s", err)
		}
	}()

	return postgresContainer
}
