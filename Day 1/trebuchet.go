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
	for i := 0; i < 1; i++ {
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
	"one":   "o1e",
	"two":   "t2o",
	"three": "t3e",
	"four":  "f4r",
	"five":  "f5e",
	"six":   "s6x",
	"seven": "s7n",
	"eight": "e8t",
	"nine":  "n9e",
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
	var numbers []int
	var cache string

	for _, r := range str {
		cache = cache + string(r)
		if unicode.IsDigit(r) {
			numbers = append(numbers, getNumber(r))
		}
	}
	return numbers
}

func constructNumber(arr []int) (number int) {
	if len(arr) == 1 { // one item
		number = arr[0]*10 + arr[0]
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
