package color

import (
	"testing"
)

func TestStyle(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		styles   []int
		expected string
	}{
		{"Empty string", "", []int{1}, "\x1b[1m\x1b[0m"},
		{"Non-empty string", "test", []int{1}, "\x1b[1mtest\x1b[0m"},
		{"Different style", "test", []int{2}, "\x1b[2mtest\x1b[0m"},
		{"Multiple styles", "test", []int{1, 2}, "\x1b[1;2mtest\x1b[0m"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := Style(tt.text, tt.styles...)
			if output != tt.expected {
				t.Errorf("got %v, want %v", output, tt.expected)
			}
		})
	}
}

func TestRed(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected string
	}{
		{"Empty string", "", "\x1b[31m\x1b[0m"},
		{"Non-empty string", "test", "\x1b[31mtest\x1b[0m"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := Red(tt.text)
			if output != tt.expected {
				t.Errorf("got %v, want %v", output, tt.expected)
			}
		})
	}
}
