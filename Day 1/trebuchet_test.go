package main

import (
	"testing"
)

func TestExampleInput(t *testing.T) {
	var tests = map[string]int{
		"two1nine":         29,
		"eightwothree":     83,
		"abcone2threexyz":  13,
		"xtwone3four":      24,
		"4nineeightseven2": 42,
		"zoneight234":      14,
		"7pqrstsixteen":    76,
	}

	for s, n := range tests {
		newLine := replaceSpelledNumbers(s)
		numbers := collectNumbers(newLine)
		con := constructNumber(numbers)
		if con != n {
			t.Errorf("Want: %d, Got: %d \n lines: %s vs %s", n, con, s, newLine)
		}
	}
}

func TestTrebuchet(t *testing.T) {
	type test struct {
		path string
		sum  int
	}

	var testCase = test{
		path: "example",
		sum:  281,
	}

	result := Trebuchet(testCase.path)

	if result != testCase.sum {
		t.Errorf("Want: %d, Got: %d\n", testCase.sum, result)
	}
}
