package log

import (
	"os"

	"github.com/ethereum/go-ethereum/log"
)

func InitLogger() {
	log.SetDefault(log.NewLogger(log.NewTerminalHandlerWithLevel(os.Stdout, log.LevelDebug, true)))
}

func Debug(format string, v ...interface{}) {
	log.Debug(format, v...)
}

func Info(format string, v ...interface{}) {
	log.Info(format, v...)
}

func Warn(format string, v ...interface{}) {
	log.Warn(format, v...)
}

func Error(format string, v ...interface{}) {
	log.Error(format, v...)
}
