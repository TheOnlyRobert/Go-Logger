package logger

import (
	"context"
	"io"
	"log"
	"os"
	"time"
)

const (
	ApiKey    string = "api-key"
	StartTime string = "start-time"

	InfoPrefix  string = "INFO "
	WarnPrefix  string = "WARN "
	ErrorPrefix string = "ERROR "
)

type Logger struct {
	InfoLogger  *log.Logger
	WarnLogger  *log.Logger
	ErrorLogger *log.Logger

	InfoOutput  io.Writer
	WarnOutput  io.Writer
	ErrorOutput io.Writer
}

var logger *Logger

func New() *Logger {
	logger = &Logger{
		InfoOutput:  os.Stdout,
		WarnOutput:  os.Stdout,
		ErrorOutput: os.Stderr,
	}

	return logger
}

func (l *Logger) println(logger *log.Logger, logCat LogCat, startTime time.Time, apiKey string, v ...interface{}) {
	msg := BuildMessage(logCat, startTime, apiKey, v...)

	logger.Println(msg)
}

func (l *Logger) Info(ctx context.Context, logCat LogCat, v ...interface{}) {
	if l.InfoLogger == nil {
		l.InfoLogger = log.New(l.InfoOutput, InfoPrefix, log.Ldate|log.Ltime)
	}

	apiKey, _ := ctx.Value(ApiKey).(string)
	startTime, _ := ctx.Value(StartTime).(time.Time)

	l.println(l.InfoLogger, logCat, startTime, apiKey, v...)
}

func (l *Logger) Warn(ctx context.Context, logCat LogCat, v ...interface{}) {
	if l.WarnLogger == nil {
		l.WarnLogger = log.New(l.WarnOutput, WarnPrefix, log.Ldate|log.Ltime)
	}

	apiKey, _ := ctx.Value(ApiKey).(string)
	startTime, _ := ctx.Value(StartTime).(time.Time)

	l.println(l.WarnLogger, logCat, startTime, apiKey, v...)
}

func (l *Logger) Error(ctx context.Context, logCat LogCat, v ...interface{}) {
	if l.ErrorLogger == nil {
		l.ErrorLogger = log.New(l.ErrorOutput, ErrorPrefix, log.Ldate|log.Ltime)
	}

	apiKey, _ := ctx.Value(ApiKey).(string)
	startTime, _ := ctx.Value(StartTime).(time.Time)

	l.println(l.ErrorLogger, logCat, startTime, apiKey, v...)
}

func Info(ctx context.Context, logCat LogCat, v ...interface{}) {
	logger.Info(ctx, logCat, v...)
}

func Warn(ctx context.Context, logCat LogCat, v ...interface{}) {
	logger.Warn(ctx, logCat, v...)
}

func Error(ctx context.Context, logCat LogCat, v ...interface{}) {
	logger.Error(ctx, logCat, v...)
}
