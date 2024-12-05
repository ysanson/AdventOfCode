package main

import (
	"slices"
	"strings"

	"github.com/ysanson/AdventOfCode/pkg"
	"github.com/ysanson/AdventOfCode/pkg/execute"
)

func parseOrders(ordering string) map[int][]int {
	result := make(map[int][]int)
	for _, order := range strings.Split(ordering, "\n") {
		numbers := strings.Split(order, "|")
		first, second := pkg.MustAtoi(numbers[0]), pkg.MustAtoi(numbers[1])
		result[first] = append(result[first], second)
	}
	return result
}

func parseUpdates(updates string) [][]int {
	upd := strings.Split(updates, "\n")
	result := make([][]int, len(upd))
	for row, line := range upd {
		update := strings.Split(line, ",")
		numbers := make([]int, len(update))
		for i, num := range update {
			numbers[i] = pkg.MustAtoi(num)
		}
		result[row] = numbers
	}
	return result
}

func repairUpdate(orders map[int][]int, update []int) int {
	reordered := make([]int, 0, len(update))
	for idx, num := range update {
		reordered = append(reordered, num)
		if wrong := pkg.IntersectHash(orders[num], reordered[:idx]); len(wrong) != 0 {
			reordered = slices.DeleteFunc(reordered, func(x int) bool {
				return slices.Contains(wrong, x)
			})
			reordered = append(reordered, wrong...)
		}
	}
	return reordered[len(reordered)/2]
}

func processUpdate(orders map[int][]int, update []int) (int, int) {
	for idx, page := range update {
		if len(pkg.IntersectHash(orders[page], update[:idx])) != 0 {
			// If we have a number that should be placed after, we return
			return 0, repairUpdate(orders, update)
		}
	}
	return update[len(update)/2], 0
}

func run(input string) (interface{}, interface{}) {
	data := strings.Split(input, "\n\n")
	ordering := parseOrders(data[0])
	updates := parseUpdates(data[1])
	part1, part2 := 0, 0
	for _, update := range updates {
		first, second := processUpdate(ordering, update)
		part1 += first
		part2 += second
	}

	return part1, part2
}

func main() {
	execute.Run(run, Tests, Puzzle, false)
}
