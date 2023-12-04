package main

import (
	"regexp"
	"testing"
)

func Test_sumNumbersWithAdjacentSymbols(t *testing.T) {
	var given = []string{"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598.."}

	result := 0
	for i := 0; i < len(given); i++ {
		if i == 0 {
			result += getSumNumbersAdjacentToSymbolsOfOneLine(given[i], given[i+1])
		} else if i == len(given)-1 {
			result += getSumNumbersAdjacentToSymbolsOfOneLine(given[i], given[i-1])
		} else {
			result += getSumNumbersAdjacentToSymbolsOfTwoLines(given[i-1], given[i], given[i+1])
		}
	}

	if result != 4361 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 4361)
	}
}

func Test_sumNumbersWithAdjacentSymbols2(t *testing.T) {
	var given = []string{"12.......*..",
		"+.........34",
		".......-12..",
		"..78........",
		"..*....60...",
		"78..........",
		".......23...",
		"....90*12...",
		"............",
		"2.2......12.",
		".*.........*",
		"1.1.......56"}

	result := 0
	for i := 0; i < len(given); i++ {
		if i == 0 {
			result += getSumNumbersAdjacentToSymbolsOfOneLine(given[i], given[i+1])
		} else if i == len(given)-1 {
			result += getSumNumbersAdjacentToSymbolsOfOneLine(given[i], given[i-1])
		} else {
			result += getSumNumbersAdjacentToSymbolsOfTwoLines(given[i-1], given[i], given[i+1])
		}
	}

	if result != 413 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 413)
	}
}

func Test_sumNumbersWithAdjacentSymbols3(t *testing.T) {
	var given = []string{"*12.......*..",
		"..........34",
		".......-12..",
		"..78........",
		"..*....60...",
		"78..........",
		".......23...",
		"....90/12...",
		"............",
		"2.2......12.",
		".*.......*..",
		"1.1.......56"}

	result := 0
	for i := 0; i < len(given); i++ {
		if i == 0 {
			result += getSumNumbersAdjacentToSymbolsOfOneLine(given[i], given[i+1])
		} else if i == len(given)-1 {
			result += getSumNumbersAdjacentToSymbolsOfOneLine(given[i], given[i-1])
		} else {
			result += getSumNumbersAdjacentToSymbolsOfTwoLines(given[i-1], given[i], given[i+1])
		}
	}

	if result != 413 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 413)
	}
}

func Test_sumNumbersWithAdjacentSymbols4(t *testing.T) {
	var given = []string{"....401.............425.......323......791......697...............963............................................420........................",
		"...*..................................%......#.....*....290.........................492.............656...@953.....................+830.....",
		"..159...........823...33.717.....572.......806...896......-.....335....834......815.............791....*..............776..................."}

	result := 0
	for i := 0; i < len(given); i++ {
		if i == 0 {
			result += getSumNumbersAdjacentToSymbolsOfOneLine(given[i], given[i+1])
		} else if i == len(given)-1 {
			result += getSumNumbersAdjacentToSymbolsOfOneLine(given[i], given[i-1])
		} else {
			result += getSumNumbersAdjacentToSymbolsOfTwoLines(given[i-1], given[i], given[i+1])
		}
	}

	if result != 6479 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 6479)
	}
}

func Test_sumNumbersWithAdjacentSymbols5(t *testing.T) {
	var given = []string{"12.......*..",
		"+.........34",
		".......-12..",
		"..78........",
		"..*....60...",
		"78.........9",
		".5.....23..$",
		"8...90*12...",
		"............",
		"2.2......12.",
		".*.........*",
		"1.1..503+.56"}

	result := 0
	for i := 0; i < len(given); i++ {
		if i == 0 {
			result += getSumNumbersAdjacentToSymbolsOfOneLine(given[i], given[i+1])
		} else if i == len(given)-1 {
			result += getSumNumbersAdjacentToSymbolsOfOneLine(given[i], given[i-1])
		} else {
			result += getSumNumbersAdjacentToSymbolsOfTwoLines(given[i-1], given[i], given[i+1])
		}
	}

	if result != 925 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 925)
	}
}

func Test_sumNumbersWithAdjacentSymbols6(t *testing.T) {
	var given = []string{"12.......*12",
		"..........34",
		".......-12..",
	}

	result := 0
	for i := 0; i < len(given); i++ {
		if i == 0 {
			result += getSumNumbersAdjacentToSymbolsOfOneLine(given[i], given[i+1])
		} else if i == len(given)-1 {
			result += getSumNumbersAdjacentToSymbolsOfOneLine(given[i], given[i-1])
		} else {
			result += getSumNumbersAdjacentToSymbolsOfTwoLines(given[i-1], given[i], given[i+1])
		}
	}

	if result != 58 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 58)
	}
}

func Test_doesItHaveAdjacentSymbolsBeforeOrAfter(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  bool
	}{
		{"Example1", "+.........34", false},
		{"Example2", "34.&.........", false},
		{"Example3", "%.34.&.........", false},
		{"Example4", ".........&34", true},
		{"Example5", "....555^....", true},
		{"Example6", "....*555....", true},
		{"Example7", "*5555.......", true},
	}

	for _, tt := range tests {
		// tt is the test
		t.Run(tt.name, func(t *testing.T) {
			allNumbersIndexes := regexp.MustCompile("[0-9]+").FindAllStringIndex(tt.input, -1)
			for i := 0; i < len(allNumbersIndexes); i++ {
				numberIdx := NumberIndex{allNumbersIndexes[i][0], allNumbersIndexes[i][1]}
				andSo := doesItHaveAdjacentSymbolsBeforeOrAfter(tt.input, numberIdx)
				if andSo != tt.want {
					t.Errorf("Restul = %t; wanted %t", andSo, tt.want)
				}
			}

		})
	}
}

func Test_aa(t *testing.T) {
	given := "534...534"
	index := regexp.MustCompile("[0-9]+").FindAllStringIndex(given, -1)
	for i := 0; i < len(index); i++ {
		println(index[i][0])
	}

}

func Test_aa2(t *testing.T) {
	given := "534.*..*.534"
	regexp.MustCompile("[*]+").FindAllStringIndex(given, -1)

}

func Test_sumGearRatio(t *testing.T) {
	var given = []string{"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598.."}

	result := 0
	for i := 1; i < len(given)-1; i++ {
		result += getGearRatio(given[i-1], given[i], given[i+1])

	}

	if result != 467835 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 467835)
	}
}

func Test_sumGearRatio2(t *testing.T) {
	var given = []string{".......229.............727..425............../..........................228........#..$...............385....&...........695................",
		"...................978.*....*....................700....*.........256..@.........625......311.170...+................642*...................",
		".492.....983.........-.577.743..................*....267.521........%....479.......................696.980......................*..........."}

	result := 0
	for i := 1; i < len(given)-1; i++ {
		result += getGearRatio(given[i-1], given[i], given[i+1])

	}

	if result != 1320551 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 1320551)
	}
}
