package main

import (
	"fmt"
	"math"
)

type Race struct {
	time     int
	distance int
}

func main() {
	race1 := Race{62, 644}
	race2 := Race{73, 1023}
	race3 := Race{75, 1240}
	race4 := Race{65, 1023}

	races := [4]Race{race1, race2, race3, race4}

	theRace := Race{62737565, 644102312401023}
	bigRace := [1]Race{theRace}
	wins := float64(1)
	bigWin := float64(1)
	for i := 0; i < len(races); i++ {
		winPossibilities := calculateWinningPossibilities(races[i])
		wins *= winPossibilities
	}

	for i := 0; i < len(bigRace); i++ {
		winPossibilities := calculateWinningPossibilities(bigRace[i])
		bigWin *= winPossibilities
	}

	fmt.Printf("Multi Race: %f\n", wins)
	fmt.Printf("BigRace: %f\n", bigWin)
}

func calculateWinningPossibilities(race Race) float64 {
	raceTimef := float64(race.time)
	raceDistancef := float64(race.distance)
	lowestPossibleValue := math.Floor(raceDistancef / raceTimef)

	for i := lowestPossibleValue; i < raceTimef; i++ {
		if i*(raceTimef-i) > raceDistancef {
			return raceTimef + 1 - (i * 2)
		}
	}

	return 0

}
