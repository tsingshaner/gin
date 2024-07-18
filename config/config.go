package config

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/spf13/viper"
	"github.com/tsingshaner/go-pkg/color"
	"github.com/tsingshaner/go-pkg/log/console"
)

func init() {
	args := parseArgs()
	loadConfig(args)
	s = NewStore()

	if !args.silence {
		showConfig()
	}
}

type args struct {
	filePath string
	fileType string
	silence  bool
}

func loadConfig(args *args) {
	viper.SetConfigFile(args.filePath + args.fileType)

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func parseArgs() *args {
	args := &args{}

	flag.StringVar(&args.filePath, "config", "config.json", "Path to the configuration file")
	flag.BoolVar(&args.silence, "silence", false, "Silence the output of config loading")
	flag.Parse()

	args.fileType = filepath.Ext(args.filePath)

	if dir, err := os.Getwd(); err == nil {
		args.filePath = filepath.Join(dir, args.filePath)
		args.filePath = args.filePath[0 : len(args.filePath)-len(args.fileType)]
	}

	if !args.silence {
		showArgs(args)
	}

	return args
}

func showArgs(args *args) {
	console.Info(
		"will load configuration from %s",
		color.UnsafeCyan(args.filePath)+color.UnsafeYellow(args.fileType),
	)
}

func showConfig() {
	sb := &strings.Builder{}
	sb.WriteString("config loaded success")
	formatStruct(sb, "    * ", s)

	console.Trace(sb.String())
}

func formatStruct(sb *strings.Builder, prefix string, obj any) {
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	} else if v.Kind() != reflect.Struct {
		sb.WriteString(fmt.Sprintf("%v", obj))
		return
	}

	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		sb.WriteString(fmt.Sprintf(
			"\n%s%s %s: ",
			prefix,
			color.UnsafeCyan(t.Field(i).Name),
			color.UnsafeYellow(t.Field(i).Type.String())),
		)

		formatStruct(sb, "  "+prefix, field.Interface())
	}
}
