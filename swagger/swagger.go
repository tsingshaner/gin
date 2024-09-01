package swagger

import (
	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwag "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag"
)

type Options struct {
	Enable   bool   `mapstructure:"enable"`
	DocsBase string `mapstructure:"docsBase"`

	Version     string `mapstructure:"version"`
	Title       string `mapstructure:"title"`
	Description string `mapstructure:"description"`
	Host        string `mapstructure:"-"`
	ApiBase     string `mapstructure:"-"`

	Server ginSwag.Config `mapstructure:"server"`
}

func New(conf ginSwag.Config) gin.HandlerFunc {
	return ginSwag.WrapHandler(files.Handler, func(c *ginSwag.Config) {
		c.Title = conf.Title
	})
}

func MergeDocsOptions(spec *swag.Spec, opts *Options) {
	spec.BasePath = opts.ApiBase
	spec.Host = opts.Host
	spec.Version = opts.Version

	if opts.Title != "" {
		spec.Title = opts.Title
	}

	if opts.Description != "" {
		spec.Description = opts.Description
	}
}
