package main

import (
	"fmt"
	"testing"
)

func Test_findNumberOfWinsMultipleRaces(t *testing.T) {
	race1 := Race{7, 9}
	race2 := Race{15, 40}
	race3 := Race{30, 200}

	races := [3]Race{race1, race2, race3}
	wins := float64(1)
	for i := 0; i < len(races); i++ {
		winPossibilities := calculateWinningPossibilities(races[i])
		wins *= winPossibilities
	}
	if wins != 288 {
		t.Errorf("Result was incorrect, got: %f, want: %d.", wins, 288)
	}
}

func Test_findNumberOfWinsInABigRace(t *testing.T) {
	race1 := Race{71530, 940200}

	races := [1]Race{race1}
	wins := float64(1)
	for i := 0; i < len(races); i++ {
		winPossibilities := calculateWinningPossibilities(races[i])
		wins *= winPossibilities
	}
	if wins != 71503 {
		t.Errorf("Result was incorrect, got: %f, want: %d.", wins, 71503)
	}
}

func Test_findNumberOfWins(t *testing.T) {
	race1 := Race{7, 9}
	race2 := Race{15, 40}
	race3 := Race{30, 200}
	fmt.Printf("Race1: %f\n", calculateWinningPossibilities(race1))
	fmt.Printf("Race2: %f\n", calculateWinningPossibilities(race2))
	fmt.Printf("Race3: %f\n", calculateWinningPossibilities(race3))
}
