package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	var calibrationValue = 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		calibrationValue += buildCalibrationNumber(line)
	}
	fmt.Println(calibrationValue)
	readFile.Close()
}

func buildCalibrationNumber(line string) int {
	fmt.Println("Line found: " + line)
	line = replaceStringNumber(line)
	strArr := []rune(line)
	fmt.Println("Line found: " + line)
	var left = ""
	var right = ""
	for i := 0; i < len(strArr); i++ {
		leftChar := strArr[i]
		rightChar := strArr[len(strArr)-i-1]
		if left == "" && unicode.IsDigit(leftChar) {
			left = string(leftChar)

		}

		if right == "" && unicode.IsDigit(rightChar) {
			right = string(rightChar)

		}

	}
	partialCalibrationValue := left + right

	fmt.Println("Partial calibration value: " + partialCalibrationValue)

	s, err := strconv.Atoi(partialCalibrationValue)
	if err != nil {
		fmt.Println("Can't convert this to an int!")
	} else {
		return s
	}
	return 0
}

func replaceStringNumber(line string) string {
	numbers := [9][2]string{
		{"one", "1"},
		{"two", "2"},
		{"three", "3"},
		{"four", "4"},
		{"five", "5"},
		{"six", "6"},
		{"seven", "7"},
		{"eight", "8"},
		{"nine", "9"}}

	strArr := []rune(line)
	var s = ""
	var partial = ""
	for i := 0; i < len(strArr); i++ {
		if unicode.IsDigit(strArr[i]) {
			s += string(strArr[i])
		} else {
			partial += string(strArr[i])
			for _, elem := range numbers {
				if strings.Contains(partial, elem[0]) {
					s += elem[1]
					partial = partial[len(partial)-len(elem[0])+1:]
				}
			}
		}

	}
	return s

	return line
}
