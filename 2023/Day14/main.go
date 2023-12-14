package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/ysanson/AdventOfCode/2023/pkg"
	"github.com/ysanson/AdventOfCode/2023/pkg/execute"
)

func parseInput(input string) [][]rune {
	lines := strings.Split(input, "\n")
	result := make([][]rune, len(lines))
	for idx, line := range lines {
		result[idx] = []rune(line)
	}
	return result
}

func countRocksUntilBlocks(s []rune, startAt int) int {
	idx := startAt
	rocksOccurrence := 0
	for idx < len(s) && s[idx] != '#' {
		if s[idx] == 'O' {
			rocksOccurrence++
		}
		idx++
	}
	return rocksOccurrence
}

func replaceRocks(col []rune, startAt, rocksCount int) []rune {
	for i := startAt; i < len(col); i++ {
		if col[i] != '#' {
			if rocksCount > 0 {
				col[i] = 'O'
				rocksCount--
			} else {
				col[i] = '.'
			}
		} else {
			break
		}
	}
	return col
}

func moveRocks(col []rune) []rune {
	newCol := slices.Clone(col)
	idx := 0
	for idx < len(newCol) {
		if newCol[idx] != '.' {
			// If '#' or 'O', move to next char
			idx++
		} else {
			// Count rocks until next block and moves them
			nbRocks := countRocksUntilBlocks(newCol, idx)
			if nbRocks > 0 {
				newCol = replaceRocks(newCol, idx, nbRocks)
			}
			newBlocks := slices.Index(newCol[idx+nbRocks:], '#')
			if newBlocks != -1 {
				idx += slices.Index(newCol[idx+nbRocks:], '#') + nbRocks
			} else {
				idx = len(col)
			}
		}
	}
	return newCol
}

func print(field [][]rune) {
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			fmt.Printf("%c", field[i][j])
		}
		fmt.Println()
	}
}

func countRocks(line []rune) int {
	occ := 0
	for _, c := range line {
		if c == 'O' {
			occ++
		}
	}
	return occ
}

func countRockWeight(field [][]rune) int {
	nbLines := len(field)
	result := 0
	for idx, line := range field {
		result += (countRocks(line) * (nbLines - idx))
	}
	return result
}

func run(input string) (any, any) {
	part1, part2 := 0, 0
	field := parseInput(input)
	inverted := pkg.TransposeMatrix(field)
	for idx, col := range inverted {
		inverted[idx] = moveRocks(col)
	}
	inverted = pkg.TransposeMatrix(inverted)
	part1 = countRockWeight(inverted)

	return part1, part2
}

func main() {
	execute.Run(run, Tests, Puzzle, true)
}
