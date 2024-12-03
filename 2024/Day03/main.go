package main

import (
	"regexp"
	"strings"

	"github.com/ysanson/AdventOfCode/pkg"
	"github.com/ysanson/AdventOfCode/pkg/execute"
)

func calculateMulFunctions(line string) int {
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	numbers := regexp.MustCompile(`\d+`)
	result := 0
	for _, match := range re.FindAllString(line, -1) {
		nums := numbers.FindAllString(match, -1)
		result += pkg.MustAtoi(nums[0]) * pkg.MustAtoi(nums[1])
	}
	return result
}

func removeDontFunctions(line string) string {
	if idx := strings.Index(line, "don't()"); idx != -1 {
		if doIdx := strings.Index(line[idx+6:], "do()"); doIdx != -1 {
			return removeDontFunctions(line[:idx+6] + line[idx+doIdx+6:]) // Remove the portion between dont and do, and reiterate
		} else {
			return line[:idx+6] // No matching do, we keep only what's before the dont
		}
	} else {
		return line // No dont, we return the whole line
	}
}

func run(input string) (interface{}, interface{}) {
	return calculateMulFunctions(input), calculateMulFunctions(removeDontFunctions(input))
}

func main() {
	execute.Run(run, Tests, Puzzle, false)
}
