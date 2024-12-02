package main

import (
	"slices"
	"strings"

	"github.com/ysanson/AdventOfCode/pkg"
	"github.com/ysanson/AdventOfCode/pkg/execute"
)

func getWinningNumbers(game string) []int {
	winning := game[strings.IndexRune(game, ':')+1 : strings.IndexRune(game, '|')]
	numbers := strings.Split(strings.TrimSpace(winning), " ")
	nums := make([]int, len(numbers))
	for idx, number := range numbers {
		nums[idx] = pkg.MustAtoi(number)
	}
	return nums
}

func getCardNumbers(game string) []int {
	extract := game[strings.IndexRune(game, '|')+1:]
	numbers := strings.Split(strings.TrimSpace(extract), " ")
	nums := make([]int, len(numbers))
	for idx, number := range numbers {
		nums[idx] = pkg.MustAtoi(number)
	}
	return nums
}

func getNumbers(game string) ([]int, []int) {
	return getWinningNumbers(game), getCardNumbers(game)
}

func countWinningNumbers(expectedNumbers []int, gameNumbers []int) int {
	winningNumbers := 0
	for _, num := range gameNumbers {
		if slices.Contains(expectedNumbers, num) {
			winningNumbers++
		}
	}
	return winningNumbers
}

func run(input string) (interface{}, interface{}) {
	lines := strings.Split(input, "\n")
	part1 := 0
	var expectedNumbers, gameNumbers []int
	winningNumbers := 0

	copies := pkg.CreateSlice(len(lines), 1)

	for index, line := range lines {
		expectedNumbers, gameNumbers = getNumbers(pkg.StandardizeSpaces(line))
		winningNumbers = countWinningNumbers(expectedNumbers, gameNumbers)
		for j := 0; j < winningNumbers; j++ {
			if index+j+1 < len(lines) {
				copies[index+j+1] += copies[index]
			}
		}
		if winningNumbers > 0 {
			part1 += 1 << (winningNumbers - 1)
		}

	}

	part2 := pkg.Reduce(copies, func(acc, current int) int { return acc + current }, 0)

	return part1, part2
}

func main() {
	execute.Run(run, Tests, Puzzle, true)
}
