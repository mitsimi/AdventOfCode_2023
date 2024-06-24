package main

import (
	"fmt"
	"slices"
	"strings"
	"util"
)

func main() {
	file := "input"
	fmt.Println("Part 1:", part1(file))
	fmt.Println("Part 2:", part2(file))
}

func part1(input string) int {
	var result int = 0
	for _, line := range util.ReadLines(input) {
		var seq []int = []int{}
		for _, n := range strings.Split(line, " ") {
			seq = append(seq, util.Atoi(n))
		}

		result += getNext(seq)
	}
	return result
}

func part2(input string) int {
	var result int = 0
	for _, line := range util.ReadLines(input) {
		var seq []int = []int{}
		for _, n := range strings.Split(line, " ") {
			seq = append(seq, util.Atoi(n))
		}
		slices.Reverse(seq)
		result += getNext(seq)
	}
	return result
}

func getNext(seq []int) int {
	vars := []int{}
	for i := range len(seq) - 1 {
		vars = append(vars, seq[i+1]-seq[i])
	}
	if !slices.ContainsFunc(vars, func(x int) bool { return x != 0 }) {
		return seq[len(seq)-1]
	}

	return seq[len(seq)-1] + getNext(vars)
}
