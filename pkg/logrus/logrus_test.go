package logrus

import (
	"errors"
	"github.com/hanaboso/go-log/pkg"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLogger(t *testing.T) {
	logger := NewLogger()
	logger.SetLevel(pkg.FATAL)
	logger.SetLevel(pkg.ERROR)
	logger.SetLevel(pkg.WARNING)
	logger.SetLevel(pkg.INFO)
	logger.SetLevel(pkg.DEBUG)

	data := make(map[string]interface{})
	data["key"] = "val"
	data["key2"] = 22

	logger.WithFields(data)
	logger.Debug("f: %s", "debug")
	logger.Info("f: %s", "info")
	logger.Warn("f: %s", "warning")
	logger.Error(errors.New("error"))
	logger.ErrorWrap("failed testing", errors.New("cause"))

	defer func() {
		if r := recover(); r != nil {
			require.True(t, true)
		} else {
			require.True(t, false)
		}
	}()
	logger.Fatal(errors.New("fatal"))
}
