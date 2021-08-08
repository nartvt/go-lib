package golib

import (
	"errors"
	"testing"
)

func Test_Debug(t *testing.T) {
	err := errors.New("Debug logger testing")
	Debug("Debugger", err)
	t.Log("Say bye")
}
