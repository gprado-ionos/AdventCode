package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	var mapPipes []string
	for fileScanner.Scan() {
		mapPipes = append(mapPipes, fileScanner.Text())
	}
	startPoint := findStartingPoint(mapPipes)
	possibleDirectionsToStart := findAllAdjacentPoint(startPoint, mapPipes)
	possibleDirectionsToStart = filterValidDirections(possibleDirectionsToStart, mapPipes)
	farthest := findFarthestPoint(startPoint, possibleDirectionsToStart, mapPipes)
	fmt.Println(farthest)

	drawOutputFiles(startPoint, possibleDirectionsToStart, mapPipes)
	countEnclosedTiles()

}

func countEnclosedTiles() {
	var pipeFiles []string = []string{"pipe0.txt", "pipe1.txt", "pipe2.txt"}
	for _, pipeFile := range pipeFiles {
		readFile, err := os.Open(pipeFile)

		if err != nil {
			fmt.Println(err)
		}
		fileScanner := bufio.NewScanner(readFile)

		fileScanner.Split(bufio.ScanLines)
		var mapPipes []string
		for fileScanner.Scan() {
			mapPipes = append(mapPipes, fileScanner.Text())
		}
		enclosedCount := countEnclosedTilesInPipe(mapPipes)
		fmt.Println(pipeFile + ":" + strconv.Itoa(enclosedCount))
	}
}

func countEnclosedTilesInPipe(pipes []string) int {
	var enclosedCount int
	for i := 1; i < len(pipes)-1; i++ {
		for j := 1; j < len(pipes[i])-1; j++ {

			//if pipes[i][j] != '*' {
			//	for isEnclosed(i, j, pipes) {
			//
			//	}
			//	enclosedCount++
			//}
		}
	}
	return enclosedCount

}

func findStartingPoint(mapPipes []string) [2]int {
	var startPoint [2]int
	for i := 0; i < len(mapPipes); i++ {
		for j := 0; j < len(mapPipes[i]); j++ {
			if mapPipes[i][j] == 'S' {
				startPoint = [2]int{i, j}
			}
		}
	}
	return startPoint
}

func findAllAdjacentPoint(startPoint [2]int, pipes []string) [][]int {
	var adjacentPoints [][]int
	var lineAbove = startPoint[0] - 1
	var lineBelow = startPoint[0] + 1
	var columnLeft = startPoint[1] - 1
	var columnRight = startPoint[1] + 1
	adjacentPoints = append(adjacentPoints, []int{lineAbove, startPoint[1]})
	adjacentPoints = append(adjacentPoints, []int{lineBelow, startPoint[1]})
	adjacentPoints = append(adjacentPoints, []int{startPoint[0], columnLeft})
	adjacentPoints = append(adjacentPoints, []int{startPoint[0], columnRight})

	return adjacentPoints
}

func filterValidDirections(start [][]int, pipes []string) [][]int {
	var validDirections [][]int
	for i := 0; i < len(start); i++ {
		if start[i][0] >= 0 && start[i][0] < len(pipes) && start[i][1] >= 0 && start[i][1] < len(pipes[start[i][0]]) {
			if pipes[start[i][0]][start[i][1]] != '.' {
				validDirections = append(validDirections, start[i])
			}
		}
	}
	return validDirections
}

func findFarthestPoint(startPoint [2]int, validDirections [][]int, pipes []string) int {
	var farthestPoint int

	for i := 0; i < len(validDirections); i++ {
		isValid := true
		var tmpFarthestPoint int = 0
		var tmpValidDirection [2]int = [2]int(validDirections[i])
		var tmpStartPoint [2]int = startPoint
		for isValid {
			tmpFarthestPoint++
			tmpValidDirection, tmpStartPoint, isValid = getValidNextStep(tmpValidDirection, tmpStartPoint, pipes)
		}
		if tmpFarthestPoint > farthestPoint {
			if !isValid && string(pipes[tmpStartPoint[0]][tmpStartPoint[1]]) == "S" {
				farthestPoint = tmpFarthestPoint / 2
			} else {
				farthestPoint = tmpFarthestPoint
			}
		}

	}
	return farthestPoint
}

