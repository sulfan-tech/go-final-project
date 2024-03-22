package logger

import (
	"log"
	"os"
)

// LogLevel type represents different log levels.
type LogLevel int

// Log levels.
const (
	DebugLevel LogLevel = iota
	InfoLevel
	WarningLevel
	ErrorLevel
)

// Logger interface represents a logger instance.
type Logger interface {
	Debug(message string)
	Info(message string)
	Warning(message string)
	Error(message string)
}

// LoggerImpl struct represents an implementation of the Logger interface.
type LoggerImpl struct {
	logger *log.Logger
}

// NewLogger creates a new Logger instance with the given log level.
func NewLogger(level LogLevel) *LoggerImpl {
	var prefix string
	switch level {
	case DebugLevel:
		prefix = "[DEBUG] "
	case InfoLevel:
		prefix = "[INFO] "
	case WarningLevel:
		prefix = "[WARNING] "
	case ErrorLevel:
		prefix = "[ERROR] "
	}

	return &LoggerImpl{
		logger: log.New(os.Stdout, prefix, log.Ldate|log.Ltime),
	}
}

// Debug logs a debug message.
func (l *LoggerImpl) Debug(message string) {
	l.logger.Println(message)
}

// Info logs an informational message.
func (l *LoggerImpl) Info(message string) {
	l.logger.Println(message)
}

// Warning logs a warning message.
func (l *LoggerImpl) Warning(message string) {
	l.logger.Println(message)
}

// Error logs an error message.
func (l *LoggerImpl) Error(message string) {
	l.logger.Println(message)
}
