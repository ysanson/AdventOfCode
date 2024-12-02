package main

import (
	"slices"
	"strconv"
	"strings"

	"github.com/ysanson/AdventOfCode/pkg"
	"github.com/ysanson/AdventOfCode/pkg/execute"
)

func down(line []int) [][]int {
	sequences := make([][]int, 0, len(line))
	currentLine := line
	sequences = append(sequences, currentLine)
	step := 1
	for !pkg.IsSameElement(currentLine) {
		currentLine = make([]int, len(currentLine)-1)
		for i := 0; i < len(sequences[step-1])-1; i++ {
			currentLine[i] = sequences[step-1][i+1] - sequences[step-1][i]
		}
		sequences = append(sequences, currentLine)
		step++
	}
	return sequences
}

func up(sequences [][]int) (int, int) {
	slices.Reverse(sequences)
	seq := sequences[1:]
	part1, part2 := sequences[0][0], sequences[0][0]
	for _, arr := range seq {
		part1 += arr[len(arr)-1]
		part2 = arr[0] - part2
	}
	return part1, part2
}

func extractNumbers(input string) []int {
	parts := strings.Fields(input)
	nums := make([]int, len(parts))
	for index, num := range parts {
		parsed, _ := strconv.ParseInt(num, 10, 0)
		nums[index] = int(parsed)
	}
	return nums
}

func run(input string) (any, any) {
	lines := strings.Split(input, "\n")
	part1, part2 := 0, 0
	for _, line := range lines {
		right, left := up(down(extractNumbers(line)))
		part1 += right
		part2 += left
	}
	return part1, part2
}

func main() {
	execute.Run(run, Tests, Puzzle, true)
}
