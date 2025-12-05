package main

import (
	"math"
	"slices"
	"strings"

	"github.com/ysanson/AdventOfCode/pkg/execute"
)

func convertArrayToNumber(digits []int) (result int) {
	slices.Reverse(digits)
	for i, digit := range digits {
		result += digit * int(math.Pow10(i))
	}
	return
}

func findHighestNComb(digits, result []int, length int) int {
	if length == 0 {
		return convertArrayToNumber(result)
	}
	max := slices.Max(digits[:len(digits)-length+1])
	maxIndex := slices.Index(digits, max)
	if maxIndex == len(digits)-length {
		return convertArrayToNumber(slices.Concat(result, digits[maxIndex:]))
	}
	return findHighestNComb(digits[maxIndex+1:], append(result, max), length-1)
}

func convertStringToIntArray(number string) (result []int) {
	for _, digit := range number {
		result = append(result, int(digit-'0'))
	}
	return
}

func run(input string) (any, any) {
	part1, part2 := 0, 0
	for _, line := range strings.Split(input, "\n") {
		number := convertStringToIntArray(line)
		part1 += findHighestNComb(number, make([]int, 0, 2), 2)
		part2 += findHighestNComb(number, make([]int, 0, 12), 12)
	}

	return part1, part2
}

func main() {
	execute.Run(run, Tests, Puzzle, true)
}
