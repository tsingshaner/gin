package logger

import (
	"fmt"
	"strings"
)

func formatPath(path string) string {
	return Style(fmt.Sprintf("%-40s", path), ColorGreen)
}

func formatHandlerName(handlerName string) string {
	handlerNameTokens := strings.Split(handlerName, "/")[3:]
	return Style(strings.Join(handlerNameTokens, "/"), ColorCyan)
}

func FormatMethod(method string) string {
	var methodColor int
	switch method {
	case "GET":
		methodColor = ColorGreen
	case "POST":
		methodColor = ColorYellow
	case "PUT":
		methodColor = ColorBlue
	case "DELETE":
		methodColor = ColorRed
	case "PATCH":
		methodColor = ColorMagenta
	default:
		methodColor = ColorCyan
	}

	return Style(fmt.Sprintf("%6s", method), methodColor)
}

func PrintRouter(method, absolutePath, handlerName string, nuHandlers int) {
	method = FormatMethod(method)
	absolutePath = formatPath(absolutePath)
	handlerName = formatHandlerName(handlerName)

	fmt.Printf("%s %s --{%d}-> %s\x1b[0m\n", method, absolutePath, nuHandlers, handlerName)
}
