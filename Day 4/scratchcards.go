package main

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"util"
)

func main() {
	path := "input"
	cs := util.NewScanner(path)
	cards := collect(cs)
	sum := sumCopies(cards)
	fmt.Println(sum)
}

type scratchcard struct {
	id      int
	winNums []int
	myNums  []int
	wins    int
}

func collect(fs *util.FileScanner) []scratchcard {
	var cards []scratchcard
	for fs.Scan() {
		var card scratchcard
		line := fs.Text()
		lineSplit := strings.Split(line, ":") // Split for {Card \d, ...}

		// Set card id
		card.id = extractCardId(lineSplit[0])

		// Split number types
		// [0] winning numbers
		// [1] "rolled" numbers
		numbers := strings.Split(lineSplit[1], "|")
		card.winNums = extractNumbers(numbers[0])
		card.myNums = extractNumbers(numbers[1])

		card.wins = countWinning(card)

		cards = append(cards, card)
	}

	return cards
}

func sumCopies(cards []scratchcard) int {
	var cardCount = map[int]int{}
	for i, card := range cards {
		// Every card counts towards the total
		cardCount[i]++

		// At least one winning number
		for j := 1; j <= card.wins; j++ {
			cardCount[i+j] += cardCount[i]
		}

	}

	var sum int
	for _, value := range cardCount {
		sum += value
	}
	return sum
}

func countWinning(card scratchcard) (count int) {
	for _, mine := range card.myNums {
		if slices.Contains(card.winNums, mine) {
			// Count up for every winning number
			count++
		}
	}
	return count
}

func extractCardId(line string) int {
	re, err := regexp.Compile("\\d+")
	util.CheckErr(err)
	id := re.FindString(line)
	num, err := strconv.Atoi(id)
	util.CheckErr(err)
	return num
}

func extractNumbers(line string) []int {
	var nums []int
	for _, n := range strings.Fields(line) {
		nums = append(nums, util.Atoi(n))
	}
	return nums
}
