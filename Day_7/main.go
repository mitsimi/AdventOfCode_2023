package main

import (
	"fmt"
	"sort"
	"strings"
	"util"
)

type Set struct {
	Hand [5]Card
	Type HandType
	Bid  int
}

var game []Set

func main() {
	lines := util.ReadLines("input")
	// Collect all hands in the game
	for _, line := range lines {
		str := strings.Split(line, " ")
		cards := [5]Card(getCardArr([]rune(str[0])))
		handType := getHandType(cards)
		game = append(game, Set{
			Hand: cards,
			Type: handType,
			Bid:  util.Atoi(str[1]),
		})
	}

	// Sort array based on winning hierachy
	sort.SliceStable(game, func(i, j int) bool {
		if game[i].Type == game[j].Type {
			for idx, iCard := range game[i].Hand {
				jCard := game[j].Hand[idx]

				if iCard == jCard {
					continue
				}

				return iCard < jCard
			}
		}

		return game[i].Type < game[j].Type
	})

	var sum int = 0
	for i, set := range game {
		sum += set.Bid * (i + 1)
	}
	fmt.Println(sum)
	//printGame(game)
}

func getCardArr(arr []rune) []Card {
	var cards []Card
	for _, r := range arr {
		cards = append(cards, getCard(r))
	}
	return cards
}

func printGame(game []Set) {
	fmt.Println()
	for _, set := range game {
		for _, c := range set.Hand {
			fmt.Print(c.String())
		}
		fmt.Printf("\t%d\t%s\n", set.Bid, set.Type.String())

		fmt.Println("-----------------------")
	}
}
