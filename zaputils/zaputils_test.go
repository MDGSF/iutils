package zaputils

import "testing"

func TestNewZapLog(t *testing.T) {
	logger, sugar := NewZapLog("123", "456")
	if logger == nil {
		t.Errorf("logger == nil")
	}
	if sugar == nil {
		t.Errorf("sugar == nil")
	}
	sugar.Infof("hello zap")
}

func TestNewZapConsole(t *testing.T) {
	sugar := NewZapConsole()
	if sugar == nil {
		t.Errorf("sugar == nil")
	}
	sugar.Infof("hello zap")
}
