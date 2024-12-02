package main

import (
	"slices"
	"strings"

	"github.com/ysanson/AdventOfCode/pkg"
	"github.com/ysanson/AdventOfCode/pkg/execute"
)

type Iter struct {
	iteration  int
	rocksCount int
}

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

func fieldToString(field [][]rune) string {
	var sb strings.Builder
	sb.Grow(len(field[0]) * len(field))
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			sb.WriteRune(field[i][j])
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

func tiltNorth(field [][]rune) [][]rune {
	transpose := pkg.TransposeMatrix(field)
	for idx, col := range transpose {
		transpose[idx] = moveRocks(col)
	}
	return pkg.TransposeMatrix(transpose)
}

func tiltWest(field [][]rune) [][]rune {
	for idx, line := range field {
		field[idx] = moveRocks(line)
	}
	return field
}

func tiltSouth(field [][]rune) [][]rune {
	transpose := pkg.TransposeMatrix(field)
	for idx, col := range transpose {
		slices.Reverse(transpose[idx])
		transpose[idx] = moveRocks(col)
		slices.Reverse(transpose[idx])
	}
	return pkg.TransposeMatrix(transpose)
}

func tiltEast(field [][]rune) [][]rune {
	for idx, line := range field {
		slices.Reverse(field[idx])
		field[idx] = moveRocks(line)
		slices.Reverse(field[idx])
	}
	return field
}

func rotateCycle(field [][]rune, maxIterations int) int {
	records := make(map[string]Iter, 20)
	f := field
	idx := 0
	var iteration Iter
	exists := false
	//Search for the loop
	for !exists {
		f = tiltNorth(f)
		f = tiltWest(f)
		f = tiltSouth(f)
		f = tiltEast(f)

		hash := fieldToString(f)
		iteration, exists = records[hash]
		if !exists {
			records[hash] = Iter{
				iteration:  idx + 1,
				rocksCount: countRockWeight(f),
			}
		}
		idx++
	}
	// Loop is found
	loopLength := idx - iteration.iteration
	resultIdx := iteration.iteration + (maxIterations-iteration.iteration)%(loopLength)
	for _, results := range records {
		if results.iteration == resultIdx {
			return results.rocksCount
		}
	}
	return -1
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
	part2 = rotateCycle(field, 1000000000)
	return part1, part2
}

func main() {
	execute.Run(run, Tests, Puzzle, true)
}