func getValidNextStep(point [2]int, startPoint [2]int, pipes []string) ([2]int, [2]int, bool) {
	tmpPoint := [2]int(point)
	var isValid bool = true
	var lineAbove = startPoint[0] - 1
	var lineBelow = startPoint[0] + 1
	var columnLeft = startPoint[1] - 1
	var columnRight = startPoint[1] + 1
	line := point[0]
	column := point[1]
	lineStr := pipes[line]
	shape := string(lineStr[column])
	switch shape {
	case "|":
		if line == lineAbove {
			startPoint = [2]int{line - 1, startPoint[1]}
		} else if line == lineBelow {
			startPoint = [2]int{line + 1, startPoint[1]}
		} else {
			isValid = false
		}
	case "J":
		if column == columnRight {
			startPoint = [2]int{startPoint[0] - 1, columnRight}
		} else if line == lineBelow {
			startPoint = [2]int{lineBelow, startPoint[1] - 1}
		} else {
			isValid = false
		}
	case "L":
		if column == columnLeft {
			startPoint = [2]int{startPoint[0] - 1, columnLeft}
		} else if line == lineBelow {
			startPoint = [2]int{lineBelow, startPoint[1] + 1}
		} else {
			isValid = false
		}
	case "F":
		if line == lineAbove {
			startPoint = [2]int{lineAbove, startPoint[1] + 1}
		} else if column == columnLeft {
			startPoint = [2]int{startPoint[0] + 1, columnLeft}
		} else {
			isValid = false
		}
	case "7":
		if line == lineAbove {
			startPoint = [2]int{lineAbove, startPoint[1] - 1}
		} else if column == columnRight {
			startPoint = [2]int{startPoint[0] + 1, columnRight}
		} else {
			isValid = false
		}
	case "-":
		if column == columnRight {
			startPoint = [2]int{startPoint[0], columnRight + 1}
		} else if column == columnLeft {
			startPoint = [2]int{startPoint[0], columnLeft - 1}
		} else {
			isValid = false
		}
	case ".":
		isValid = false
	case "S":
		isValid = false
	}
	return startPoint, tmpPoint, isValid
}

func drawOutputFiles(startPoint [2]int, validDirections [][]int, pipes []string) {
	var farthestPoint int

	for i := 0; i < len(validDirections); i++ {
		tmpPipe := make([]string, len(pipes))

		copy(tmpPipe, pipes)
		fileName := "pipe" + strconv.Itoa(i) + ".txt"
		isValid := true
		var tmpFarthestPoint int = 0
		var tmpValidDirection [2]int = [2]int(validDirections[i])
		var tmpStartPoint [2]int = startPoint
		tmpPipe[tmpValidDirection[0]] = replaceAtIndex(tmpPipe[tmpValidDirection[0]], '*', tmpValidDirection[1])
		for isValid {
			tmpFarthestPoint++
			tmpValidDirection, tmpStartPoint, isValid = getValidNextStep(tmpValidDirection, tmpStartPoint, pipes)
			tmpPipe[tmpValidDirection[0]] = replaceAtIndex(tmpPipe[tmpValidDirection[0]], '*', tmpValidDirection[1])
		}
		writeToFile(tmpPipe, fileName)
		if tmpFarthestPoint > farthestPoint {
			if !isValid && string(pipes[tmpStartPoint[0]][tmpStartPoint[1]]) == "S" {
				farthestPoint = tmpFarthestPoint / 2
			} else {
				farthestPoint = tmpFarthestPoint
			}
		}

	}
}

func writeToFile(pipe []string, fileName string) {
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	for _, v := range pipe {
		fmt.Fprintln(f, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err1 := f.Close()
	if err1 != nil {
		fmt.Println(err1)
		return
	}
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}
