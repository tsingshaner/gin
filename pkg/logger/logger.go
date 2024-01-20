package logger

import (
	"fmt"
	"log"
	"os"

	"github.com/lab-online/pkg/color"
)

const (
	LevelDebug = 1 << iota
	LevelInfo
	LevelWarn
	LevelError
)

var level int

func InitLog(level int) {
	log.SetFlags(log.Lshortfile)
	SetLevel(level)
}

var debug = NewLogger(color.PrefixDebug)

func Debug(v ...any) {
	if level&LevelDebug != 0 {
		debug.Println(v...)
	}
}

var info = NewLogger(color.PrefixInfo)

func Info(msg ...any) {
	if level&LevelInfo != 0 {
		info.Println(msg...)
	}
}

var warn = NewLogger(color.PrefixWarn)

func Warn(msg ...any) {
	if level&LevelWarn != 0 {
		warn.Println(msg...)
	}
}

var errorLogger = NewLogger(color.PrefixError)

func Error(msg ...any) {
	if level&LevelError != 0 {
		errorLogger.Println(msg...)
	}
}

func NewLogger(prefix string) *log.Logger {
	return log.New(os.Stdout, prefix+" ", 0)
}

func SetLevel(l int) {
	fmt.Println(color.PrefixInfo, "log level:", l)
	level = l
}
