package logging

import (
	"github.com/sirupsen/logrus"
	"sync"
)

var once sync.Once

func getLogLevel(logLevel string) logrus.Level {
	switch logLevel {
	case "panic":
		return logrus.PanicLevel
	case "fatal":
		return logrus.FatalLevel
	case "error":
		return logrus.ErrorLevel
	case "warning":
		return logrus.WarnLevel
	case "info":
		return logrus.InfoLevel
	case "debug":
		return logrus.DebugLevel
	case "trace":
		return logrus.TraceLevel
	default:
		return logrus.InfoLevel
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
