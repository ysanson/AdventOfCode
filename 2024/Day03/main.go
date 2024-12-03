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
		if doIdx := strings.Index(line[idx:], "do()"); doIdx != -1 {
			return removeDontFunctions(line[:idx] + line[idx+doIdx:]) // Remove the portion between dont and do, and reiterate
		} else {
			return line[:idx] // No matching do, we keep only what's before the dont
		}
	} else {
		return line // No dont, we return the whole line
	}
}

func splitRemove(line string) string {
	if strings.Contains(line, "don't()") {
		var sb strings.Builder
		for _, part := range strings.Split(line, "do()") {
			if dontIdx := strings.Index(part, "don't()"); dontIdx != -1 {
				sb.WriteString(part[:dontIdx])
			} else {
				sb.WriteString(part)
			}
		}
		return sb.String()
	}
	return line
}

func run(input string) (interface{}, interface{}) {
	return calculateMulFunctions(input), calculateMulFunctions(removeDontFunctions(input))
}

func main() {
	execute.Run(run, Tests, Puzzle, false)
}
