package main

import (
	"testing"
)

const (
	TestTemplatePath = "../../static/web/template"
)

var (
	TestDataBindings = struct {
		Title          string
		WindowLocation string
	}{
		Title:          "ttyd",
		WindowLocation: "https://192.168.0.200:7681",
	}
)

func Test_render(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Case 1",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := render(TestTemplatePath, TestDataBindings); (err != nil) != tt.wantErr {
				t.Errorf("render() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
