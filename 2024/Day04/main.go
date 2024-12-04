package main

import (
	"strings"

	"github.com/ysanson/AdventOfCode/pkg/execute"
)

func countHorizontalWords(input string) int {
	return strings.Count(input, "XMAS") + strings.Count(input, "SAMX")
}

func countVerticalWords(lines []string) int {
	var invertedInput strings.Builder
	for i := range len(lines[0]) {
		for _, line := range lines {
			invertedInput.WriteByte(line[i])
		}
		invertedInput.WriteRune('\n')
	}
	return countHorizontalWords(invertedInput.String())
}

func countDiagonalWords(lines []string) int {
	count := 0
	for row := 0; row < len(lines)-3; row++ {
		for col := 0; col < len(lines[0])-3; col++ {
			zig := []byte{lines[row][col], lines[row+1][col+1], lines[row+2][col+2], lines[row+3][col+3]}
			if string(zig) == "XMAS" || string(zig) == "SAMX" {
				count++
			}
			zag := []byte{lines[row][col+3], lines[row+1][col+2], lines[row+2][col+1], lines[row+3][col]}
			if string(zag) == "XMAS" || string(zag) == "SAMX" {
				count++
			}
		}
	}
	return count
}

func countMas(lines []string) int {
	count := 0
	for row := 1; row < len(lines)-1; row++ {
		for col := 1; col < len(lines[0])-1; col++ {
			masCount := 0
			if lines[row][col] == 'A' {
				if lines[row-1][col-1] == 'M' && lines[row+1][col+1] == 'S' {
					masCount++
				} else if lines[row-1][col-1] == 'S' && lines[row+1][col+1] == 'M' {
					masCount++
				}
				if lines[row+1][col-1] == 'M' && lines[row-1][col+1] == 'S' {
					masCount++
				} else if lines[row+1][col-1] == 'S' && lines[row-1][col+1] == 'M' {
					masCount++
				}
				if masCount == 2 {
					count++
				}
			}
		}
	}
	return count
}

func run(input string) (interface{}, interface{}) {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	return countHorizontalWords(input) + countVerticalWords(lines) + countDiagonalWords(lines), countMas(lines)
}

func main() {
	execute.Run(run, Tests, Puzzle, false)
}
