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

	var pos []string
	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			continue
		}

		matches := reg.FindAllString(line, 3)
		assert(len(matches) != 3, fmt.Sprintf("count of matches wrong: %d", len(matches)))
		network[matches[0]] = Branch{Left: matches[1], Right: matches[2]}
		if strings.HasSuffix(matches[0], "A") {
			pos = append(pos, matches[0])
		}
	}

	//for node := range network {
	//	if !strings.HasSuffix(node, start) {
	//		continue
	//	}
	//	steps := 0
	//	for !strings.HasSuffix(node, end) {
	//		next_dir := direction[steps%len(direction)]
	//		if next_dir == 'L' {
	//			node = network[node][0]
	//		} else {
	//			node = network[node][1]
	//		}
	//		steps++
	//	}
	//	results = append(results, steps)
	//}
	//val := results[0]
	//for i := 1; i < len(results); i++ {
	//	val = lcm(val, results[i])
	//}

	results := []int{}
	for i := range pos {
		var steps int = 0
		for !strings.HasSuffix(pos[i], "Z") {
			move := movements[steps%len(movements)]
			switch move {
			case "R":
				pos[i] = network[pos[i]].Right
			case "L":
				pos[i] = network[pos[i]].Left
			default:
				panic("unknown movement")
			}
			steps++
		}
		results = append(results, steps)
	}
	fmt.Println("---------")
	val := results[0]
	for i := 1; i < len(results); i++ {
		val = lcm(val, results[i])
	}
	fmt.Println(val)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func assert(b bool, msg string) {
	if b {
		panic(msg)
	}
}
