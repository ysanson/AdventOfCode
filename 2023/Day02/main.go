package main

import (
	"strings"

	"github.com/ysanson/AdventOfCode/pkg"
	"github.com/ysanson/AdventOfCode/pkg/execute"
)

type Color string

const (
	Red   Color = " red"
	Green Color = " green"
	Blue  Color = " blue"
)

func getGameId(input string) int {
	game := input[strings.Index(input, " ")+1 : strings.Index(input, ":")]
	return pkg.MustAtoi(game)
}

func getColoredBalls(input string, color Color) int {
	index := strings.Index(input, string(color))
	if index != -1 {
		extractedNum := input[strings.LastIndex(input[:index], " "):index]
		return pkg.MustAtoi(strings.TrimSpace(extractedNum))
	} else {
		return 0
	}
}

func checkGame(game string, nbRed int, nbGreen int, nbBlue int) (int, int) {
	gameId := getGameId(game)
	sets := strings.Split(game[strings.Index(game, ":"):], ";")
	isDoable := true
	var red, green, blue = 0, 0, 0
	var minRed, minGreen, minBlue = 0, 0, 0

	for _, set := range sets {
		if set != "" {
			red = getColoredBalls(set, Red)
			green = getColoredBalls(set, Green)
			blue = getColoredBalls(set, Blue)
			minRed = pkg.Max(minRed, red)
			minGreen = pkg.Max(minGreen, green)
			minBlue = pkg.Max(minBlue, blue)
			if red > nbRed || green > nbGreen || blue > nbBlue {
				isDoable = false
			}
		}
	}
	if isDoable {
		return gameId, pkg.Multiply(minRed, minGreen, minBlue)
	} else {
		return 0, pkg.Multiply(minRed, minGreen, minBlue)
	}
}

func run(inputs string) (interface{}, interface{}) {
	lines := strings.Split(inputs, "\n")
	sumOfGames := 0
	sumOfMins := 0
	for _, line := range lines {
		if line != "" {
			gameId, power := checkGame(line, 12, 13, 14)
			sumOfGames += gameId
			sumOfMins += power
		}
	}

	return sumOfGames, sumOfMins
}

func main() {
	execute.Run(run, Tests, Puzzle, true)
}
