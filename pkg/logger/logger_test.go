package logger

import (
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestInitializeZapGlobals_InfoLevel(t *testing.T) {
	initializeZapGlobals("info")

	if zap.L().Core().Enabled(zapcore.DebugLevel) {
		t.Error("debug level should be disabled when LOG_LEVEL is info")
	}
	if !zap.L().Core().Enabled(zapcore.InfoLevel) {
		t.Error("info level should be enabled")
	}
}

func TestInitializeZapGlobals_DebugLevel(t *testing.T) {
	initializeZapGlobals("debug")

	if !zap.L().Core().Enabled(zapcore.DebugLevel) {
		t.Error("debug level should be enabled")
	}
}
