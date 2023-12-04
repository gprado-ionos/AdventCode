package main

import "testing"

func Test_sumScratchcardsIntersections(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  int
	}{
		{"Card1", "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53", 8},
		{"Card2", "Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19", 2},
		{"Card3", "Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1", 2},
		{"Card4", "Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83", 1},
		{"Card5", "Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36", 0},
		{"Card6", "Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11", 0},
	}

	for _, tt := range tests {
		// tt is the test
		t.Run(tt.name, func(t *testing.T) {
			scratchcard := buildScratchcard(tt.input)
			intersection := findIntersection(scratchcard)
			got := calculateResult(intersection)
			if got != tt.want {
				t.Errorf("The game (%s) restul = %d; wanted %d", tt.input, got, tt.want)
			}
		})
	}
}

func Test_sumScratchcards(t *testing.T) {
	var arrayScratchcards = []string{"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
	}

	finalNumberOfScratchcards := 0
	var scratchcards map[int]Scratchcard = make(map[int]Scratchcard)

	for i := 0; i < len(arrayScratchcards); i++ {
		// tt is the test
		scratchcard := buildScratchcard(arrayScratchcards[i])
		scratchcards[scratchcard.cardNumber] = scratchcard
	}
	for i := 1; i <= len(scratchcards); i++ {

		finalNumberOfScratchcards += scratchcards[i].copies
		intersection := findIntersection(scratchcards[i])
		pointsToDistribute := PointsToDistribute{len(intersection), scratchcards[i].copies}
		distributeCopies(scratchcards, pointsToDistribute, scratchcards[i].cardNumber)

	}

	if finalNumberOfScratchcards != 30 {
		t.Errorf("Result = %d; wanted %d", finalNumberOfScratchcards, 30)
	}
}
