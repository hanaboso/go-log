package logrus

import (
	"fmt"
	"github.com/hanaboso/go-log/pkg"
	"github.com/sirupsen/logrus"
	"runtime"
)

// Logger logrus
type Logger struct {
	log  *logrus.Logger
	data map[string]interface{}
}

// Debug debug
func (log *Logger) Debug(message string, args ...interface{}) {
	log.log.WithFields(log.data).Debugf(message, args...)
	log.data = nil
}

// Info info
func (log *Logger) Info(message string, args ...interface{}) {
	log.log.WithFields(log.data).Infof(message, args...)
	log.data = nil
}

// Warn warn
func (log *Logger) Warn(message string, args ...interface{}) {
	log.log.WithFields(log.data).Warnf(message, args...)
	log.data = nil
}

// Error error
func (log *Logger) Error(err error) {
	_, file, line, _ := runtime.Caller(1)
	_, file2, line2, _ := runtime.Caller(2)
	if log.data == nil {
		log.data = make(map[string]interface{})
	}
	log.data["stacktrace"] = fmt.Sprintf("%s: %d\n%s: %d", file, line, file2, line2)

	log.log.WithFields(log.data).Error(err.Error())
	log.data = nil
}

// ErrorWrap wraps error
func (log *Logger) ErrorWrap(message string, err error) {
	log.log.WithFields(log.data).Errorf("%s, reason: %v", message, err)
	log.data = nil
}

// Fatal fatal
func (log *Logger) Fatal(err error) {
	_, file, line, _ := runtime.Caller(1)
	if log.data == nil {
		log.data = make(map[string]interface{})
	}
	log.data["stacktrace"] = fmt.Sprintf("%s: %d", file, line)

	log.log.WithFields(log.data).Fatalf(err.Error())
	log.data = nil
}

// WithFields preregister fields into logger
func (log *Logger) WithFields(data map[string]interface{}) pkg.Logger {
	log.data = data

	return log
}

// SetLevel set level
func (log Logger) SetLevel(level pkg.LogLevel) {
	switch level {
	case pkg.DEBUG:
		log.log.SetLevel(logrus.DebugLevel)
	case pkg.INFO:
		log.log.SetLevel(logrus.InfoLevel)
	case pkg.WARNING:
		log.log.SetLevel(logrus.WarnLevel)
	case pkg.ERROR:
		log.log.SetLevel(logrus.ErrorLevel)
	case pkg.FATAL:
		log.log.SetLevel(logrus.FatalLevel)
	}
}

// NewLogger new logrus logger
func NewLogger() pkg.Logger {
	log := logrus.New()
	log.ExitFunc = func(code int) {
		panic(code)
	}

	return &Logger{
		log:  log,
		data: nil,
	}
}
