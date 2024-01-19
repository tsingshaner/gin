package color

import (
	"fmt"
	"strings"
)

const (
	ColorBlack = 30 + iota
	ColorRed
	ColorGreen
	ColorYellow
	ColorBlue
	ColorMagenta
	ColorCyan
	ColorWhite
)

const (
	StyleReset = iota
	FontBold
	FontDim
	FontItalic
	FontUnderline
)

func Style(text string, styles ...int) string {
	if len(styles) == 0 {
		return text
	}

	var style strings.Builder
	for _, s := range styles {
		style.WriteString(fmt.Sprintf(";%d", s))
	}
	return fmt.Sprintf("\x1b[%sm%s\x1b[0m", style.String()[1:], text)
}

func style(text string, style int) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", style, text)
}

func Red(text string) string {
	return style(text, ColorRed)
}

func Green(text string) string {
	return style(text, ColorGreen)
}

func Yellow(text string) string {
	return style(text, ColorYellow)
}

func Blue(text string) string {
	return style(text, ColorBlue)
}

func Magenta(text string) string {
	return style(text, ColorMagenta)
}

func Cyan(text string) string {
	return style(text, ColorCyan)
}

func White(text string) string {
	return style(text, ColorWhite)
}

func Black(text string) string {
	return style(text, ColorBlack)
}

func Bold(text string) string {
	return style(text, FontBold)
}

func Dim(text string) string {
	return style(text, FontDim)
}

func Italic(text string) string {
	return style(text, FontItalic)
}

func Underline(text string) string {
	return style(text, FontUnderline)
}

func Log[T string | int](tag T, color int, msg []any) {
	fmt.Print(Style(fmt.Sprintf("%5v: ", tag), color, FontBold))
	fmt.Println(msg...)
}

var (
	PrefixDebug = Style("debug:", ColorBlue, FontBold)
	PrefixInfo  = Style(" info:", ColorGreen, FontBold)
	PrefixWarn  = Style(" warn:", ColorYellow, FontBold)
	PrefixError = Style("error:", ColorRed, FontBold)
	PrefixFatal = Style("fatal:", ColorRed, FontBold)
)
