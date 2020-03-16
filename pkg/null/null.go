package null

import (
	"github.com/hanaboso/go-log/pkg"
)

// Logger null logger
type Logger struct{}

// Debug debug
func (log *Logger) Debug(message string, args ...interface{}) {}

// Info info
func (log *Logger) Info(message string, args ...interface{}) {}

// Warn warn
func (log *Logger) Warn(message string, args ...interface{}) {}

// Error error
func (log *Logger) Error(err error) {}

// Fatal fatal
func (log *Logger) Fatal(err error) {}

// WithFields preregister fields into logger
func (log *Logger) WithFields(data map[string]interface{}) pkg.Logger {
	return log
}

// SetLevel set level
func (log Logger) SetLevel(level pkg.LogLevel) {}

// NewLogger new logrus logger
func NewLogger() pkg.Logger {
	return &Logger{}
}
