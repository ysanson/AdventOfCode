package main

import (
	"strings"

	"github.com/ysanson/AdventOfCode/pkg/execute"
)

func findPossibleCombinations(towels []string, pattern string, previousPatterns *map[string]int) (combinations int) {
	if comb, ok := (*previousPatterns)[pattern]; ok {
		return comb
	}
	if len(pattern) == 0 {
		combinations = 1
	} else {
		for _, towel := range towels {
			if after, found := strings.CutPrefix(pattern, towel); found {
				if comb := findPossibleCombinations(towels, after, previousPatterns); comb > 0 {
					combinations += comb
					(*previousPatterns)[after] = comb
				}
			}
		}
	}
	return
}

func run(input string) (interface{}, interface{}) {
	part1, part2 := 0, 0
	lines := strings.Split(input, "\n")
	towels, patterns := strings.Split(lines[0], ", "), lines[2:]
	previousPatterns := make(map[string]int)
	for _, pattern := range patterns {
		if count := findPossibleCombinations(towels, pattern, &previousPatterns); count > 0 {
			part1++
			part2 += count
		}
	}

	return part1, part2
}

func main() {
	execute.Run(run, Tests, Puzzle, false)
}
