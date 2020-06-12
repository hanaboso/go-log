package pkg

// LogLevel log level
type LogLevel int

const (
	// DEBUG debug log level
	DEBUG LogLevel = iota
	// INFO info log level
	INFO
	// WARNING warning log level
	WARNING
	// ERROR error log level
	ERROR
	// FATAL fatal log level
	FATAL
)

// Logger logger interface
type Logger interface {
	Debug(message string, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(err error)
	ErrorWrap(message string, err error)
	Fatal(err error)
	WithFields(data map[string]interface{}) Logger
	SetLevel(level LogLevel)
}
