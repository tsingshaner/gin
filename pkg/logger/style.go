package logger

import "fmt"

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

	style := ""
	for _, s := range styles {
		style += fmt.Sprintf(";%d", s)
	}
	return fmt.Sprintf("\x1b[%sm%s\x1b[0m", style[1:], text)
}
