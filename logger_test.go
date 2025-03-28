package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

const (
	InfoMsg  = "Test Info message"
	DebugMsg = "Test Debug message"
	ErrorMsg = "Test Error message"
)

func removeLogFile() {
	files, err := os.ReadDir("./")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".log") {
			err := os.Remove(file.Name())
			if err != nil {
				panic(err)
			}
		}
	}
}

func fetchLogContent() ([]byte, error) {
	files, err := os.ReadDir("./")
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".log") {

			// Read the log file
			content, err := os.ReadFile(file.Name())
			if err != nil {
				return nil, err
			}
			return content, nil
		}
	}
	return nil, nil
}

func TestWriteLogToFile(t *testing.T) {
	// Create Logger
	logger := NewLogger()

	// Write log to file
	logger.Debug(DebugMsg)
	logger.Info(InfoMsg)
	logger.Error(ErrorMsg)

	contentBytes, err := fetchLogContent()
	if err != nil {
		t.Error(err)
	}

	content := string(contentBytes)

	// Verify the content of the log file
	if !strings.Contains(content, DebugMsg) {
		t.Error("Debug message not found in the log file")
	}
	if !strings.Contains(content, InfoMsg) {
		t.Error("Info message not found in the log file")
	}
	if !strings.Contains(content, ErrorMsg) {
		t.Error("Error message not found in the log file")
	}

	removeLogFile()
}

func TestWriteLogWithDebugLevel(t *testing.T) {
	logger := NewLogger()
	logger.SetLevel(Debug)

	// Write log to file
	logger.Debug(DebugMsg)
	logger.Info(InfoMsg)
	logger.Error(ErrorMsg)

	contentBytes, err := fetchLogContent()
	if err != nil {
		t.Error(err)
	}

	content := string(contentBytes)

	// Verify the content of the log file
	if !strings.Contains(content, DebugMsg) {
		t.Error("Debug message not found in the log file")
	}
	if !strings.Contains(content, InfoMsg) {
		t.Error("Info message not found in the log file")
	}
	if !strings.Contains(content, ErrorMsg) {
		t.Error("Error message not found in the log file")
	}

	removeLogFile()
}

func TestWriteLogWithInfoLevel(t *testing.T) {
	logger := NewLogger()
	logger.SetLevel(Info)

	// Write log to file
	logger.Debug(DebugMsg)
	logger.Info(InfoMsg)
	logger.Error(ErrorMsg)

	contentBytes, err := fetchLogContent()
	if err != nil {
		t.Error(err)
	}

	content := string(contentBytes)

	// Verify the content of the log file
	if strings.Contains(content, DebugMsg) {
		t.Error("Debug message found in the log file")
	}
	if !strings.Contains(content, InfoMsg) {
		t.Error("Info message not found in the log file")
	}
	if !strings.Contains(content, ErrorMsg) {
		t.Error("Error message not found in the log file")
	}

	removeLogFile()
}

func TestWriteLogWithErrorLevel(t *testing.T) {
	logger := NewLogger()
	logger.SetLevel(Error)

	// Write log to file
	logger.Debug(DebugMsg)
	logger.Info(InfoMsg)
	logger.Error(ErrorMsg)

	contentBytes, err := fetchLogContent()
	if err != nil {
		t.Error(err)
	}

	content := string(contentBytes)

	// Verify the content of the log file
	if strings.Contains(content, DebugMsg) {
		t.Error("Debug message found in the log file")
	}
	if !strings.Contains(content, InfoMsg) {
		t.Error("Info message not found in the log file")
	}
	if !strings.Contains(content, ErrorMsg) {
		t.Error("Error message not found in the log file")
	}

	removeLogFile()
}

func TestRotateLog(t *testing.T) {
	// Create Logger
	logger := NewLogger()

	// Set max-file-zise to 1MB
	logger.SetMaxFileSize(1)

	// Write log to file over 1MB
	// One line of log at least 41 bytes
	// 1024 * 1024 = 1048576 bytes
	// 41 * 255770 = 10485770 bytes almost 11 Log files
	for i := 0; i < 255770; i++ {
		//logger.Info(strconv.Itoa(i))
		logger.Info("")
	}

	cnt := 0

	files, err := os.ReadDir("./")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".log") {
			cnt++
			fmt.Printf("Found log file: %s\n", file.Name())
		}
	}

	if cnt != 11 {
		t.Error("Log file rotated incorrectly")
	}

	removeLogFile()
}
