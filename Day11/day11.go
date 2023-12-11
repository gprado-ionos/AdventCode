package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	var space []string
	for fileScanner.Scan() {
		space = append(space, fileScanner.Text())
	}
	emptyLines, emptyColumns, galaxies := expandSpace(space)
	distance := getSumDistanceAllGalaxyPairs(space, emptyLines, emptyColumns, galaxies)
	fmt.Println(distance)
}

func expandSpace(space []string) (map[int]int, map[int]int, [][]int) {
	emptyColumns := make(map[int]int)
	emptyLines := make(map[int]int)
	var galaxies [][]int
	for i := 0; i < len(space); i++ {
		galaxyFinder := regexp.MustCompile(`[^.]+`).FindAllStringIndex(space[i], -1)
		if len(galaxyFinder) == 0 {
			emptyLines[i] = i
		} else {
			for j := 0; j < len(galaxyFinder); j++ {
				for k := 0; k < galaxyFinder[j][1]-galaxyFinder[j][0]; k++ {
					galaxies = append(galaxies, []int{i, galaxyFinder[j][k]})
				}

			}
			galaxyFinder = append(galaxyFinder, []int{len(space[i]), len(space[i])})
		}
	}
	for i := 0; i < len(space[0]); i++ {
		var emptyColumn = true
		for j := 0; j < len(space); j++ {
			if space[j][i] != '.' {
				emptyColumn = false
			}
		}
		if emptyColumn {
			emptyColumns[i] = i
		}
	}
	return emptyLines, emptyColumns, galaxies
}

func getSumDistanceAllGalaxyPairs(space []string, emptyLines map[int]int, emptyColumns map[int]int, galaxies [][]int) int {
	var sumDistance = 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			sumDistance += getDistance(galaxies[i], galaxies[j], emptyLines, emptyColumns)
		}
	}
	return sumDistance
}

func getDistance(positionGalaxy1 []int, positionGalaxy2 []int, emptyLines map[int]int, emptyColumns map[int]int) int {
	expandedSpace := findExpandedSpace(positionGalaxy1, positionGalaxy2, emptyLines, emptyColumns)

	return abs(positionGalaxy1[0]-positionGalaxy2[0]) + abs(positionGalaxy1[1]-positionGalaxy2[1]) + expandedSpace

}

func findExpandedSpace(positionGalaxy1 []int, positionGalaxy2 []int, lines map[int]int, columns map[int]int) int {
	expandedSpace := 0
	if positionGalaxy1[0] < positionGalaxy2[0] {
		for i := positionGalaxy1[0]; i < positionGalaxy2[0]; i++ {
			_, contains := lines[i]
			if contains {
				expandedSpace += 999999
			}
		}
	} else {
		for i := positionGalaxy2[0]; i < positionGalaxy1[0]; i++ {
			_, contains := lines[i]
			if contains {
				expandedSpace += 999999
			}
		}
	}

	if positionGalaxy1[1] < positionGalaxy2[1] {
		for i := positionGalaxy1[1]; i < positionGalaxy2[1]; i++ {
			_, contains := columns[i]
			if contains {
				expandedSpace += 999999
			}
		}
	} else {
		for i := positionGalaxy2[1]; i < positionGalaxy1[1]; i++ {
			_, contains := columns[i]
			if contains {
				expandedSpace += 999999
			}
		}
	}
	return expandedSpace
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
