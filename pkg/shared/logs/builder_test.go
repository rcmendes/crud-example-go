package logs

import (
	"bytes"
	"log"
	"os"
	"testing"
)

func TestShouldBuildDefaultLogger(t *testing.T) {
	builderInstance := NewLoggerBuilder()
	builder, ok := builderInstance.(*builder)

	if !ok {
		t.Error("Fail on create a 'LoggerBuilder' instance.")
	}

	if builder.flags != log.Ldate|log.Ltime|log.Lshortfile {
		t.Error("Fail on set the default flag (log.Ldate|log.Ltime|log.Lshortfile) value.")
	}

	if builder.output != os.Stdout {
		t.Error("Fail on set the default writer (os.Stdout) value.")
	}

	if builder.level != LogLevelInfo {
		t.Error("Fail on set the default log level (INFO) value.")
	}

	logger := builderInstance.Build()

	if logger == nil {
		t.Error("Fail on build/create a 'Logger' instance.")
	}
}

func TestShouldBuildParametrizedLogger(t *testing.T) {
	builderInstance := NewLoggerBuilder()

	builder := builderInstance.(*builder)

	flags := log.Llongfile | log.Ldate

	builderInstance.SetFlags(flags)

	if builder.flags != flags {
		t.Error("Fail on set the parameterized flag value.")
	}

	var buffer bytes.Buffer

	builderInstance.SetWriter(&buffer)

	if builder.output != &buffer {
		t.Error("Fail on set the parameterized Writer value.")
	}

	logLevel := LogLevelError
	builderInstance.SetLogLevel(logLevel)

	if builder.level != logLevel {
		t.Error("Fail on set the parameterized log level value.")
	}

}
