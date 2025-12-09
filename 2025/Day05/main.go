package main

import (
	"slices"
	"strconv"
	"strings"

	"github.com/ysanson/AdventOfCode/pkg/execute"
)

type Range struct {
	start int
	end   int
}

func buildRanges(ranges []string) []Range {
	rangesArray := make([]Range, 0, len(ranges))
	for _, r := range ranges {
		nums := strings.Split(r, "-")
		start, _ := strconv.Atoi(nums[0])
		end, _ := strconv.Atoi(nums[1])
		rangesArray = append(rangesArray, Range{start: start, end: end})
	}
	slices.SortFunc(rangesArray, func(a, b Range) int {
		if a.start < b.start {
			return -1
		} else if a.start > b.start {
			return 1
		} else {
			return a.end - b.end
		}
	})
	return rangesArray
}

func countTotalRange(ranges []Range) int {
	consolidatedRanges := make([]Range, 0, len(ranges))
	for _, r := range ranges {
		if len(consolidatedRanges) == 0 {
			consolidatedRanges = append(consolidatedRanges, r)
			continue
		}
		last := consolidatedRanges[len(consolidatedRanges)-1]
		if r.start < last.start && r.end < last.end {
			consolidatedRanges = append(consolidatedRanges, r)
			continue
		} else if r.start > last.end && r.end > last.end {
			consolidatedRanges = append(consolidatedRanges, r)
			continue
		}
		if r.start < last.start {
			last.start = r.start
		}
		if r.end > last.end {
			last.end = r.end
		}
		consolidatedRanges[len(consolidatedRanges)-1] = last
	}
	sum := 0
	for _, r := range consolidatedRanges {
		sum += r.end - r.start + 1
	}
	return sum
}

func run(input string) (any, any) {
	part1, part2 := 0, 0
	parts := strings.Split(input, "\n\n")
	ranges, ids := strings.Split(parts[0], "\n"), strings.Split(parts[1], "\n")
	rangesArray := buildRanges(ranges)
	for _, id := range ids {
		num, _ := strconv.Atoi(id)
		for _, r := range rangesArray {
			if num >= r.start && num <= r.end {
				part1++
				break
			}
		}
	}
	part2 = countTotalRange(rangesArray)
	return part1, part2
}

func main() {
	execute.Run(run, Tests, Puzzle, true)
}
