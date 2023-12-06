package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var mapTypes []string = []string{"seed-to-soil", "soil-to-fertilizer", "fertilizer-to-water",
	"water-to-light", "light-to-temperature", "temperature-to-humidity", "humidity-to-location"}

type AlmanacMap struct {
	mapType     string
	source      int
	destination int
}

type Map struct {
	source      int
	destination int
	range_      int
}

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	var seedBaselinesPart1 map[int]int
	var seedBaselinesPart2 map[int]int
	var mapping []string
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(seedBaselinesPart1) == 0 || len(seedBaselinesPart2) == 0 {
			seedBaselinesPart1 = readSeeds(line)
			seedBaselinesPart2 = readSeedsRange(line)
		}
		mapping = append(mapping, line)
	}
	var mapOfAlmanacs map[string][]AlmanacMap = make(map[string][]AlmanacMap)
	for _, mapType := range mapTypes {
		seedBaselinesPart1 = buildMapping(mapping, mapType, seedBaselinesPart1, mapOfAlmanacs)
		seedBaselinesPart2 = buildMapping(mapping, mapType, seedBaselinesPart2, mapOfAlmanacs)
	}
	location := findLowestLocation(mapOfAlmanacs)
	fmt.Println(location)
}

func readSeedsRange(line string) map[int]int {
	var seeds map[int]int = make(map[int]int)
	if strings.Contains(line, "seeds") {
		line = regexp.MustCompile(`[^0-9 ]+`).ReplaceAllString(line, "")

		values := strings.Split(strings.TrimSpace(line), " ")

		for i := 0; i < len(values); i += 2 {
			seed, err := strconv.Atoi(values[i])
			range_, err1 := strconv.Atoi(values[i+1])

			if err != nil || err1 != nil {
				fmt.Errorf("Can't convert this to an int!")
			} else {
				seeds[seed] = range_
			}
		}

	}
	return seeds
}

func readSeeds(line string) map[int]int {
	var seeds map[int]int = make(map[int]int)
	if strings.Contains(line, "seeds") {
		line = regexp.MustCompile(`[^0-9 ]+`).ReplaceAllString(line, "")
		values := strings.Split(line, " ")

		for _, i := range values {
			j, err := strconv.Atoi(i)
			if err != nil {
				fmt.Errorf("Can't convert this to an int!" + i)
			} else {
				seeds[j] = 0
			}

		}

	}
	return seeds
}

func buildMapping(entries []string, mapType string, baselines map[int]int, mapOfAlmanacs map[string][]AlmanacMap) map[int]int {
	found := false
	var newBaseline map[int]int = make(map[int]int)
	for i := 0; i < len(entries); i++ {
		if strings.Contains(entries[i], mapType) {
			found = true
			continue
		}
		if found && (strings.TrimSpace(entries[i]) != "" && !strings.Contains(entries[i], "map")) {
			values := strings.Split(entries[i], " ")
			dest, _ := strconv.Atoi(values[0])
			source, _ := strconv.Atoi(values[1])
			range_, _ := strconv.Atoi(values[2])
			mapEntry := Map{source, dest, range_}
			createAlmanacs(mapEntry, baselines, mapOfAlmanacs, mapType, newBaseline)
		}
		if found && (strings.TrimSpace(entries[i]) == "" || strings.Contains(entries[i], "map") || i == len(entries)-1) {
			for source, range_ := range baselines {
				for j := 0; j <= range_; j++ {
					baseline := source + j

					almanacMap := new(AlmanacMap)
					almanacMap.mapType = mapType

					almanacMap.source = baseline
					almanacMap.destination = baseline
					prepareNewBaseline(newBaseline, almanacMap.destination)
					mapOfAlmanacs[mapType] = append(mapOfAlmanacs[mapType], *almanacMap)
				}
			}
			break
		}
	}
	return newBaseline
}

func prepareNewBaseline(baselines map[int]int, baseline int) {
	rangeIncreased := false
	for source, range_ := range baselines {
		if baseline-source+range_ == 1 || baseline-source+range_ == 0 {
			baselines[source] = range_ + 1
			rangeIncreased = true
			break
		}
	}
	if !rangeIncreased {
		baselines[baseline] = 0
	}
}

func createAlmanacs(entry Map, baselines map[int]int, mapOfAlmanacs map[string][]AlmanacMap, mapType string, newBaseline map[int]int) {

	for source, range_ := range baselines {
		valueConsumed := -1
		rangeConsumed := 0
		for j := 0; j < range_; j++ {
			baseline := source + j
			if entry.source <= baseline && baseline <= entry.source+entry.range_ {
				if valueConsumed == -1 {
					valueConsumed = baseline

				}
				rangeConsumed++

				almanacMap := new(AlmanacMap)
				almanacMap.mapType = mapType

				almanacMap.source = baseline
				almanacMap.destination = entry.destination + (baseline - entry.source)
				prepareNewBaseline(newBaseline, almanacMap.destination)
				mapOfAlmanacs[mapType] = append(mapOfAlmanacs[mapType], *almanacMap)
			}
		}
		if valueConsumed != -1 {
			addToNewBaseline(baselines, source, valueConsumed, rangeConsumed)
		}
	}
}

func addToNewBaseline(baselines map[int]int, baseline int, valueConsumed int, rangeConsumed int) {
	if valueConsumed == baseline {
		if rangeConsumed == baselines[baseline] {
			baselines[baseline] = -1
			return
		} else {
			baselines[baseline+rangeConsumed] = baselines[baseline] - rangeConsumed
		}
	} else {
		baselines[baseline] = valueConsumed - baselines[baseline]
		createRemainingBaselineIfNeeded(baselines, baseline, valueConsumed, rangeConsumed)
	}
}

func createRemainingBaselineIfNeeded(baselines map[int]int, baseline int, valueConsumed int, rangeConsumed int) {
	if valueConsumed-baseline+rangeConsumed < baselines[baseline] {
		baselines[valueConsumed+rangeConsumed+1] = baselines[baseline] - (valueConsumed - baseline + rangeConsumed)
	}
}

func findLowestLocation(mapOfAlmanacs map[string][]AlmanacMap) int {
	almanacs := mapOfAlmanacs["humidity-to-location"]
	lowestLocation := -1
	for _, almanac := range almanacs {
		if lowestLocation == -1 || almanac.destination < lowestLocation {
			lowestLocation = almanac.destination
		}
	}
	return lowestLocation
}
