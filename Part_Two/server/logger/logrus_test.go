package logger

import (
	"reflect"
	"testing"
)

func TestNewLogger(t *testing.T) {
	tests := []struct {
		name    string
		notWant *Logger
	}{
		{
			"Test new logger eq nil",
			nil,
		}, {
			"Test new logger eq empty Logger",
			&Logger{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLogger(); reflect.DeepEqual(got, tt.notWant) {
				t.Errorf("NewLogger() = %v, want %v", got, tt.notWant)
			}
		})
	}
}
