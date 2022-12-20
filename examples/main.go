package main

import "github.com/vspaz/simplelogger/pkg/logging"

func main() {
	textLogger := logging.GetTextLogger("info").Logger
	textLogger.Info("foobar")

	jsonLogger := logging.GetJsonLogger("info").Logger
	jsonLogger.Info("foobar")
}
