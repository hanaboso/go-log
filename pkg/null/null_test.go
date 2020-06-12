package null

import (
	"errors"
	"github.com/hanaboso/go-log/pkg"
	"testing"
)

func TestLogger(t *testing.T) {
	logger := NewLogger()
	logger.SetLevel(pkg.WARNING)
	logger.WithFields(nil).Info("")
	logger.Debug("")
	logger.Warn("")
	logger.Error(nil)
	logger.ErrorWrap("", errors.New(""))
	logger.Fatal(nil)
}
