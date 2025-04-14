package main

import (
	"tw.com.wd.utils/logger/log"
)

func main() {
	logger := log.NewLogger()
	defer logger.Close()

	logger.Info("This is an information message")
	logger.Debug("This is a debug message")
	logger.Error("This is an error message")

}
