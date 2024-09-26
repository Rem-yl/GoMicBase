package zlog

import (
	"bytes"
	"testing"
)

func TestInfo(t *testing.T) {
	var logOutput bytes.Buffer

	SetLogOutput(&logOutput)
	Info("this is test.")
	if !bytes.Contains(logOutput.Bytes(), []byte("this is test.")) {
		t.Errorf("Expected log message not found in output")
	}
}

func TestInfof(t *testing.T) {
	var logOutput bytes.Buffer

	SetLogOutput(&logOutput)
	Infof("this is test: %s.", "rem")
	if !bytes.Contains(logOutput.Bytes(), []byte("this is test: rem.")) {
		t.Errorf("Expected log message not found in output")
	}
}
