package log

import (
	"os"
	"strings"
	"testing"
)

const (
	InfoMsg  = "Test Info message"
	DebugMsg = "Test Debug message"
	ErrorMsg = "Test Error message"
)

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
		}
	}

	if cnt != 11 {
		t.Error("Log file rotated incorrectly")
	}

	removeLogFile()
}

func TestWriteLogWithFormat(t *testing.T) {
	// Create Logger
	logger := NewLogger()

	logger.Debugf("Debug message with string: %s", "test-string")
	logger.Debugf("Debug message with integer: %d", 12345)
	logger.Debugf("Debug message with object: %v", logger)

	logger.Infof("Info message with string: %s", "test-string")
	logger.Infof("Info message with integer: %d", 12345)
	logger.Infof("Info message with object: %v", logger)

	logger.Errorf("Error message with string: %s", "test-string")
	logger.Errorf("Error message with integer: %d", 12345)
	logger.Errorf("Error message with object: %v", logger)

	contentBytes, err := fetchLogContent()
	if err != nil {
		t.Error(err)
	}

	content := string(contentBytes)

	// Verify the content of the log file
	if !strings.Contains(content, "Debug message with string: test-string") {
		t.Error("Debug message with string not found in the log file")
	}
	if !strings.Contains(content, "Debug message with integer: 12345") {
		t.Error("Debug message with integer not found in the log file")
	}
	if !strings.Contains(content, "Debug message with object") {
		t.Error("Debug message with object not found in the log file")
	}

	if !strings.Contains(content, "Info message with string: test-string") {
		t.Error("Info message with string not found in the log file")
	}
	if !strings.Contains(content, "Info message with integer: 12345") {
		t.Error("Info message with integer not found in the log file")
	}
	if !strings.Contains(content, "Info message with object") {
		t.Error("Info message with object not found in the log file")
	}

	if !strings.Contains(content, "Error message with string: test-string") {
		t.Error("Error message with string not found in the log file")
	}
	if !strings.Contains(content, "Error message with integer: 12345") {
		t.Error("Error message with integer not found in the log file")
	}
	if !strings.Contains(content, "Error message with object") {
		t.Error("Error message with object not found in the log file")
	}

	removeLogFile()
}

func BenchmarkWriteLogToFile(b *testing.B) {
	logger := NewLogger()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info(InfoMsg)
	}
}
func BenchmarkWriteLogToFile2(b *testing.B) {
	logger := NewLogger()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info(InfoMsg)
	}
}

func BenchmarkWriteLogToFile3(b *testing.B) {
	logger := NewLogger()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info(InfoMsg)
	}
}

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
