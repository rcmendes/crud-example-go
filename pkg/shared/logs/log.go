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
	debugLogger.Println(message...)
}

func Info(message ...interface{}) {
	infoLogger.Println(message...)
}

func Warning(message ...interface{}) {
	warningLogger.Println(message...)
}

func Error(message ...interface{}) {
	errorLogger.Println(message...)
}
