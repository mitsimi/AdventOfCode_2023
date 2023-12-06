package main

import (
	"fmt"
	"log"
	"regexp"
	"slices"
	"strings"
	"time"
	"util"
)

func main() {
	sc := util.NewScanner("input")
	sc.Scan()
	// parse all numbers of the first line
	seedIds := parseNums.FindAllString(sc.Text(), -1)
	log.Println("parsing seeds")
	for _, id := range seedIds {
		seeds = append(seeds, util.Atoi(id))
	}

	// parse map content
	log.Println("started parsing mappings")
	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			continue
		}

		if slices.Contains(keyword, line) {
			key := line
			log.Printf("parse mappings for %s \n", key)
			for sc.Scan() {
				line := sc.Text()
				log.Printf("parsing line: %s\n", line)
				if line == "" {
					break
				}
				parseMap(key, line)
			}
		}
	}
	log.Println("finished parsing mappings")

	lowest := findLowestLocation()
	fmt.Println(lowest)
}

var parseNums = regexp.MustCompile("[0-9]+")

var keyword = []string{
	"seed-to-soil map:",
	"soil-to-fertilizer map:",
	"fertilizer-to-water map:",
	"water-to-light map:",
	"light-to-temperature map:",
	"temperature-to-humidity map:",
	"humidity-to-location map:",
}

var keywordToMap = map[string]map[int]int{
	"seed-to-soil map:":            seedToSoil,
	"soil-to-fertilizer map:":      soilToFertilizer,
	"fertilizer-to-water map:":     fertilizerToWater,
	"water-to-light map:":          waterToLight,
	"light-to-temperature map:":    lightToTemperature,
	"temperature-to-humidity map:": temperatureToHumidity,
	"humidity-to-location map:":    humidityToLocation,
}

var seeds = []int{}

var seedToSoil = make(map[int]int)
var soilToFertilizer = make(map[int]int)
var fertilizerToWater = make(map[int]int)
var waterToLight = make(map[int]int)
var lightToTemperature = make(map[int]int)
var temperatureToHumidity = make(map[int]int)
var humidityToLocation = make(map[int]int)

func parseMap(keyword, line string) {
	fields := strings.Fields(line) // <dest range start> <source range start> <length>
	dest := util.Atoi(fields[0])
	src := util.Atoi(fields[1])
	length := util.Atoi(fields[2])

	convMap := keywordToMap[keyword]
	log.Println("create mappings")
	t := time.Now()
	for i := 0; i < length; i++ {
		convMap[src+i] = dest + i
	}
	log.Printf("finished mappings within %d ms\n", time.Since(t).Milliseconds())
}

func findLowestLocation() int {
	var locs []int
	log.Println("looking for lowest location")
	for _, seed := range seeds {

		soil := seedToSoil[seed]
		if soil == 0 {
			soil = seed
		}

		fert := soilToFertilizer[soil]
		if fert == 0 {
			fert = soil
		}

		water := fertilizerToWater[fert]
		if water == 0 {
			water = fert
		}

		light := waterToLight[water]
		if light == 0 {
			light = water
		}

		temp := lightToTemperature[light]
		if temp == 0 {
			temp = light
		}

		hum := temperatureToHumidity[temp]
		if hum == 0 {
			hum = temp
		}

		loc := humidityToLocation[hum]
		if loc == 0 {
			loc = hum
		}

		locs = append(locs, loc)
	}

	return slices.Min(locs)
}
