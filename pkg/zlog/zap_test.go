package zlog

import (
	"bytes"
	"reflect"
	"testing"

	"go.uber.org/zap/zapcore"
)

func TestSetLogOutput(t *testing.T) {
	var logOutput bytes.Buffer

	SetLogOutput(&logOutput)
	l.Info("this is test.")
	if !bytes.Contains(logOutput.Bytes(), []byte("this is test.")) {
		t.Errorf("Expected log message not found in output")
	}
}

func TestNewDefaultEncoderConfig(t *testing.T) {
	encoder := newDefaultEncoderConfig()
	if reflect.TypeOf(encoder) != reflect.TypeOf(&zapcore.EncoderConfig{}) {
		t.Errorf("newDefaultEncoderConfig() does not return *zapcore.EncoderConfig, got %T", encoder)
	}
}
