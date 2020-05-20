//nolint: testpackage
package log

import (
	"testing"
)

func TestNew(t *testing.T) {
	New().Trace("hello world", Obj("id", 5))
	New().Info("hello world", Obj("id", 5))
	New().Debug("hello world", Obj("id", 5))
	New().Error("hello world", Obj("id", 5))
}
