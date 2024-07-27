package gen

import (
	"github.com/tsingshaner/go-pkg/conf"
	"github.com/tsingshaner/go-pkg/log/console"
)

func Read[T any]() *T {
	c := conf.New(new(T), conf.ParseArgs())
	if err := c.Load(); err != nil {
		console.Fatal("load config error: %v", err)
	}

	return c.Value
}
