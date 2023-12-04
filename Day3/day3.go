package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

type NumberIndex struct {
	initial int
	end     int
}

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	var lines []string
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	total := part1(lines)
	total2 := part2(lines)
	fmt.Println(total)
	fmt.Println(total2)
}

func part2(lines []string) int {
	total := 0
	for i := 1; i < len(lines)-1; i++ {
		total += getGearRatio(lines[i-1], lines[i], lines[i+1])

	}
	return total
}

func part1(lines []string) int {
	total := 0
	for i := 0; i < len(lines); i++ {
		if i == 0 {
			total += getSumNumbersAdjacentToSymbolsOfOneLine(lines[i], lines[i+1])
		} else if i == len(lines)-1 {
			total += getSumNumbersAdjacentToSymbolsOfOneLine(lines[i], lines[i-1])
		} else {
			total += getSumNumbersAdjacentToSymbolsOfTwoLines(lines[i-1], lines[i], lines[i+1])
		}
	}
	return total
}

func getSumNumbersAdjacentToSymbolsOfOneLine(lineUnderEval string, lineToCompare string) int {
	partialSum := 0
	allNumbersIndexes := regexp.MustCompile("[0-9]+").FindAllStringIndex(lineUnderEval, -1)
	for i := 0; i < len(allNumbersIndexes); i++ {
		numberIdx := NumberIndex{allNumbersIndexes[i][0], allNumbersIndexes[i][1]}
		if doesItHaveAdjacentSymbolsBeforeOrAfter(lineUnderEval, numberIdx) {
			number, _ := strconv.Atoi(lineUnderEval[numberIdx.initial:numberIdx.end])
			partialSum += number
		} else {
			if doesTheAdjacentLineHaveSymbols(numberIdx, lineToCompare) {
				number, _ := strconv.Atoi(lineUnderEval[numberIdx.initial:numberIdx.end])
				partialSum += number
			}
		}
	}

	return partialSum
}

func doesItHaveAdjacentSymbolsBeforeOrAfter(line string, numberIdx NumberIndex) bool {
	if len(line) == numberIdx.end {
		if line[numberIdx.initial-1] != '.' {
			return true
		}
	} else if numberIdx.initial > 0 && len(line) > numberIdx.end-1 {
		if line[numberIdx.initial-1] != '.' || line[numberIdx.end] != '.' {
			return true
		}
	} else if numberIdx.initial == 0 {
		if line[numberIdx.end] != '.' {
			return true
		}
	} else {
		if line[numberIdx.initial-1] != '.' {
			return true
		}
	}
	return false
}

func doesTheAdjacentLineHaveSymbols(numberIdx NumberIndex, lineForComparison string) bool {
	if numberIdx.initial > 0 && len(lineForComparison) > numberIdx.end {
		substringOfNextLineToBeEval := lineForComparison[numberIdx.initial-1 : numberIdx.end+1]
		if doesItHaveSymbols(substringOfNextLineToBeEval) {
			return true
		}
	} else if numberIdx.initial == 0 {
		substringOfNextLineToBeEval := lineForComparison[:numberIdx.end+1]
		if doesItHaveSymbols(substringOfNextLineToBeEval) {
			return true
		}
	} else {
		substringOfNextLineToBeEval := lineForComparison[numberIdx.initial-1:]
		if doesItHaveSymbols(substringOfNextLineToBeEval) {
			return true
		}
	}
	return false
}

func doesItHaveSymbols(stringToEval string) bool {
	allString := regexp.MustCompile("[^0-9.]").FindAllString(stringToEval, -1)
	if allString != nil && len(allString) > 0 {
		return true
	}
	return false
}

func getSumNumbersAdjacentToSymbolsOfTwoLines(previousLine string, lineUnderEval string, nextLine string) int {
	partialSum := 0

	allNumbersIndexes := regexp.MustCompile("[0-9]+").FindAllStringIndex(lineUnderEval, -1)
	for i := 0; i < len(allNumbersIndexes); i++ {
		numberIdx := NumberIndex{allNumbersIndexes[i][0], allNumbersIndexes[i][1]}
		if doesItHaveAdjacentSymbolsBeforeOrAfter(lineUnderEval, numberIdx) {
			number, _ := strconv.Atoi(lineUnderEval[numberIdx.initial:numberIdx.end])
			partialSum += number
		} else {
			if doesTheAdjacentLineHaveSymbols(numberIdx, previousLine) || doesTheAdjacentLineHaveSymbols(numberIdx, nextLine) {
				number, _ := strconv.Atoi(lineUnderEval[numberIdx.initial:numberIdx.end])
				partialSum += number
			}
		}
	}

	return partialSum
}

