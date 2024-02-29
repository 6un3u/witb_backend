package utils

import (
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"github.com/natefinch/lumberjack"
	slogecho "github.com/samber/slog-echo"
	slogformatter "github.com/samber/slog-formatter"
	slogmulti "github.com/samber/slog-multi"
)

var SlogLogger = slog.New(
	slogformatter.NewFormatterHandler(
		slogformatter.TimezoneConverter(seoulTime),
		slogformatter.TimeFormatter(time.DateTime, nil),
	)(
		slogmulti.Fanout(
			slog.NewJSONHandler(os.Stdout, nil),
			slog.NewJSONHandler(lumberjackLogger, nil),
		),
	),
)

var SlogConfig = slogecho.Config{
	WithRequestBody: true,
	// WithResponseBody:   true,
	WithRequestHeader:  true,
	WithResponseHeader: true,
}

var seoulTime, _ = time.LoadLocation("Asia/Seoul")

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
