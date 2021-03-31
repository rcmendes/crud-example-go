package logs

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
)

func TestShouldCreateLogger(t *testing.T) {
	loggerInstance := newLogger(os.Stdout, log.Ldate, LogLevelError)

	if loggerInstance == nil {
		t.Errorf("Failed on create a logger.")
	}

	instance := loggerInstance.(*logger)

	if instance.debugLogger == nil {
		t.Errorf("Failed on set the DEBUG log in the Logger instance.")
	}

	if instance.infoLogger == nil {
		t.Errorf("Failed on set the INFO log in the Logger instance.")
	}

	if instance.warningLogger == nil {
		t.Errorf("Failed on set the WARNING log in the Logger instance.")
	}

	if instance.errorLogger == nil {
		t.Errorf("Failed on set the ERROR log in the Logger instance.")
	}

	if instance.level != LogLevelError {
		t.Errorf("Failed on set the log level the Logger instance.")
	}
}

func TestShouldLogDebug(t *testing.T) {
	var buffer bytes.Buffer

	logger := newLogger(&buffer, log.Ldate, LogLevelDebug)

	message := "Log Testing 123"

	logger.Debug(message)
	logger.Info(message)
	logger.Warning(message)
	logger.Error(message)

	data := buffer.String()

	if !strings.Contains(data, message) {
		t.Errorf("Fail on log message '%s'.\n", message)
	}

	if !strings.Contains(data, "DEBUG") {
		t.Errorf("Should log in: '%s'.", "DEBUG")
	}

	if !strings.Contains(data, "INFO") {
		t.Errorf("Should log in: '%s'.", "INFO")
	}

	if !strings.Contains(data, "WARNING") {
		t.Errorf("Should log in: '%s'.", "WARNING")
	}

	if !strings.Contains(data, "ERROR") {
		t.Errorf("Should log in: '%s'.", "ERROR")
	}

}

func TestShouldLogInfo(t *testing.T) {
	var buffer bytes.Buffer

	logger := newLogger(&buffer, log.Ldate, LogLevelInfo)

	message := "Log Testing 123"

	logger.Debug(message)
	logger.Info(message)
	logger.Warning(message)
	logger.Error(message)

	data := buffer.String()

	if !strings.Contains(data, message) {
		t.Errorf("Fail on log message '%s'.\n", message)
	}

	if strings.Contains(data, "DEBUG") {
		t.Errorf("Should NOT log in: '%s'.", "DEBUG")
	}

	if !strings.Contains(data, "INFO") {
		t.Errorf("Should log in: '%s'.", "INFO")
	}

	if !strings.Contains(data, "WARNING") {
		t.Errorf("Should log in: '%s'.", "WARNING")
	}

	if !strings.Contains(data, "ERROR") {
		t.Errorf("Should log in: '%s'.", "ERROR")
	}

}

func TestShouldLogWarning(t *testing.T) {
	var buffer bytes.Buffer

	logger := newLogger(&buffer, log.Ldate, LogLevelWarning)

	message := "Log Testing 123"

	logger.Debug(message)
	logger.Info(message)
	logger.Warning(message)
	logger.Error(message)

	data := buffer.String()

	if !strings.Contains(data, message) {
		t.Errorf("Fail on log message '%s'.\n", message)
	}

	if strings.Contains(data, "DEBUG") {
		t.Errorf("Should NOT log in: '%s'.", "DEBUG")
	}

	if strings.Contains(data, "INFO") {
		t.Errorf("Should NOT log in: '%s'.", "INFO")
	}

	if !strings.Contains(data, "WARNING") {
		t.Errorf("Should log in: '%s'.", "WARNING")
	}

	if !strings.Contains(data, "ERROR") {
		t.Errorf("Should log in: '%s'.", "ERROR")
	}

}

func TestShouldLogError(t *testing.T) {
	var buffer bytes.Buffer

	logger := newLogger(&buffer, log.Ldate, LogLevelError)

	message := "Log Testing 123"

	logger.Debug(message)
	logger.Info(message)
	logger.Warning(message)
	logger.Error(message)

	data := buffer.String()

	if !strings.Contains(data, message) {
		t.Errorf("Fail on log message '%s'.\n", message)
	}

	if strings.Contains(data, "DEBUG") {
		t.Errorf("Should NOT log in: '%s'.", "DEBUG")
	}

	if strings.Contains(data, "INFO") {
		t.Errorf("Should NOT log in: '%s'.", "INFO")
	}

	if strings.Contains(data, "WARNING") {
		t.Errorf("Should NOT log in: '%s'.", "WARNING")
	}

	if !strings.Contains(data, "ERROR") {
		t.Errorf("Should log in: '%s'.", "ERROR")
	}

}
