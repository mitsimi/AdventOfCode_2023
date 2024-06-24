#!/bin/bash
source .env
echo -n "Please number of the day: "
read number

dir=Day_$number
mkdir $dir

cd $dir

curl --cookie "session=$AOC_COOKIE" https://adventofcode.com/2023/day/$number/input > input
touch main.go
echo "package main" > main.go
echo "
func main() {
    file := \"input\"
    fmt.Println(\"Part 1:\", part1(file))
    fmt.Println(\"Part 2:\", part2(file))
}

func part1(input string) int {
	panic(\"not implemented\")
}

func part2(input string) int {
	panic(\"not implemented\")
}
" >> main.go

touch example
touch main_test.go
echo "package main" > main_test.go
echo "
import (
	\"testing\"
	\"github.com/stretchr/testify/assert\"
)

func TestMain(t *testing.T) {
tests := []struct {
		name     string
		method   func(string) int
		expected int
		input    string
	}{
		{
			name:     \"Part 1\",
			method:   part1,
			expected: 999,
			input:    "example",
		},
		//{
		//	name:     \"Part 2\",
		//	method:   part1,
		//	expected: 999,
		//	input:    example,
		//},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.method(test.input))
		})
	}
}
" >> main_test.go

go mod init $dir
cd ..
go work use $dir