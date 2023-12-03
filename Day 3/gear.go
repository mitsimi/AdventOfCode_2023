package main

import (
	"fmt"
	"unicode"
	"util"
)

func main() {
	gearSchematic()
}

var table [][]rune

func gearSchematic() {
	path := "input"
	table = buildArray(util.NewScanner(path))
	nums := collectNumbers()
	fmt.Println(nums)
	fmt.Println(util.Sum(nums))
}

func buildArray(scanner *util.FileScanner) [][]rune {
	defer scanner.Close()

	var schematic [][]rune
	for scanner.Scan() {
		var line []rune
		for _, v := range scanner.Text() {
			line = append(line, v)
		}
		schematic = append(schematic, line)
	}

	return schematic
}

func collectNumbers() []int {
	var numCollection []int
	addToCollection := func(num int) {
		if num == 0 {
			return
		}
		numCollection = append(numCollection, num)
	}

	for i, row := range table {
		var numFrags []rune
		var symbolFound = false
		for j, r := range row {
			if unicode.IsNumber(r) {
				numFrags = append(numFrags, r)
				if !symbolFound {
					symbolFound = lookupTable(i, j)
				}
			} else {
				if symbolFound {
					addToCollection(constructNumber(numFrags))
					symbolFound = false
				}
				numFrags = []rune{}
			}

			if j == len(row)-1 && symbolFound {
				addToCollection(constructNumber(numFrags))
				symbolFound = false
			}
		}
	}
	return numCollection
}

func constructNumber(arr []rune) int {
	var num = 0
	for _, n := range arr {
		num = num*10 + util.Atoi(string(n))
	}
	return num
}

func lookupTable(row, col int) bool {
	// We start top right
	row--
	col--

	// Set boundaries
	rowBound := row + 3
	columnBound := col + 3

	// Preventing out of bounds with -1
	if row < 0 {
		row++
	}
	if col < 0 {
		col++
	}

	// Go through radius of 1
	for i := row; i < rowBound; i++ {
		if i > len(table)-1 {
			break
		}
		for j := col; j < columnBound; j++ {
			if j > len(table[i])-1 {
				break
			}

			char := table[i][j]
			if !unicode.IsNumber(char) && char != '.' {
				return true
			}
		}
	}
	return false
}
