package main

import (
	"slices"
	"sort"
)

type HandType int

const (
	High   HandType = iota // 0
	Pair                   // 1
	DPair                  // 2
	Threes                 // 3
	Full                   // 4
	Fours                  // 5
	Fives                  // 6
)

func (t HandType) Compare(s HandType) int {
	return int(t) - int(s)
}

func (t HandType) String() string {
	switch t {
	case High:
		return "High"
	case Pair:
		return "Pair"
	case DPair:
		return "DPair"
	case Threes:
		return "Threes"
	case Full:
		return "Full"
	case Fours:
		return "Fours"
	case Fives:
		return "Fives"
	}

	return "UNDEFINED"
}

func getHandType(cards [5]Card) HandType {
	var count map[Card]int = map[Card]int{}

	for _, card := range cards {
		if _, ok := count[card]; ok {
			count[card] = count[card] + 1
			continue
		}
		count[card] = 1
	}

	keys := sortMap(count)
	if slices.Contains(keys, Joker) {
		if keys[0] != Joker {
			count[keys[0]] = count[keys[0]] + count[Joker]
			count[Joker] = 0
		} else if keys[0] == Joker && len(keys) > 1 {
			count[keys[1]] = count[keys[1]] + count[Joker]
			count[Joker] = 0
		}
	}
	keys = sortMap(count)

	switch count[keys[0]] {
	case 5:
		return Fives
	case 4:
		return Fours
	case 3:
		if count[keys[1]] == 2 {
			return Full
		}
		return Threes
	case 2:
		if count[keys[1]] == 2 {
			return DPair
		}
		return Pair
	case 1:
		return High
	}

	return 1 << 32
}

func sortMap(count map[Card]int) []Card {
	keys := make([]Card, 0, len(count))
	for key := range count {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return count[keys[i]] > count[keys[j]]
	})
	return keys
}
