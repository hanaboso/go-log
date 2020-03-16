package null

import (
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
	logger.Fatal(nil)
}
