package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"
	"util"
)

type Mapping struct {
	Dest int
	Src  int
	Len  int
}

var parseNums = regexp.MustCompile("[0-9]+")
var seeds []int
var ranges []Mapping

func main() {
	file, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}
	content := strings.Split(string(file), "\n\n")

	nums := util.Atoi_Arr(parseNums.FindAllString(strings.Join(content[:1], ""), -1))
	for i := 0; i < len(nums); i += 2 {
		for offset := range nums[i+1] {
			seeds = append(seeds, nums[i]+offset)
		}
	}
	fmt.Println("parsed seeds")

	for i, block := range content[1:] {
		for _, line := range strings.Split(block, "\n")[1:] {
			if line == "" {
				break
			}

			nums := util.Atoi_Arr(parseNums.FindAllString(line, -1))
			ranges = append(ranges, Mapping{Dest: nums[0], Src: nums[1], Len: nums[2]})
		}

		new := make([]int, len(seeds))
		for i, seed := range seeds {
			found := false
			for _, r := range ranges {
				if r.Src <= seed && seed < (r.Src+r.Len) {
					new[i] = seed - r.Src + r.Dest
					found = true
					break
				}
			}

			if !found {
				new[i] = seed
			}

		}
		seeds = new
		ranges = []Mapping{}

		fmt.Printf("Blocks done: %d / %d \n", i+1, len(content[1:]))
	}

	util.Sout(slices.Min(seeds))
}
