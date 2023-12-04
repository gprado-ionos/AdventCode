package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Scratchcard struct {
	cardNumber       int
	winningNumbers   []string
	numbersScratched []string
	copies           int
}

type PointsToDistribute struct {
	wins   int
	copies int
}

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	result := 0
	finalNumberOfScratchcards := 0
	var scratchcards map[int]Scratchcard = make(map[int]Scratchcard)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		scratchcard := buildScratchcard(line)
		scratchcards[scratchcard.cardNumber] = scratchcard
	}
	for i := 1; i <= len(scratchcards); i++ {
		finalNumberOfScratchcards += scratchcards[i].copies
		intersection := findIntersection(scratchcards[i])
		pointsToDistribute := PointsToDistribute{len(intersection), scratchcards[i].copies}
		distributeCopies(scratchcards, pointsToDistribute, scratchcards[i].cardNumber)
		result += calculateResult(intersection)
	}

	fmt.Println(result)
	fmt.Println(finalNumberOfScratchcards)
}

func distributeCopies(scratchcards map[int]Scratchcard, pointsToDistribute PointsToDistribute, winningCardNumber int) {

	for i := winningCardNumber + 1; i < len(scratchcards) && pointsToDistribute.wins > 0; i++ {
		// First we get a "copy" of the entry
		if entry, ok := scratchcards[i]; ok {

			// Then we modify the copy
			entry.copies += pointsToDistribute.copies

			// Then we reassign map entry
			scratchcards[i] = entry
		}
		pointsToDistribute.wins--
	}

}

func calculateResult(intersection []int) int {
	result := 0
	for i := 1; i <= len(intersection); i++ {
		if i == 1 {
			result = 1
		} else {
			result = result * 2
		}
	}
	return result
}

func buildScratchcard(value string) Scratchcard {
	input := strings.Split(value, ":")
	allNumbersFromInput := strings.Split(input[1], "|")
	winningNumbers := strings.Split(allNumbersFromInput[0], " ")
	numbersScratched := strings.Split(allNumbersFromInput[1], " ")
	cardNumber, _ := strconv.Atoi(regexp.MustCompile(`[^0-9]+`).ReplaceAllString(input[0], ""))
	scratchcard := Scratchcard{cardNumber, winningNumbers, numbersScratched, 1}
	return scratchcard
}

func findIntersection(scratchcard Scratchcard) []int {
	var result []int
	for _, winningNumber := range scratchcard.winningNumbers {
		winningNum, err := strconv.Atoi(winningNumber)
		if err == nil {
			for _, numberScratched := range scratchcard.numbersScratched {
				numberScr, err := strconv.Atoi(numberScratched)
				if err == nil {
					if winningNum == numberScr {
						result = append(result, winningNum)
					}
				}

			}
		}

	}
	return result
}
