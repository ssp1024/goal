package goal

import (
	"fmt"
	"io"
	"os"
	"time"
)

type logLevel int

const (
	NoneLevel  logLevel = 0
	DebugLevel logLevel = 10
	InfoLevel  logLevel = 20
	WarnLevel  logLevel = 30
	ErrorLevel logLevel = 40
	FatalLevel logLevel = 50
)

var (
	logWriter    io.Writer = os.Stdout
	currLevel              = InfoLevel
	levelStrings           = []string{"NONE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL"}

	now = time.Now
)

func (l logLevel) String() string {
	index := l / 10
	return levelStrings[index]
}

//SetLogLevel update log level, default value is InfoLevel.
//	Debugf("NOT SHOW")
// 	Infof("info msg")
// 	Warnf("warn msg")
// 	SetLogLevel(DebugLevel)
// 	Debugf("debug msg")
//
// 	//Output:
// 	// 2020-01-11T18:03:00 [INFO] info msg
// 	// 2020-01-11T18:03:00 [WARN] warn msg
// 	// 2020-01-11T18:03:00 [DEBUG] debug msg
func SetLogLevel(level logLevel) {
	currLevel = level
}

//SetLogWriter update log message writer, default value is STDOUT.
func SetLogWriter(w io.Writer) {
	logWriter = w
}

func logf(level logLevel, format string, args ...interface{}) {
	if level < currLevel {
		return
	}
	fmt.Fprintf(logWriter, "%s [%s] %s\n", now().Format("2006-01-02T15:04:05"), level.String(), fmt.Sprintf(format, args...))
}

//Fatalf log fatal message, then exit process with code 1.
func Fatalf(format string, args ...interface{}) {
	logf(FatalLevel, format, args...)
	os.Exit(1)
}

//Warnf log warn message.
func Warnf(format string, args ...interface{}) {
	logf(WarnLevel, format, args...)
}

//Errorf log error message.
func Errorf(format string, args ...interface{}) {
	logf(ErrorLevel, format, args...)
}

//Infof log info message.
func Infof(format string, args ...interface{}) {
	logf(InfoLevel, format, args...)
}

//Debugf log debug message.
func Debugf(format string, args ...interface{}) {
	logf(DebugLevel, format, args...)
}
