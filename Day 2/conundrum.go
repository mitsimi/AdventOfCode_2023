package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"util"
)

func main() {
	path := "input"
	fmt.Printf("Sum: %d\n", cubeConundrum(path))
}

func cubeConundrum(filePath string) int {
	scanner := util.NewScanner(filePath)
	defer scanner.Close()

	var power []int
	for scanner.Scan() {
		var game cubeGame
		var lowest gameRound
	
		lineSplit := strings.Split(scanner.Text(), ":") // Split for {Game \d, ...}
		// Set game id
		game.SetID(extractGameID(lineSplit[0]))

		// Extract rounds
		var rounds = strings.Split(lineSplit[1], ";")
		
		// Process single rounds
		for _, r := range rounds {
			newRound := extractRound(r)
			game.AddRound(newRound)
			
			if newRound.GetRed() > lowest.GetRed() {
				lowest.SetRed(newRound.GetRed())
			}
			
			if newRound.GetGreen() > lowest.GetGreen() {
				lowest.SetGreen(newRound.GetGreen())
			}
			
			if newRound.GetBlue() > lowest.GetBlue() {
				lowest.SetBlue(newRound.GetBlue())
			}
		}
		
		power = append(power, lowest.GetRed() * lowest.GetGreen() * lowest.GetBlue())
	}
	
	return util.Sum(power)
}

func extractGameID(line string) int {
	re, err := regexp.Compile("\\d+")
	util.CheckErr(err)
	id := re.FindString(line)
	num, err := strconv.Atoi(id)
	util.CheckErr(err)
	return num
}

func extractRound(str string) gameRound {
	var round gameRound
	var cubeLine = strings.Split(str, ",") // <count> <color>

	for _, cube := range cubeLine {
		split := strings.Fields(cube)
		count := split[0]
		color := split[1]

		switch color {
		case "red":
			round.SetRed(util.Atoi(count))
		case "green":
			round.SetGreen(util.Atoi(count))
		case "blue":
			round.SetBlue(util.Atoi(count))
		}
	}
	return round
}
