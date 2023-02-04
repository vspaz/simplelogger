package logging

import (
	"github.com/sirupsen/logrus"
	"sync"
)

type Level uint32

var once sync.Once

const (
	PanicLevel = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

func getLogLevel(logLevel string) int {
	switch logLevel {
	case "panic":
		return
	case "fatal":
		return FatalLevel
	case "error":
		return ErrorLevel
	case "warning":
		return WarnLevel
	case "info":
		return InfoLevel
	case "debug":
		return DebugLevel
	case "trace":
		return TraceLevel
	default:
		return InfoLevel
	}
}

func formatterFactory(formatterType string) logrus.Formatter {
	switch formatterType {
	case "json":
		formatter := new(logrus.JSONFormatter)
		formatter.PrettyPrint = true
		return formatter
	case "text":
		formatter := new(logrus.TextFormatter)
		formatter.TimestampFormat = "2006-01-02 15:04:05.000"
		formatter.FullTimestamp = true
		return formatter
	default:
		panic("invalid formatter: " + formatterType)
	}
}

func createLogger(logLevel logrus.Level, formatter *logrus.Formatter) *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(*formatter)
	logger.SetReportCaller(true)
	logger.SetLevel(logLevel)
	return logger
}

func configureLogger(logLevel string, formatterType string) *logrus.Logger {
	formatter := formatterFactory(formatterType)
	logrusLogLevel := logrus.Level(getLogLevel(logLevel))
	return createLogger(logrusLogLevel, &formatter)
}

func setLogLevel(logLevels ...string) string {
	logLevel := "info"
	if len(logLevels) > 0 {
		logLevel = logLevels[0]
	}
	return logLevel
}

type singletonLogger struct {
	Logger *logrus.Logger
}

var Logger *singletonLogger

func GetTextLogger(logLevels ...string) *singletonLogger {
	once.Do(
		func() {
			Logger = &singletonLogger{
				Logger: configureLogger(setLogLevel(logLevels...), "text"),
			}
		},
	)
	return Logger
}

func GetJsonLogger(logLevels ...string) *singletonLogger {
	once.Do(
		func() {
			Logger = &singletonLogger{
				Logger: configureLogger(setLogLevel(logLevels...), "json"),
			}
		},
	)
	return Logger
}
