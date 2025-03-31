package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

type LogLevel int

const (
	Debug LogLevel = iota + 1
	Info
	Error
)

const MaxFileSize = int(1 << 30) // 1 GB

var LevelPrefix = map[LogLevel]string{
	Debug: "DEBUG",
	Info:  "INFO",
	Error: "ERROR",
}

type Logger struct {
	iLogger     *log.Logger // Logger for Info Level
	dLogger     *log.Logger // Logger for Debug Level
	eLogger     *log.Logger // Logger for Error Level
	logFile     *os.File    // File to write logs
	maxFileSize int64       // Max file size
	fileName    string      // File name
	level       LogLevel    // Log level
	no          int         // Number of log files
	lock        sync.Mutex  // Mutex lock
}

func NewLogger() *Logger {
	f, err := buildLogFile()
	if err != nil {
		fmt.Println("Failed to open log file")
		panic(err)
	}

	return &Logger{
		iLogger:     log.New(f, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		dLogger:     log.New(f, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile),
		eLogger:     log.New(f, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
		logFile:     f,
		maxFileSize: 100 << 20,
		fileName:    f.Name(),
		level:       Debug,
		no:          0,
		lock:        sync.Mutex{},
	}
}

// Build file for logger with datetime as filename
func buildLogFile() (*os.File, error) {
	// Get the current time
	currTime := time.Now()

	// Build the filename
	filename := fmt.Sprintf("log-%s.log", currTime.Format(time.DateOnly))
	// Create the file
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("Failed to open log file")
	}
	return file, nil
}

func (l *Logger) SetLevel(level LogLevel) {
	l.level = level
}

// Set the max file size for the logger between 1MB and 1024MB
func (l *Logger) SetMaxFileSize(maxFileSize int64) error {
	if maxFileSize <= 0 || maxFileSize > 1024 {
		fmt.Println("Invalid max file size")
		return errors.New("Invalid max file size")
	}

	l.maxFileSize = maxFileSize << 20
	fmt.Printf("Max File Size: %d\n", l.maxFileSize)
	return nil
}

// Rotate the log file when it reaches the max size limit and create a new one with number of log files
func (l *Logger) rotateLogFile() error {
	l.lock.Lock()
	defer l.lock.Unlock()

	// Change file name of Log file
	l.logFile.Close()

	// Rename the file
	l.no = l.no + 1
	tmpName := l.logFile.Name()[:len(l.logFile.Name())-4]
	os.Rename(l.logFile.Name(), fmt.Sprintf("%s.%d.log", tmpName, l.no))

	// Create new log file
	currTime := time.Now()
	newFilename := fmt.Sprintf("log-%s.log", currTime.Format(time.DateOnly))
	// Create the file
	newFile, err := os.OpenFile(newFilename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("Failed to open log file")
	}

	l.logFile = newFile
	l.dLogger.SetOutput(newFile)
	l.iLogger.SetOutput(newFile)
	l.eLogger.SetOutput(newFile)

	return nil
}

func (l *Logger) write(level LogLevel, message string) {
	// Check if the file size is greater than the max file size
	fInfo, err := l.logFile.Stat()
	if err != nil {
		fmt.Println("Failed to get file info")
		panic(err)
	}

	if fInfo.Size() > l.maxFileSize {
		err := l.rotateLogFile()
		if err != nil {
			fmt.Println("Failed to rotate log file")
			panic(err)
		}
	}

	switch level {
	case Debug:
		if l.level == Debug {
			l.dLogger.Println(message)
		}
	case Info:
		l.iLogger.Println(message)
	case Error:
		l.eLogger.Println(message)
	}
}

func (l *Logger) Debug(message string) {
	l.write(Debug, message)
}

func (l *Logger) Info(message string) {
	l.write(Info, message)
}

func (l *Logger) Error(message string) {
	l.write(Error, message)
}

func (l *Logger) Close() {
	// Close the loggers
	l.logFile.Close()
}
