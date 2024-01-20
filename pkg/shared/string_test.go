package shared

import (
	"testing"
)

func TestUpperCamelCaseToLowerCameCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Empty string", "", ""},
		{"Single uppercase letter", "A", "a"},
		{"Single lowercase letter", "a", "a"},
		{"Uppercase camel case", "HelloWorld", "helloWorld"},
		{"Lowercase camel case", "helloWorld", "helloWorld"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := UpperCamelCaseToLowerCameCase(tt.input)
			if output != tt.expected {
				t.Errorf("got %v, want %v", output, tt.expected)
			}
		})
	}
}
