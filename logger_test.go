package main

import (
	"errors"
	"testing"

	"github.com/sirupsen/logrus"
)

func Test_Debug(t *testing.T) {
	logrus.Debug()
	err := errors.New("Debug logger testing")
	Debug("Debugger", err)
	t.Log("Say bye")
}
