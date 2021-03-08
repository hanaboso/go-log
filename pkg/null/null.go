package null

import (
	"github.com/hanaboso/go-log/pkg"
)

// Logger null logger
type Logger struct{}

// Debug debug
func (log Logger) Debug(_ string, _ ...interface{}) {}

// Info info
func (log Logger) Info(_ string, _ ...interface{}) {}

// Warn warn
func (log Logger) Warn(_ string, _ ...interface{}) {}

// Error error
func (log Logger) Error(_ error) {}

// ErrorWrap wraps error
func (log Logger) ErrorWrap(_ string, _ error) {}

// Fatal fatal
func (log Logger) Fatal(_ error) {}

// FatalWrap wraps error
func (log Logger) FatalWrap(_ string, _ error) {}

// WithFields preregister fields into logger
func (log Logger) WithFields(_ map[string]interface{}) pkg.Logger {
	return log
}

// SetLevel set level
func (log Logger) SetLevel(_ pkg.LogLevel) {}

// NewLogger new logrus logger
func NewLogger() pkg.Logger {
	return &Logger{}
}
