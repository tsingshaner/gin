package logger

import (
	"fmt"
)

func Log[T string | int](tag T, color int, msg []any) {
	fmt.Print(Style(fmt.Sprintf("%5v: ", tag), color, FontBold))
	fmt.Println(msg...)
}

func Info(msg ...any) {
	Log("info", ColorGreen, msg)
}

func Error(msg ...any) {
	Log("error", ColorRed, msg)
}

func Warn(msg ...any) {
	Log("warn", ColorYellow, msg)
}

func Http(msg ...any) {
	Log("http", ColorCyan, msg)
}
