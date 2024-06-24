package main

import (
	"fmt"
	"regexp"
	"strings"
	"util"
)

type Branch struct {
	Right string
	Left  string
}

var reg = regexp.MustCompile(`\w{3}`)

func main() {
	sc := util.NewScanner("input")
	sc.Scan()
	movements := strings.Split(sc.Text(), "")
	var network map[string]Branch = make(map[string]Branch)

	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			continue
		}

		matches := reg.FindAllString(line, 3)
		assert(len(matches) != 3, fmt.Sprintf("count of matches wrong: %d", len(matches)))
		network[matches[0]] = Branch{Left: matches[1], Right: matches[2]}
	}

	var pos string = "AAA"
	var goal string = "ZZZ"
	var steps int
	for i := 0; true; i++ {
		prev := pos
		switch movements[i%len(movements)] {
		case "R":
			pos = network[pos].Right
		case "L":
			pos = network[pos].Left
		default:
			panic("unknown movement")
		}

		steps++
		fmt.Printf("%s --> %s\n", prev, pos)
		if pos == goal {
			break
		}
	}
	assert(pos != goal, "did not reach goal")
	fmt.Println(steps)
}

func assert(b bool, msg string) {
	if b {
		panic(msg)
	}
}