func getGearRatio(previousLine string, lineUnderEval string, nextLine string) int {
	partialSum := 0

	allNumbersIndexes := regexp.MustCompile("[*]+").FindAllStringIndex(lineUnderEval, -1)
	for i := 0; i < len(allNumbersIndexes); i++ {
		gear := []int{-1, -1}
		gearIdx := allNumbersIndexes[i][0]
		findGearInLine(lineUnderEval, gearIdx, gear)
		findGearInAdjacentLine(previousLine, gearIdx, gear)
		findGearInAdjacentLine(nextLine, gearIdx, gear)
		partialSum += calculateGearRatio(gear)
	}

	return partialSum
}

func calculateGearRatio(gear []int) int {
	if gear[0] != -1 && gear[1] != -1 {
		return gear[0] * gear[1]
	}
	return 0
}

func findGearInAdjacentLine(line string, gearIdx int, gear []int) {
	partOfGear := getNumberSameIndexGear(line, gearIdx)
	if partOfGear != -1 {
		if addToGearArray(gear, partOfGear) {
			return
		}
	}
	findGearInLine(line, gearIdx, gear)
}

func findGearInLine(line string, gearIdx int, gear []int) {
	if len(line) == gearIdx+1 {
		partOfGear := getNumberBeforeGear(line, gearIdx)
		if partOfGear != -1 {
			if addToGearArray(gear, partOfGear) {
				return
			}
		}
	} else if gearIdx == 0 {
		partOfGear := getNumberAfterGear(line, gearIdx)
		if partOfGear != -1 {
			if addToGearArray(gear, partOfGear) {
				return
			}
		}

	} else {
		partOfGear := getNumberBeforeGear(line, gearIdx)
		if partOfGear != -1 {
			if addToGearArray(gear, partOfGear) {
				return
			}
		}
		partOfGear = getNumberAfterGear(line, gearIdx)
		if partOfGear != -1 {
			if addToGearArray(gear, partOfGear) {
				return
			}
		}
	}
}

func addToGearArray(gear []int, partOfGear int) bool {
	if gear[0] == -1 {
		gear[0] = partOfGear
		return false
	} else {
		gear[1] = partOfGear
		return true
	}
}

func getNumberSameIndexGear(line string, gearIdx int) int {
	if unicode.IsDigit(rune(line[gearIdx])) {
		allNumbersIndexes := regexp.MustCompile("[0-9]+").FindAllStringIndex(line, -1)
		for i := 0; i < len(allNumbersIndexes); i++ {
			if allNumbersIndexes[i][0] <= gearIdx && allNumbersIndexes[i][1] >= gearIdx {
				number, _ := strconv.Atoi(line[allNumbersIndexes[i][0]:allNumbersIndexes[i][1]])
				return number
			}
		}
	}
	return -1
}
func getNumberBeforeGear(line string, gearIdx int) int {
	if unicode.IsDigit(rune(line[gearIdx-1])) {
		allNumbersIndexes := regexp.MustCompile("[0-9]+").FindAllStringIndex(line, -1)
		for i := 0; i < len(allNumbersIndexes); i++ {
			if allNumbersIndexes[i][1] == gearIdx {
				number, _ := strconv.Atoi(line[allNumbersIndexes[i][0]:allNumbersIndexes[i][1]])
				return number
			}
		}
	}
	return -1
}

func getNumberAfterGear(line string, gearIdx int) int {
	if unicode.IsDigit(rune(line[gearIdx+1])) {
		allNumbersIndexes := regexp.MustCompile("[0-9]+").FindAllStringIndex(line, -1)
		for i := 0; i < len(allNumbersIndexes); i++ {
			if allNumbersIndexes[i][0] == gearIdx+1 {
				number, _ := strconv.Atoi(line[allNumbersIndexes[i][0]:allNumbersIndexes[i][1]])
				return number
			}
		}
	}
	return -1
}
