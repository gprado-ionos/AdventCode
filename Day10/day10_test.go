package main

import (
	"testing"
)

func Test_findStartPoint(t *testing.T) {
	var mapPipes = []string{
		"..F7.",
		".FJ|.",
		"SJ.L7",
		"|F--J",
		"LJ...",
	}

	startPoint := findStartingPoint(mapPipes)

	if startPoint[0] != 2 || startPoint[1] != 0 {
		t.Errorf("Result was incorrect, got position: %d and %d, want: %d and %d.", startPoint[0], startPoint[1], 2, 0)
	}
}

func Test_findAllAdjacentFromStartingPoint(t *testing.T) {
	var mapPipes = []string{
		"..F7.",
		".FJ|.",
		"SJ.L7",
		"|F--J",
		"LJ...",
	}

	startPoint := findStartingPoint(mapPipes)
	possibleDirectionsToStart := findAllAdjacentPoint(startPoint, mapPipes)
	possibleDirectionsToStart = filterValidDirections(possibleDirectionsToStart, mapPipes)

	if len(possibleDirectionsToStart) != 2 {
		t.Errorf("Result was incorrect, got  %d starting points want: %d .", len(possibleDirectionsToStart), 2)
	}
}

func Test_findFarthestFromStartPoint(t *testing.T) {
	var mapPipes = []string{
		"..F7.",
		".FJ|.",
		"SJ.L7",
		"|F--J",
		"LJ...",
	}

	startPoint := findStartingPoint(mapPipes)
	possibleDirectionsToStart := findAllAdjacentPoint(startPoint, mapPipes)
	possibleDirectionsToStart = filterValidDirections(possibleDirectionsToStart, mapPipes)
	farthest := findFarthestPoint(startPoint, possibleDirectionsToStart, mapPipes)

	if farthest != 8 {
		t.Errorf("Result was incorrect, got  %d  want: %d .", farthest, 8)
	}
}
