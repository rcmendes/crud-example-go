package logs

import (
	"io"
	"log"
)

const LogLevelDebug = 1
const LogLevelInfo = 2
const LogLevelWarning = 3
const LogLevelError = 4

type Logger interface {
	Debug(message ...interface{})
	Info(message ...interface{})
	Warning(message ...interface{})
	Error(message ...interface{})
}

type logger struct {
	debugLogger   *log.Logger
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
	level         int
}

func newLogger(output io.Writer, flags int, logLevel int) Logger {
	return &logger{
		debugLogger:   log.New(output, "DEBUG: ", flags),
		infoLogger:    log.New(output, "INFO: ", flags),
		warningLogger: log.New(output, "WARNING: ", flags),
		errorLogger:   log.New(output, "ERROR: ", flags),
		level:         logLevel,
	}
}

func (l *logger) Debug(message ...interface{}) {
	if l.level == LogLevelDebug {
		l.debugLogger.Println(message...)
	}
}

func (l *logger) Info(message ...interface{}) {
	if l.level <= LogLevelInfo {
		l.infoLogger.Println(message...)
	}
}

func (l *logger) Warning(message ...interface{}) {
	if l.level <= LogLevelWarning {
		l.warningLogger.Println(message...)
	}
}

func (l *logger) Error(message ...interface{}) {
	if l.level <= LogLevelError {
		l.errorLogger.Println(message...)
	}
}
