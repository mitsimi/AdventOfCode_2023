package main

import (
	"testing"
	"util/array"
)

func TestOverlap(t *testing.T) {
	var tests = map[string][]int{
		"two1nine":        {2, 1, 9},
		"eightwothree":    {8, 2, 3},
		"abcone2threexyz": {1, 2, 3},
		"xtwone3four":     {2, 1, 3, 4},
		"4nineightseven2": {4, 9, 8, 7, 2},
		"zoneight":        {1, 8},
		"7pqrstsixteen":   {7, 6},
		"twoneight":       {2, 1, 8},
	}

	for s, n := range tests {
		t.Run(s, func(t *testing.T) {
			str := replaceSpelledNumbers(s)
			res := collectNumbers(str)
			if !array.Equal(n, res) {
				t.Errorf("Want: %v, Got: %v", n, res)
			}
		})
	}
}

func TestExampleInput(t *testing.T) {
	var tests = map[string]int{
		"two1nine":         29,
		"eightwothree":     83,
		"abcone2threexyz":  13,
		"xtwone3four":      24,
		"4nineeightseven2": 42,
		"zoneight":         18,
		"7pqrstsixteen":    76,
		"twoneight":        28,
	}

	for s, n := range tests {
		t.Run(s, func(t *testing.T) {
			newLine := replaceSpelledNumbers(s)
			numbers := collectNumbers(newLine)
			con := constructNumber(numbers)
			if con != n {
				t.Errorf("Want: %d, Got: %d \n lines: %s vs %s", n, con, s, newLine)
			}
		})
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
