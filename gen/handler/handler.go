package handler

import (
	_ "embed"
	"flag"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"text/template"

	"github.com/tsingshaner/gin/gen"
	"github.com/tsingshaner/go-pkg/log/console"
)

//go:embed handler.go.template
var tplStr string
var tpl *template.Template

func init() {
	var err error
	if tpl, err = template.New("HandlerFunc").Parse(tplStr); err != nil {
		console.Fatal("%+v", err)
	}
}

type apiMap = map[string]any

type handlerConf struct {
	Api apiMap `mapstructure:"api"`
}

type handler struct {
	Name   string
	Method string
}

type apiData struct {
	RouterName       string
	ParentRouterName string
	Path             string
	Guards           []string
	Handlers         []handler
	Child            []*apiData
}

type tplData struct {
	Chains []string
	Guards []string
	Api    []*apiData
}

var handlerChains []string
var handlerGuards []string

func Gen() {
	c := gen.Read[handlerConf]()
	tplInfo := &tplData{
		Chains: handlerChains,
		Guards: handlerGuards,
		Api:    parseApiData(c.Api, "", "r").Child,
	}

	output := flag.String("output", "handler_gen.go", "handler file output")
	flag.Parse()
	if pwd, err := os.Getwd(); err == nil {
		*output = filepath.Join(pwd, *output)
	}

	if writer, err := os.OpenFile(*output, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644); err == nil {
		if err := tpl.Execute(writer, tplInfo); err != nil {
			console.Fatal("%+v", err)
		}
	}

	console.Info("gen handler success")
}

func parseApiData(a apiMap, base, parentRouterName string) *apiData {
	data := &apiData{
		ParentRouterName: parentRouterName,
		RouterName:       parentRouterName + parseRouterName(base),
		Path:             base,
	}

	for k, v := range a {
		switch strings.ToLower(k) {
		case "guards":
			if guards, ok := v.([]any); ok {
				for _, g := range guards {
					if name, ok := g.(string); ok {
						guard := name + "Guard"
						data.Guards = append(data.Guards, guard)
						addGuard(guard)
					}
				}
			}
		// case "any":

		case "get", "post", "put", "delete", "patch", "options", "head":
			if name, ok := v.(string); ok {
				addChain(name + "Chain")
				data.Handlers = append(data.Handlers, handler{
					Name:   name + "Chain",
					Method: strings.ToUpper(k),
				})
			}
		default:
			if k[0] == '/' {
				if sub, ok := v.(apiMap); ok {
					data.Child = append(
						data.Child,
						parseApiData(sub, k, data.RouterName),
					)
				}
			}
		}
	}

	return data
}

func addChain(chain string) {
	if !slices.Contains(handlerChains, chain) {
		handlerChains = append(handlerChains, chain)
	}
}

func addGuard(guard string) {
	if !slices.Contains(handlerGuards, guard) {
		handlerGuards = append(handlerGuards, guard)
	}
}

func parseRouterName(path string) string {
	return strings.ReplaceAll(strings.ReplaceAll(path, ":", "_"), "/", "_")
}
