package config

import (
	"reflect"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		want    *Config
		wantErr bool
	}{
		{
			"Test load config okay",
			"../../config.yaml",
			&Config{Port: ":8080"},
			false,
		}, {
			"Test load config random path",
			"ascvgd.json",
			nil,
			true,
		}, {
			"Test load config error file",
			"../../docker-compose.yaml",
			&Config{},
			false,
		}, {
			"Test load config error file",
			"../../Dockerfile",
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadConfig(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewConfig(t *testing.T) {
	tests := []struct {
		name string
		want *Config
	}{
		{
			name: "Test create config",
			want: &Config{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
