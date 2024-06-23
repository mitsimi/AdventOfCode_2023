package main

import (
	"fmt"
	"regexp"
	"strings"
	"util"
)

type race struct {
	Time  int
	Dists []int
}

var parseNums = regexp.MustCompile("[0-9]+")
var tries []race

func main() {
	sc := util.NewScanner("input")
	var nums [][]int
	for sc.Scan() {
		line := strings.ReplaceAll(sc.Text(), " ", "")
		nums = append(nums, util.Atoi_Arr(parseNums.FindAllString(line, -1)))
	}
	sc.Close()

	// Generate possibilities
	for _, time := range nums[0] {
		try := race{Time: time, Dists: make([]int, time+1)}
		for i := range time + 1 {
			try.Dists[i] = i * (time - i)
		}
		tries = append(tries, try)
	}

	// Filter wins
	var result int = 1
	for i, record := range nums[1] {
		var count int
		for _, dist := range tries[i].Dists {
			if dist > record {
				count++
			}
		}
		result *= count
	}

	fmt.Println(result)
}
