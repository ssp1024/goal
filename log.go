package goal

import (
	"fmt"
	"io"
	"os"
	"time"
)

type LogLevel int

const (
	NoneLevel  LogLevel = 0
	DebugLevel LogLevel = 10
	InfoLevel  LogLevel = 20
	WarnLevel  LogLevel = 30
	ErrorLevel LogLevel = 40
	FatalLevel LogLevel = 50
)

var (
	logWriter    io.Writer = os.Stdout
	logLevel               = InfoLevel
	levelStrings           = []string{"NONE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

func (l LogLevel) String() string {
	index := l / 10
	return levelStrings[index]
}

func SetLogLevel(level LogLevel) {
	logLevel = level
}

func SetLogWriter(w io.Writer) {
	logWriter = w
}

func logf(level LogLevel, format string, args ...interface{}) {
	if level < logLevel {
		return
	}
	fmt.Fprintf(logWriter, "%s [%s] %s\n", time.Now().Format("2006-01-02T15:04:05"), level.String(), fmt.Sprintf(format, args...))
}

func Fatalf(format string, args ...interface{}) {
	logf(FatalLevel, format, args...)
	os.Exit(1)
}

func Warnf(format string, args ...interface{}) {
	logf(WarnLevel, format, args...)
}

func Errorf(format string, args ...interface{}) {
	logf(ErrorLevel, format, args...)
}

func Infof(format string, args ...interface{}) {
	logf(InfoLevel, format, args...)
}

func Debugf(format string, args ...interface{}) {
	logf(DebugLevel, format, args...)
}
