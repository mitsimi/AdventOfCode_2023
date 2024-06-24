package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	tests := []struct {
		method   func(string) int
		expected int
		input    string
	}{
		{
			method:   part1,
			expected: 114,
			input:    "example",
		},
		//{
		//	method:   part1,
		//	expected: 999,
		//	input:    "example",
		//},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			assert.Equal(t, test.expected, test.method(test.input))
		})
	}
}
