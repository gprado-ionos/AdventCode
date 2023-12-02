package main

import "testing"

func Test_gameOfCubes(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  bool
	}{
		{"Game1", "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", true},
		{"Game2", "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", true},
		{"Game3", "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", false},
		{"Game4", "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", false},
		{"Game5", "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", true},
	}

	var numberOfCubes map[string]int = map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	for _, tt := range tests {
		// tt is the test
		t.Run(tt.name, func(t *testing.T) {
			game := buildGame(tt.input)
			got := validateGame(game, numberOfCubes)
			if got != tt.want {
				t.Errorf("The game (%s) restul = %t; wanted %t", tt.input, got, tt.want)
			}
		})
	}
}

func Test_parseGameValue(t *testing.T) {
	var given = "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	game := buildGame(given)
	if game.gameValue != 1 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", game.gameValue, 1)
	}
}

func Test_parseGamePlay(t *testing.T) {
	var given = "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	game := buildGame(given)
	if len(game.play) != 3 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", len(game.play), 3)
	}
}

func Test_validateFewestPlayable(t *testing.T) {
	var given = "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	var fewest map[string]int = map[string]int{
		"red":   4,
		"green": 2,
		"blue":  6,
	}
	var colors = []string{"red", "green", "blue"}

	game := buildGame(given)
	playable := validateFewestPlayable(game)
	fewestPlayable := calculateFewestPlayable(playable)
	for _, color := range colors {
		if playable[color] != fewest[color] {
			t.Errorf("Result was incorrect, got: %d, want: %d.", playable[color], fewest[color])
		}
	}

	if fewestPlayable != 48 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", fewestPlayable, 48)
	}
}
