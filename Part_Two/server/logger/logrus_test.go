package logger

import (
	"reflect"
	"testing"
)

func TestNewLogger(t *testing.T) {
	tests := []struct {
		name string
		want *Logger
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLogger(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}
