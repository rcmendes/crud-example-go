package logs

import (
	"io"
	"log"
	"os"
)

type LoggerBuilder interface {
	SetWriter(out io.Writer)
	SetFlags(flag int)
	SetLogLevel(level int)
	Build() Logger
}

type builder struct {
	output io.Writer
	flags  int
	level  int
}

func NewLoggerBuilder() LoggerBuilder {
	return &builder{
		output: os.Stdout,
		flags:  log.Ldate | log.Ltime | log.Lshortfile,
		level:  LogLevelInfo,
	}
}

func (b *builder) Build() Logger {
	return newLogger(b.output, b.flags, b.level)
}

func (b *builder) SetWriter(out io.Writer) {
	b.output = out
}
func (b *builder) SetFlags(flags int) {
	b.flags = flags
}

func (b *builder) SetLogLevel(level int) {
	b.level = level
}

// output, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
// if err != nil {
// 	log.Fatal(err)
// }
