package config

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/tsingshaner/gin/app"
	"github.com/tsingshaner/go-pkg/conf"
	"github.com/tsingshaner/go-pkg/log"
	"github.com/tsingshaner/go-pkg/log/console"
)

type (
	store struct {
		app.Options  `mapstructure:",squash"`
		ConsoleLevel log.Level `mapstructure:"consoleLevel"`
	}
	Options = conf.Options
)

var (
	s          *conf.Config[store]
	opts       *conf.Options
	currentDir string
)

func Store() *store {
	return s.Value
}

func Init() {
	opts = conf.ParseArgs()
	s = conf.New(&store{}, opts)
	if err := s.Load(); err != nil {
		console.Fatal("load config error: %v", err)
	}
	console.SetLevel(Store().ConsoleLevel)
}

// ************************** for test **************************

func init() {
	_, file, _, _ := runtime.Caller(0)
	currentDir = filepath.Dir(file)
}

// GetPath get the absolute path of the file base on the config dir, use for test
func getPath(path string) string {
	return filepath.Join(currentDir, path)
}

// NewTestConf this func is designed for test purpose, it will fix the path of the config file
func NewTestConf(silence ...bool) *app.Options {
	option := &Options{
		FilePath: getPath("config.test.yml"),
		Silence:  true,
	}

	if len(silence) > 0 && !silence[0] {
		option.Silence = false
	}

	config := conf.New(&store{}, option)
	if err := config.Load(); err != nil {
		console.Fatal("load config error: %v", err)
	}

	for i, m := range config.Value.JWT.Methods {
		if m.Pem.PrivatePath != "" {
			config.Value.JWT.Methods[i].Pem.PrivatePath = getPath(m.Pem.PrivatePath)
			config.Value.JWT.Methods[i].Pem.PublicPath = getPath(m.Pem.PublicPath)
		}
	}

	if os.Getenv("CI") == "true" {
		l := log.LevelError | log.LevelFatal
		config.Value.Options.Logger.Level = l
		console.SetLevel(l)
	}

	return &config.Value.Options
}
