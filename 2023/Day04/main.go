package main

import (
	"slices"
	"strings"

	"github.com/ysanson/AdventOfCode/2023/pkg"
	"github.com/ysanson/AdventOfCode/2023/pkg/execute"
)

func getWinningNumbers(game string) []int {
	winning := game[strings.IndexRune(game, ':')+1 : strings.IndexRune(game, '|')]
	numbers := strings.Split(winning, " ")
	var nums []int
	for _, number := range numbers {
		if number != "" {
			nums = append(nums, pkg.MustAtoi(number))
		}
	}
	return nums
}

func getCardNumbers(game string) []int {
	extract := game[strings.IndexRune(game, '|')+1:]
	numbers := strings.Split(extract, " ")
	var nums []int
	for _, number := range numbers {
		if number != "" {
			nums = append(nums, pkg.MustAtoi(number))
		}
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

	copies := make(map[int]int)
	for i := 0; i < len(lines); i++ {
		copies[i] = 1
	}

	for index, line := range lines {
		expectedNumbers, gameNumbers = getNumbers(line)
		winningNumbers = countWinningNumbers(expectedNumbers, gameNumbers)
		for j := 0; j < winningNumbers; j++ {
			if index+j < len(lines) {
				copies[index+j+1] += copies[index]
			}
		}
		if winningNumbers > 0 {
			part1 += 1 << (winningNumbers - 1)
		}

	}

	part2 := 0
	for _, copy := range copies {
		part2 += copy
	}
	return part1, part2
}

func main() {
	execute.Run(run, Tests, Puzzle, true)
}
