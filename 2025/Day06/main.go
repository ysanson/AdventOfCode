package main

import (
	"strconv"
	"strings"

	"github.com/ysanson/AdventOfCode/pkg"
	"github.com/ysanson/AdventOfCode/pkg/execute"
)

func run(input string) (any, any) {
	lines := strings.Split(input, "\n")
	operations := make([]rune, len(strings.Fields(lines[0])))
	for i, op := range strings.Fields(lines[len(lines)-1]) {
		operations[i] = rune(op[0])
	}
	return partOne(lines[:len(lines)-1], operations), partTwo(lines[:len(lines)-1], operations)
}

func partOne(lines []string, operations []rune) int {
	problems := make([]int, len(strings.Fields(lines[0])))
	for _, line := range lines {
		for i, num := range strings.Fields(line) {
			n, _ := strconv.Atoi(num)
			if problems[i] == 0 {
				problems[i] = n
			} else {
				switch operations[i] {
				case '+':
					problems[i] += n
				case '*':
					problems[i] *= n
				}
			}
		}
	}
	return pkg.Sum(problems...)
}

func partTwo(lines []string, operations []rune) int {
	matrix := pkg.ParseIntoMatrix(padLines(lines))
	transposed := pkg.TransposeMatrix(matrix)
	currentOperation := 0
	problems := make([]int, len(operations))
	for _, line := range transposed {
		num, err := extractNumber(line)
		if err != nil {
			currentOperation++
			continue
		}
		if problems[currentOperation] == 0 {
			problems[currentOperation] = num
		} else {
			switch operations[currentOperation] {
			case '+':
				problems[currentOperation] += num
			case '*':
				problems[currentOperation] *= num
			}
		}
	}
	return pkg.Sum(problems...)
}

func padLines(lines []string) []string {
	maxLength := 0
	for _, line := range lines {
		if len(line) > maxLength {
			maxLength = len(line)
		}
	}
	for i, line := range lines {
		if len(line) < maxLength {
			lines[i] = line + strings.Repeat(" ", maxLength-len(line))
		}
	}
	return lines
}

func extractNumber(line []rune) (int, error) {
	str := string(line)
	return strconv.Atoi(strings.TrimSpace(str))
}

func main() {
	execute.Run(run, Tests, Puzzle, true)
}
