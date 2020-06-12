package zap

import (
	"github.com/hanaboso/go-log/pkg"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger zap
type Logger struct {
	log  *zap.SugaredLogger
	data map[string]interface{}
}

// Debug debug
func (log *Logger) Debug(message string, args ...interface{}) {
	log.log.With(log.fields()...).Debugf(message, args...)
}

// Info info
func (log *Logger) Info(message string, args ...interface{}) {
	log.log.With(log.fields()...).Infof(message, args...)
}

// Warn warn
func (log *Logger) Warn(message string, args ...interface{}) {
	log.log.With(log.fields()...).Warnf(message, args...)
}

// Error error
func (log *Logger) Error(err error) {
	log.log.With(log.fields()...).Error(err.Error())
}

// ErrorWrap wraps error
func (log *Logger) ErrorWrap(message string, err error) {
	log.log.With(log.fields()...).Errorf("%s, reason: %v", message, err)
}

// Fatal fatal
func (log *Logger) Fatal(err error) {
	log.log.With(log.fields()...).Panicf(err.Error())
}

// FatalWrap wraps error
func (log *Logger) FatalWrap(message string, err error) {
	log.log.With(log.fields()...).Panicf("%s, reason: %v", message, err)
}

// SetLevel set level
func (log *Logger) SetLevel(level pkg.LogLevel) {
	switch level {
	case pkg.DEBUG:
		log.log = newWithLevel(zapcore.DebugLevel)
	case pkg.INFO:
		log.log = newWithLevel(zapcore.InfoLevel)
	case pkg.WARNING:
		log.log = newWithLevel(zapcore.WarnLevel)
	case pkg.ERROR:
		log.log = newWithLevel(zapcore.ErrorLevel)
	case pkg.FATAL:
		log.log = newWithLevel(zapcore.FatalLevel)
	}
}

// WithFields preregister fields into logger
func (log *Logger) WithFields(data map[string]interface{}) pkg.Logger {
	log.data = data

	return log
}

// NewLogger new zap logger
func NewLogger() pkg.Logger {
	return &Logger{
		log:  newWithLevel(zapcore.InfoLevel),
		data: nil,
	}
}

func newWithLevel(level zapcore.Level) *zap.SugaredLogger {
	conf := zap.Config{
		Level:       zap.NewAtomicLevelAt(level),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "json",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}

	log, _ := conf.Build()

	return log.Sugar()
}

func (log *Logger) fields() []interface{} {
	l := make([]interface{}, len(log.data)*2)

	i := 0
	for key, value := range log.data {
		l[i] = key
		l[i+1] = value
		i += 2
	}

	log.data = nil

	return l
}
