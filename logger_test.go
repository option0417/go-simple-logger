package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestWriteLogToFile(t *testing.T) {
	infoMsg := "Test Info message"
	debugMsg := "Test Debug message"
	errorMsg := "Test Error message"

	// Create Logger
	logger := NewLogger()

	// Write log to file
	logger.Info(infoMsg)
	logger.Debug(debugMsg)
	logger.Error(errorMsg)

	// Verify the content of the log file
	// Find *.log file in the current directory
	os.Chdir("./")
	files, err := ioutil.ReadDir("./")
	if err != nil {
		t.Error(err)
	}
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".log") {
			t.Log("Log file found: ", file.Name())

			// Read the log file
			content, err := ioutil.ReadFile(file.Name())
			if err != nil {
				t.Error(err)
			}
			// Verify the content of the log file
			if !strings.Contains(string(content), infoMsg) {
				t.Error("Info message not found in the log file")
			}
			if !strings.Contains(string(content), debugMsg) {
				t.Error("Debug message not found in the log file")
			}
			if !strings.Contains(string(content), errorMsg) {
				t.Error("Error message not found in the log file")
			}
		}
	}
}
