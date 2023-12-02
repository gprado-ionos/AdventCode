package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var numberOfCubes map[string]int = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

type Game struct {
	gameValue int
	play      []Play
}

type Play struct {
	game map[string]int
}

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	var fewestPlayableResult = 0
	var result = 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		game := buildGame(line)
		if validateGame(game, numberOfCubes) {
			result += game.gameValue
		}
		playable := validateFewestPlayable(game)
		fewestPlayableResult += calculateFewestPlayable(playable)
	}
	fmt.Println(result)
	fmt.Println(fewestPlayableResult)
}

func calculateFewestPlayable(playable map[string]int) int {
	var result = 1
	for _, value := range playable {
		result *= value
	}
	return result

}

func validateFewestPlayable(game Game) map[string]int {
	var coloredValues map[string]int = map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	for _, play := range game.play {
		for cubeColor, cubeNumber := range play.game {
			if cubeNumber > coloredValues[cubeColor] {
				coloredValues[cubeColor] = cubeNumber
			}
		}
	}
	return coloredValues
}

func validateGame(game Game, numberOfCubes map[string]int) bool {
	var result = true
	for _, play := range game.play {
		for cubeColor, cubeNumber := range play.game {
			if cubeNumber > numberOfCubes[cubeColor] {
				result = false
			}
		}
	}
	return result

}

func buildGame(gameInput string) Game {
	var game Game
	gameInputSplit := strings.Split(gameInput, ":")
	gameResult, _ := strconv.Atoi(regexp.MustCompile(`[^0-9]+`).ReplaceAllString(gameInputSplit[0], ""))
	game.gameValue = gameResult

	gamesString := strings.Split(gameInputSplit[1], ";")
	for _, gameString := range gamesString {
		var play Play
		coloredCubeValues := strings.Split(gameString, ",")
		result := make(map[string]int)
		for _, coloredCubeValue := range coloredCubeValues {
			numberOfCubes, _ := strconv.Atoi(regexp.MustCompile(`[^0-9]+`).ReplaceAllString(coloredCubeValue, ""))
			cubeColor := regexp.MustCompile(`[0-9 ]+`).ReplaceAllString(coloredCubeValue, "")
			result[cubeColor] = numberOfCubes
		}
		play.game = result
		game.play = append(game.play, play)
	}
	return game
}
