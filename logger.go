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
	iLogger  *log.Logger
	dLogger  *log.Logger
	eLogger  *log.Logger
	logFile  *os.File
	fileName string
	level    LogLevel
	no       int
	lock     sync.Mutex
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
	filename := fmt.Sprintf("log-%s.txt", currTime.Format(time.DateOnly))
	// Create the file
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("Failed to open log file")
	}
	return file, nil
}

func (l *Logger) Info(message string) {
	l.lock.Lock()
	l.iLogger.Println(message)
	l.lock.Unlock()
}

func (l *Logger) Debug(message string) {
	l.lock.Lock()
	l.dLogger.Println(message)
	l.lock.Unlock()
}

func (l *Logger) Error(message string) {
	l.lock.Lock()
	l.eLogger.Println(message)
	l.lock.Unlock()
}

func (l *Logger) Close() {
	// Close the loggers
	l.logFile.Close()
}
