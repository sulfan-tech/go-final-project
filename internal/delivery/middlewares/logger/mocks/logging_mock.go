package mocks

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockLogger is a mock implementation of the Logger interface.
type MockLogger struct {
	DebugFunc   func(message string)
	InfoFunc    func(message string)
	WarningFunc func(message string)
	ErrorFunc   func(message string)
}

// Debug mocks the Debug method.
func (m *MockLogger) Debug(message string) {
	if m.DebugFunc != nil {
		m.DebugFunc(message)
	}
}

// Info mocks the Info method.
func (m *MockLogger) Info(message string) {
	if m.InfoFunc != nil {
		m.InfoFunc(message)
	}
}

// Warning mocks the Warning method.
func (m *MockLogger) Warning(message string) {
	if m.WarningFunc != nil {
		m.WarningFunc(message)
	}
}

// Error mocks the Error method.
func (m *MockLogger) Error(message string) {
	if m.ErrorFunc != nil {
		m.ErrorFunc(message)
	}
}

// TestLogger returns a new instance of MockLogger for testing.
func TestLogger(t *testing.T) *MockLogger {
	return &MockLogger{}
}

// TestLoggerWithBuffer returns a new instance of MockLogger that logs to a buffer for testing.
func TestLoggerWithBuffer(t *testing.T) (*MockLogger, *bytes.Buffer) {
	var buf bytes.Buffer
	logger := &MockLogger{
		DebugFunc: func(message string) {
			buf.WriteString("[DEBUG] " + message + "\n")
		},
		InfoFunc: func(message string) {
			buf.WriteString("[INFO] " + message + "\n")
		},
		WarningFunc: func(message string) {
			buf.WriteString("[WARNING] " + message + "\n")
		},
		ErrorFunc: func(message string) {
			buf.WriteString("[ERROR] " + message + "\n")
		},
	}
	return logger, &buf
}

// AssertLogged asserts that the given buffer contains the expected logged messages.
func AssertLogged(t *testing.T, buf *bytes.Buffer, expectedMessages ...string) {
	loggedMessages := buf.String()
	for _, msg := range expectedMessages {
		assert.Contains(t, loggedMessages, msg)
	}
}
