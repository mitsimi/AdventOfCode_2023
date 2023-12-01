package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(Trebuchet("input"))
	}
}

func Trebuchet(path string) int {
	file, err := os.Open(path)
	checkErr(err)
	defer func(file *os.File) {
		err := file.Close()
		checkErr(err)
	}(file)

	scanner := bufio.NewScanner(file)
	return sumScan(scanner)
}

var spelledMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func sumScan(scanner *bufio.Scanner) int {

	var numArr []int
	for scanner.Scan() {
		var line = scanner.Text()

		line = replaceSpelledNumbers(line)

		cache := collectNumbers(line)
		numArr = append(numArr, constructNumber(cache))
	}

	return arraySum(numArr)
}

func replaceSpelledNumbers(line string) string {
	var cache string
	for _, r := range line {
		cache += string(r)
		for s, d := range spelledMap {
			cache = strings.ReplaceAll(cache, s, d)
		}
	}

	return cache
}

func collectNumbers(str string) []int {
	var cache []int
	for _, r := range str {
		if unicode.IsDigit(r) {
			cache = append(cache, getNumber(r))
		}
	}
	return cache
}

func constructNumber(arr []int) (number int) {
	if len(arr) == 1 { // one item
		number = arr[0]*10 + number
	} else if len(arr) > 1 { // more than one item
		number = arr[0]*10 + arr[len(arr)-1]
	} else { // no item
		number = 0
	}
	return
}

func arraySum(arr []int) (sum int) {
	for _, n := range arr {
		sum = sum + n
	}
	return
}

func getNumber(r rune) int {
	number, err := strconv.Atoi(string(r))
	checkErr(err)
	return number
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
