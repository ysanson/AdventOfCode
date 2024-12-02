package main

import (
	"slices"
	"strings"

	"github.com/ysanson/AdventOfCode/pkg"
	"github.com/ysanson/AdventOfCode/pkg/execute"
)

func extractReport(line string) []int {
	split := strings.Split(line, " ")
	report := make([]int, len(split))
	for i, level := range split {
		report[i] = pkg.MustAtoi(level)
	}
	return report
}

func isSafe(report []int) bool {
	isAscending := report[0] < report[1]
	for i := 0; i < len(report)-1; i++ {
		dist := pkg.Abs(report[i] - report[i+1])
		if (isAscending && report[i] > report[i+1]) || (!isAscending && report[i] < report[i+1]) || dist > 3 || dist < 1 {
			return false
		}
	}
	return true
}

func checkDampenedReport(report []int) bool {
	if isSafe(report) {
		return true
	}
	for i := 0; i < len(report); i++ {
		reportCopy := make([]int, len(report))
		copy(reportCopy, report)
		compensatedReport := slices.Delete(reportCopy, i, i+1)

		if isSafe(compensatedReport) {
			return true
		}
	}
	return false
}

func run(input string) (interface{}, interface{}) {
	part1, part2 := 0, 0
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		report := extractReport(line)
		if isSafe(report) {
			part1++
		}
		if checkDampenedReport(report) {
			part2++
		}
	}
	return part1, part2
}

func main() {
	execute.Run(run, Tests, Puzzle, false)
}
