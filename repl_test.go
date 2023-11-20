package main

import (
	"reflect"
	"testing"
)

func Test_cleanInput(t *testing.T) {

	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "simple lower case input",
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			name:     "mixed case input",
			input:    "hellO World",
			expected: []string{"hello", "world"},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := cleanInput(tt.input); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("cleanInput() = %v, want %v", got, tt.expected)
			}
		})
	}
}
