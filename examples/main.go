package main

import "github.com/vspaz/simplelogger/pkg/logging"

func main() {
	logger := logging.GetTextLogger("info").Logger
	logger.Info("foobar")
}
