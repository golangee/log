//nolint: testpackage
package log

import (
	"testing"
)

func TestNew(t *testing.T) {

	logger := New("default")
	logger.Info(Log("my.logger"), Trace(), Msg("hello"))
	logger.Info(Log("my.logger"), Dbg(), Msg("hello"))
	logger.Info(Log("my.logger"), Info(), Msg("hello"))
	logger.Info(Log("my.logger"), Warn(), Msg("hello"))
	logger.Info(Log("my.logger"), Error(), Msg("hello"), ErrStack())
	/*	logger.Info(Msg("hello world"), Obj("id", 5))
		logger.Info("hello world", Obj("id", 5))
		logger.Debug("hello world", Obj("id", 5))
		logger.Warn("hello world", Obj("id", 5))
		With(logger, "sub", Obj("url", "abc")).Error("hello world", Obj("id", 5))*/
}
