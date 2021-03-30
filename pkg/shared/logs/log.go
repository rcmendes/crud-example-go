package logs

import (
	"log"
	"os"
)

var (
	debugLogger   *log.Logger
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
)

const LogLevelDebug = 1
const LogLevelInfo = 2
const LogLevelWarning = 3
const LogLevelError = 4

var logLevel = LogLevelInfo

func init() {
	// output, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	output := os.Stdout

	defaultFlags := log.Ldate | log.Ltime | log.Lshortfile

	debugLogger = log.New(output, "DEBUG: ", defaultFlags)
	infoLogger = log.New(output, "INFO: ", defaultFlags)
	warningLogger = log.New(output, "WARNING: ", defaultFlags)
	errorLogger = log.New(output, "ERROR: ", defaultFlags)
}

func Debug(message ...interface{}) {
	if logLevel == LogLevelDebug {
		debugLogger.Println(message...)
	}
}

func Info(message ...interface{}) {
	if logLevel <= LogLevelInfo {
		infoLogger.Println(message...)
	}
}

func Warning(message ...interface{}) {
	if logLevel <= LogLevelWarning {
		warningLogger.Println(message...)
	}
}

func Error(message ...interface{}) {
	if logLevel <= LogLevelError {
		errorLogger.Println(message...)
	}
}
