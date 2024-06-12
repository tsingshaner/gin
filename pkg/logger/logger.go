package logger

import (
	"fmt"
	"log"
	"os"

	"github.com/tsingshaner/go-pkg/color"
)

var (
	PrefixDebug = color.UnsafeBold(color.UnsafeBlue("debug:"))
	PrefixInfo  = color.UnsafeBold(color.UnsafeGreen(" info:"))
	PrefixWarn  = color.UnsafeBold(color.UnsafeYellow(" warn:"))
	PrefixError = color.UnsafeBold(color.UnsafeRed("error:"))
	PrefixFatal = color.UnsafeBold(color.UnsafeRed("fatal:"))
)

const (
	LevelDebug = 1 << iota
	LevelInfo
	LevelWarn
	LevelError
)

var level int = LevelDebug | LevelInfo | LevelWarn | LevelError

func InitLog(level int) {
	log.SetFlags(log.Lshortfile)
	SetLevel(level)
}

var debug = NewLogger(PrefixDebug)

func Debug(v ...any) {
	if level&LevelDebug != 0 {
		debug.Println(v...)
	}
}

var info = NewLogger(PrefixInfo)

func Info(msg ...any) {
	if level&LevelInfo != 0 {
		info.Println(msg...)
	}
}

var warn = NewLogger(PrefixWarn)

func Warn(msg ...any) {
	if level&LevelWarn != 0 {
		warn.Println(msg...)
	}
}

var errorLogger = NewLogger(PrefixError)

func Error(msg ...any) {
	if level&LevelError != 0 {
		errorLogger.Println(msg...)
	}
}

func NewLogger(prefix string) *log.Logger {
	return log.New(os.Stdout, prefix+" ", 0)
}

func SetLevel(l int) {
	fmt.Println(PrefixInfo, "log level:", l)
	level = l
}
