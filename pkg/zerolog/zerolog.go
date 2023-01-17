package zerolog

import (
	"fmt"
	"github.com/hanaboso/go-log/pkg"
	"github.com/hanaboso/go-utils/pkg/intx"
	"github.com/hanaboso/go-utils/pkg/stringx"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io"
	"runtime/debug"
	"strings"
)

type Logger struct {
	data map[string]interface{}
}

func (this Logger) Debug(message string, args ...interface{}) {
	log.Debug().Fields(this.fields()).Msg(fmt.Sprintf(message, args...))
}

func (this Logger) Info(message string, args ...interface{}) {
	log.Info().Fields(this.fields()).Msg(fmt.Sprintf(message, args...))
}

func (this Logger) Warn(message string, args ...interface{}) {
	log.Warn().Fields(this.fields()).Msg(fmt.Sprintf(message, args...))
}

func (this Logger) Error(err error) {
	log.Error().Fields(this.fields()).Err(err).Send()
}

func (this Logger) ErrorWrap(message string, err error) {
	log.Error().Fields(this.fields()).Err(fmt.Errorf("%s, reason: %v", message, err)).Send()
}

func (this Logger) Fatal(err error) {
	log.Fatal().Fields(this.fields()).Err(err).Send()
}

func (this Logger) FatalWrap(message string, err error) {
	log.Fatal().Fields(this.fields()).Err(fmt.Errorf("%s, reason: %v", message, err)).Send()
}

func (this Logger) SetLevel(level pkg.LogLevel) {
	switch level {
	case pkg.DEBUG:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		zerolog.ErrorStackMarshaler = func(err error) interface{} {
			return ParseTrace(debug.Stack())
		}
	case pkg.INFO:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case pkg.WARNING:
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case pkg.ERROR:
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case pkg.FATAL:
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	}
}

func (this Logger) WithFields(data map[string]interface{}) pkg.Logger {
	return Logger{
		data: data,
	}
}

func NewLogger(sender io.Writer) pkg.Logger {
	log.Logger = zerolog.
		New(sender).
		With().
		Timestamp().
		Stack().
		Logger()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.TimestampFieldName = "timestamp"
	zerolog.ErrorFieldName = "message"
	zerolog.ErrorStackFieldName = "trace"

	return Logger{
		data: nil,
	}
}

func (this *Logger) fields() []interface{} {
	l := make([]interface{}, len(this.data)*2)

	i := 0
	for key, value := range this.data {
		l[i] = key
		l[i+1] = value
		i += 2
	}

	this.data = nil

	return l
}

func ParseTrace(trace []byte) []interface{} {
	type frame struct {
		Function string `json:"function"`
		File     string `json:"file"`
	}

	stack := make([]interface{}, 0)
	data := string(trace)

	index := strings.Index(data, ":")
	if index < 0 {
		return nil
	}
	data = data[index+2:]

	lines := strings.Split(data, "\n")
	_ = lines

	limit := intx.Min(len(lines)-1, 16)
	for i := 6; i < limit; i += 2 {
		if strings.Contains(lines[i+1], "/rs/zerolog@") {
			continue
		}

		fns := strings.Split(strings.TrimLeft(lines[i], "\t"), "/")
		fn := fns[len(fns)-1]
		index = strings.LastIndex(fn, "(")
		if index < 0 {
			continue
		}
		fn = fn[:index]

		fileLine := strings.TrimLeft(lines[i+1], "\t")

		stack = append(stack, frame{
			File:     stringx.ToChar(fileLine, " "),
			Function: fn,
		})
	}

	return stack
}
