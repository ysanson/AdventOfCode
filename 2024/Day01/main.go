package main

import (
	"sort"
	"strings"

	"github.com/ysanson/AdventOfCode/pkg"
	"github.com/ysanson/AdventOfCode/pkg/execute"
)

func extractLists(lines []string) ([]int, []int) {
	var first = make([]int, len(lines))
	var second = make([]int, len(lines))
	for i, line := range lines {
		var elements = strings.Split(line, " ")
		first[i] = pkg.MustAtoi(elements[0])
		second[i] = pkg.MustAtoi(elements[len(elements)-1])
	}
	sort.Sort(sort.IntSlice(first))
	sort.Sort(sort.IntSlice(second))
	return first, second
}

func run(input string) (interface{}, interface{}) {
	lines := strings.Split(input, "\n")
	part1, part2 := 0, 0
	left, right := extractLists(lines)
	zipped := pkg.Zip(left, right)    // Create a interator with the 2 lists to make it simpler to iterate over
	previous, prevCalculation := 0, 0 // Used for part 2, to skip the computation in case of multiple same numbers

	rightCount := map[int]int{} // Map to count the left elements
	for _, i := range right {
		if count, ok := rightCount[i]; !ok {
			// We register the number in the map
			rightCount[i] = 1
		} else if count > 0 {
			rightCount[i]++
		}
	}

	for first, second := range zipped {
		part1 += pkg.Abs(first - second)
		if first == previous {
			part2 += prevCalculation // We skip the calculation if it's the same number as before
		} else {
			previous = first
			if count, ok := rightCount[first]; ok {
				prevCalculation = first * count
				part2 += prevCalculation
			}
		}
	}
	return part1, part2
}

func main() {
	execute.Run(run, Tests, Puzzle, true)
}
