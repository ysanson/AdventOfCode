package main

import (
	"slices"
	"strings"

	"github.com/ysanson/AdventOfCode/pkg/execute"
)

func parseInput(input string) [][]rune {
	lines := strings.Split(input, "\n")
	result := make([][]rune, len(lines))
	for idx, line := range lines {
		result[idx] = []rune(line)
	}
	return result
}

func getCol(field [][]rune, x int) []rune {
	col := make([]rune, len(field))
	for idx, line := range field {
		col[idx] = line[x]
	}
	return col
}

func sameOrOneDifferent(s1 []rune, s2 []rune, includeDiff bool) (bool, bool) {
	if includeDiff {
		oneDiff := false
		for i := 0; i < len(s1); i++ {
			if s1[i] != s2[i] {
				if !oneDiff {
					oneDiff = true
				} else {
					return false, false
				}
			}
		}
		return true, oneDiff
	} else {
		return slices.Compare(s1, s2) == 0, false
	}
}

func isLineReflected(field [][]rune, startAt int, includeSmudge bool) bool {
	diff := 0
	same, lineSmudge, smudgeIncluded := true, false, false
	for same {
		diff++
		if startAt-diff < 0 || startAt+diff+1 >= len(field) {
			if includeSmudge {
				return smudgeIncluded
			} else {
				return true
			}
		}
		up, down := field[startAt-diff], field[startAt+diff+1]
		same, lineSmudge = sameOrOneDifferent(up, down, includeSmudge)
		if lineSmudge {
			// If there's a smudge
			if smudgeIncluded {
				// If we already considered the smudge, return false
				return false
			} else {
				// Smudge is considered now
				smudgeIncluded = true
			}
		}
	}
	return false
}

func isColReflected(field [][]rune, startAt int, includeSmudge bool) bool {
	diff := 0
	same, lineSmudge, smudgeIncluded := true, false, false

	for same {
		diff++
		if startAt-diff < 0 || startAt+diff+1 >= len(field[0]) {
			if includeSmudge {
				return smudgeIncluded
			} else {
				return true
			}
		}
		left, right := getCol(field, startAt-diff), getCol(field, startAt+diff+1)
		same, lineSmudge = sameOrOneDifferent(left, right, includeSmudge)
		if lineSmudge {
			// If there's a smudge
			if smudgeIncluded {
				// If we already considered the smudge, return false
				return false
			} else {
				// Smudge is considered now
				smudgeIncluded = true
			}
		}
	}
	return false
}

func findReflection(field [][]rune, includeSmudge bool) int {
	endCol, endLine := false, false
	colIndex, lineIndex := 0, 0

	// Column search
	for i := 0; i < len(field[0])-1; i++ {
		if same, hasSmudge := sameOrOneDifferent(getCol(field, i), getCol(field, i+1), includeSmudge); same {
			if isColReflected(field, i, includeSmudge && !hasSmudge) {
				endCol = true
				colIndex = i + 1
				break // We found, no need to search further
			}
		}
	}

	if !endCol {
		// Line search
		for i := 0; i < len(field)-1; i++ {
			if same, hasSmudge := sameOrOneDifferent(field[i], field[i+1], includeSmudge); same {
				if isLineReflected(field, i, includeSmudge && !hasSmudge) {
					endLine = true
					lineIndex = i + 1
					break // We found, no need to search further
				}
			}
		}
	}

	if endCol {
		// Reflection on the column
		return colIndex
	} else if endLine {
		// Reflection on the line
		return lineIndex * 100
	}
	return 0
}

func run(input string) (any, any) {
	part1, part2 := 0, 0
	for _, field := range strings.Split(input, "\n\n") {
		pattern := parseInput(field)
		part1 += findReflection(pattern, false)
		part2 += findReflection(pattern, true)
	}

	return part1, part2
}

func main() {
	execute.Run(run, Tests, Puzzle, true)
}
