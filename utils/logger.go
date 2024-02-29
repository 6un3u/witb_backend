package utils

import (
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"github.com/natefinch/lumberjack"
	slogecho "github.com/samber/slog-echo"
	slogmulti "github.com/samber/slog-multi"
)

var SlogConfig = slogecho.Config{
	WithRequestBody: true,
	// WithResponseBody:   true,
	WithRequestHeader:  true,
	WithResponseHeader: true,
}

// Start: change log setting for ecs
const (
	ecsVersion    = "8.11.0"
	ecsTimeForamt = "2006-01-02T15:04:05.999Z"
)

func GetEcsSlogger() *slog.Logger {
	commonFields := []slog.Attr{
		{
			Key:   "ecs.version",
			Value: slog.StringValue(ecsVersion),
		},
	}

	renameFields := func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.TimeKey {
			a.Key = "@timestamp"
			a.Value = slog.StringValue(time.Now().UTC().Format(ecsTimeForamt))
			return a
		}
		if a.Key == slog.LevelKey {
			a.Key = "log.level"
			return a
		}
		if a.Key == slog.MessageKey {
			a.Key = "message"
			return a
		}

		return a
	}

	newHandlerOptions := &slog.HandlerOptions{
		ReplaceAttr: renameFields,
	}

	ecsSlogger := slog.New(
		slogmulti.Fanout(
			slog.NewJSONHandler(os.Stdout, newHandlerOptions).WithAttrs(commonFields),
			slog.NewJSONHandler(lumberjackLogger, newHandlerOptions).WithAttrs(commonFields),
		),
	)

	return ecsSlogger
}

// Start: Set lumberjack
var lumberjackLogger = &lumberjack.Logger{
	Filename:   getLogFilePath(),
	MaxSize:    20,   // A file can be up to 20M.
	MaxBackups: 5,    // Save up to 5 files at the same time
	MaxAge:     10,   // A file can be saved for up to 10 days.
	Compress:   true, // Compress with gzip.
}

func getLogFilePath() string {
	logPath := os.Getenv("LOG_PATH")
	logFileName := time.Now().Format("2006-01-02") + ".log"

	fileName := filepath.Join(logPath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			log.Println(err.Error())
		}
	}
	return fileName
}
