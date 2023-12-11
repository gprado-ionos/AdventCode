package main

import (
	"fmt"
	"regexp"
	"testing"
)

var space = []string{
	"...#......",
	".......#..",
	"#.........",
	"..........",
	"......#...",
	".#........",
	".........#",
	"..........",
	".......#..",
	"#...#.....",
}

func Test_getEmptySpaceToExpand(t *testing.T) {

	emptyLines, emptyColumns, galaxies := expandSpace(space)

	if len(galaxies) != 9 {
		t.Errorf("Result was incorrect, got: %d galaxies, want: %d.", len(galaxies), 9)
	}
	if len(emptyLines) != 2 || len(emptyColumns) != 3 {
		t.Errorf("Result was incorrect, got: %d emptyLines and %d emptyColumns, want: %d and %d.", len(emptyLines), len(emptyColumns), 2, 3)
	}
	if emptyLines[0] != 3 || emptyLines[1] != 7 {
		t.Errorf("Result was incorrect, got: %d and %d, want: %d and %d.", emptyLines[0], emptyLines[1], 3, 7)
	}
	if emptyColumns[0] != 2 || emptyColumns[1] != 5 || emptyColumns[2] != 8 {
		t.Errorf("Result was incorrect, got: %d, %d and %d, want: %d, %d and %d.", emptyColumns[0], emptyColumns[1], emptyColumns[2], 2, 5, 9)
	}
}

func Test_getSumDistanceAllGalaxyPairs(t *testing.T) {

	emptyLines, emptyColumns, galaxies := expandSpace(space)
	distance := getSumDistanceAllGalaxyPairs(space, emptyLines, emptyColumns, galaxies)

	if distance != 374 {
		t.Errorf("Result was incorrect, got: %d , want: %d.", distance, 374)
	}
}

func Test_aaa(t *testing.T) {
	galaxyFinder := regexp.MustCompile(`[^.]+`).FindAllStringIndex("...#...#..", -1)
	fmt.Print(galaxyFinder)
}
