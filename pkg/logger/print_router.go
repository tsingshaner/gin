package logger

import (
	"fmt"
	"strings"

	"github.com/tsingshaner/gin-starter/pkg/color"
)

func formatPath(path string) string {
	return color.Green(fmt.Sprintf("%-40s", path))
}

func formatHandlerName(handlerName string) string {
	handlerNameTokens := strings.Split(handlerName, "/")[3:]
	return color.Cyan(strings.Join(handlerNameTokens, "/"))
}

func FormatMethod(method string) string {
	methodColor := map[string]int{
		"GET":    color.ColorGreen,
		"POST":   color.ColorYellow,
		"PUT":    color.ColorBlue,
		"DELETE": color.ColorRed,
		"PATCH":  color.ColorMagenta,
	}[method]

	formattedMethod := fmt.Sprintf("%6s", method)

	if methodColor != 0 {
		return color.Style(formattedMethod, methodColor)
	} else {
		return color.Cyan(formattedMethod)
	}
}

func PrintRouter(method, absolutePath, handlerName string, nuHandlers int) {
	method = FormatMethod(method)
	absolutePath = formatPath(absolutePath)
	handlerName = formatHandlerName(handlerName)

	fmt.Printf("%s %s --{%d}-> %s\x1b[0m\n", method, absolutePath, nuHandlers, handlerName)
}
