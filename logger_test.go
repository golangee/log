//nolint: testpackage
package log

import (
	"testing"
)

func TestNew(t *testing.T) {
	logger := New("default")
	logger.Trace("hello world", Obj("id", 5))
	logger.Info("hello world", Obj("id", 5))
	logger.Debug("hello world", Obj("id", 5))
	With(logger, "sub", Obj("url", "abc")).Error("hello world", Obj("id", 5))
}
