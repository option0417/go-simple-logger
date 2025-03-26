package main

import (
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

var LevelPrefix = map[LogLevel]string{
	Debug: "DEBUG",
	Info:  "INFO",
	Error: "ERROR",
}

type Logger struct {
	iLogger  *log.Logger // Logger for Info Level
	dLogger  *log.Logger // Logger for Debug Level
	eLogger  *log.Logger // Logger for Error Level
	logFile  *os.File    // File to write logs
	fizeSize uint64      // Max file size
	fileName string      // File name
	level    LogLevel    // Log level
	no       int         // Number of log files
	lock     sync.Mutex  // Mutex lock
}

func NewLogger() *Logger {
	logFile, err := buildLogFile()
	if err != nil {
		fmt.Println("Failed to open log file")
		panic(err)
	}

	return &Logger{
		iLogger: log.New(logFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		dLogger: log.New(logFile, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile),
		eLogger: log.New(logFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
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

// Rotate the log file when it reaches the max size limit and create a new one with number of log files
func (l *Logger) rotateLogFile() error {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.logFile.Close()

	l.no = l.no + 1
	newFilename := fmt.Sprintf("log-%s.%d.log", l.fileName, l.no)

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

func (l *Logger) SetLevel(level LogLevel) {
	l.level = level
}

func (l *Logger) write(level LogLevel, message string) {
	switch level {
	case Debug:
		l.dLogger.Println(message)
	case Info:
		l.iLogger.Println(message)
	case Error:
		l.eLogger.Println(message)
	}
}

func (l *Logger) Info(message string) {
	l.write(Info, message)
}

func (l *Logger) Debug(message string) {
	l.write(Debug, message)
}

func (l *Logger) Error(message string) {
	l.write(Error, message)
}

func (l *Logger) Close() {
	// Close the loggers
	l.logFile.Close()
}
