package main

import (
	"flag"
	"os"

	"github.com/tsingshaner/gin/gen/folder"
	"github.com/tsingshaner/gin/gen/handler"
	"github.com/tsingshaner/go-pkg/conf"
	"github.com/tsingshaner/go-pkg/log/console"
)

type Errors struct {
	Basic map[string]string
}

type ModConf struct {
	Version   string         `mapstructure:"version"`
	Namespace string         `mapstructure:"namespace"`
	Routes    map[string]any `mapstructure:"routes"`
	Errors
}

func main() {
	modCMD := flag.NewFlagSet("mod", flag.ExitOnError)
	console.Debug("%v", os.Args[1:])
	// modCMD.Usage()
	path := modCMD.String("config", "mod.yml", "Relative path of the config file.")
	handlerOnly := modCMD.Bool("hander-only", false, "Only generate handler files.")
	constantOnly := modCMD.Bool("constant-only", false, "Only generate constant pkg.")

	switch os.Args[1] {
	case "mod":
		if err := modCMD.Parse(os.Args[2:]); err != nil {
			modCMD.Usage()
			console.Fatal("%v", err)
		}

		console.Debug("%s %v %v", *path, *handlerOnly, *constantOnly)
		folder.Mkdir()

		config := conf.New(&ModConf{}, &conf.Options{
			FilePath: *path,
		})
		handler.GenerateHandler(&handler.GenerateHandlerOptions{
			Routes:     config.Value.Routes,
			OutputPath: "",
		})
	}

	// handlerCMD := flag.NewFlagSet("handler", flag.ExitOnError)
	// constantCMD := flag.NewFlagSet("constant", flag.ExitOnError)

	// if len(flag.Args()) < 1 {
	// 	flag.Usage()
	// 	fmt.Println("Please specify a command")
	// 	fmt.Println("Commands:")
	// 	fmt.Println("  mod")
	// 	fmt.Println("  handler")
	// 	fmt.Println("  constant")
	// 	return
	// }

	// switch flag.Args()[0] {
	// case "mod":
	// 	modCMD.Parse(flag.Args()[1:])
	// 	folder.Mkdir()
	// case "handler":
	// 	handlerCMD.Parse(flag.Args()[1:])
	// 	handler.Gen()
	// case "constant":
	// 	constantCMD.Parse(flag.Args()[1:])
	// 	constant.GenerateSuccessAndErrorConstants()
	// default:
	// 	flag.Usage()
	// }
}
