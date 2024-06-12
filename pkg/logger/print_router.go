package logger

import (
	"fmt"
	"strings"

	"github.com/tsingshaner/go-pkg/color"
)

func formatPath(path string) string {
	return color.UnsafeGreen(fmt.Sprintf("%-40s", path))
}

func formatHandlerName(handlerName string) string {
	handlerNameTokens := strings.Split(handlerName, "/")[3:]
	return color.UnsafeCyan(strings.Join(handlerNameTokens, "/"))
}

var methodsTag = map[string]string{
	"GET":    color.UnsafeGreen("    GET"),
	"POST":   color.UnsafeYellow("   POST"),
	"PUT":    color.UnsafeBlue("    PUT"),
	"DELETE": color.UnsafeRed(" DELETE"),
	"PATCH":  color.UnsafeMagenta("  PATCH"),
}

func FormatMethod(method string) string {

	if formattedMethod, ok := methodsTag[method]; ok {
		return formattedMethod
	} else {
		return color.UnsafeCyan(fmt.Sprintf("%7s", method))
	}
}

func PrintRouter(method, absolutePath, handlerName string, nuHandlers int) {
	method = FormatMethod(method)
	absolutePath = formatPath(absolutePath)
	handlerName = formatHandlerName(handlerName)

	fmt.Printf("%s %s --{%d}-> %s\x1b[0m\n", method, absolutePath, nuHandlers, handlerName)
}
