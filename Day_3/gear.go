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
		for j, r := range row {
			if r == '*' {
				if ratio := searchNumbers(i, j); ratio > 0 {
					addToCollection(ratio)
				}
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

func searchNumbers(row, col int) (ratio int) {
	ratio = 1
	var found int

	// Starting top right
	row--
	col--

	// Setting boundaries
	rowBound := row + 3
	colBound := col + 3

	// staying in lower bounds
	if row < 0 {
		row++
	}
	if col < 0 {
		col++
	}

	// Looking for numbers around
	for i := row; i < rowBound; i++ {
		if i > len(table)-1 {
			break
		}

		var frags []rune
		for j := col; j < colBound; j++ {
			if j > len(table[i])-1 {
				break
			}

			if unicode.IsNumber(table[i][j]) {
				// Going to the beginning of the Number
				for ; unicode.IsNumber(table[i][j]); j-- {
					if j == 0 {
						break
					}
				}

				// Make sure it points to the beginning of the number
				// If current but next is a number
				if !unicode.IsNumber(table[i][j]) && unicode.IsNumber(table[i][j+1]) {
					j++
				}

				// Reading the number
				for ; unicode.IsNumber(table[i][j]); j++ {
					frags = append(frags, table[i][j])

					if j >= len(table[i])-1 {
						break
					}
				}

				ratio *= constructNumber(frags)
				fmt.Println(string(frags))
				frags = []rune{}
				found++
			}
		}
	}

	if found == 2 {
		return ratio
	} else {
		return 0
	}
}
