package log

import (
	// "fmt"
	"log/slog"
	// "os"
)

func InitLogger() {
	// TODO

	// opts := &slog.HandlerOptions{
	// 	Level: slog.LevelInfo,
	// 	ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
	// 		if a.Key == slog.TimeKey {
	// 			t := a.Value.Time()
	// 			format := fmt.Sprintf(
	// 				"[%s|%s]",
	// 				t.Format("01-02"),
	// 				t.Format("15:04:05.000"),
	// 			)
	// 			a.Value = slog.StringValue(format)
	// 		}
	// 		return a
	// 	},
	// }

	// logger := slog.NewTextHandler(os.Stderr, opts)
	// slog.SetDefault(logger)
}

func Debug(format string, v ...interface{}) {
	slog.Debug(format, v...)
}

func Info(format string, v ...interface{}) {
	slog.Info(format, v...)
}

func Warn(format string, v ...interface{}) {
	slog.Warn(format, v...)
}

func Error(format string, v ...interface{}) {
	slog.Error(format, v...)
}
