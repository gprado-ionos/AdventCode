package main

import "testing"

func Test_buildCalibrationNumber(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  int
	}{
		// the table itself
		{"86vbnpsixthreetwonevng should be 82", "86vbnpsixthreetwonevng", 82},
		{"seventwo4gnrsrpnfppseven2 should be 72", "seventwo4gnrsrpnfppseven2", 72},
		{"tthree4one9 should be 39", "tthree4one9", 39},
		{"eightwothree should be 83", "eightwothree", 83},
		{"abcone2threexyz should be 13", "abcone2threexyz", 13},
		{"xtwone3four should be 24", "xtwone3four", 24},
		{"4nineeightseven2 should be 42", "4nineeightseven2", 42},
		{"zoneight234 should be 14", "zoneight234", 14},
		{"7pqrstsixteen should be 76", "7pqrstsixteen", 76},
		{"2twone should be 21", "2twone", 21},
	}

	for _, tt := range tests {
		// tt is the test
		t.Run(tt.name, func(t *testing.T) {
			got := buildCalibrationNumber(tt.input)
			if got != tt.want {
				t.Errorf("buildCalibrationNumber(%s) = %d; want %d", tt.input, got, tt.want)
			}
		})
	}
}

func Test_replaceStringNumber(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  string
	}{
		// the table itself
		{"86vbnpsixthreetwonevng should be 866321", "86vbnpsixthreetwonevng", "866321"},
		{"seventwo4gnrsrpnfppseven2 should be 72472", "seventwo4gnrsrpnfppseven2", "72472"},
		{"tthree4one9 should be 3419", "tthree4one9", "3419"},
		{"eightwothree should be 823", "eightwothree", "823"},
		{"abcone2threexyz should be 123", "abcone2threexyz", "123"},
		{"xtwone3four should be 2134", "xtwone3four", "2134"},
		{"4nineeightseven2 should be 49872", "4nineeightseven2", "49872"},
		{"zoneight234 should be 18234", "zoneight234", "18234"},
		{"7pqrstsixteen should be 76", "7pqrstsixteen", "76"},
		{"2twone should be 221", "2twone", "221"},
	}

	for _, tt := range tests {
		// tt is the test
		t.Run(tt.name, func(t *testing.T) {
			got := replaceStringNumber(tt.input)
			if got != tt.want {
				t.Errorf("replaceStringNumber(%s) = %s; want %s", tt.input, got, tt.want)
			}
		})
	}
}
