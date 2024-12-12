package main

import (
	"maps"
	"math"
	"strings"

	"github.com/ysanson/AdventOfCode/pkg"
	"github.com/ysanson/AdventOfCode/pkg/execute"
)

func countStones(input map[int]int, blinks int) int {
	current, stones := input, input
	for range blinks {
		stones = make(map[int]int)
		for stone, count := range current {
			if stone == 0 {
				stones[1] += count
			} else if l := math.Floor(math.Log10(float64(stone))) + 1; int(l)%2 == 0 {
				stones[stone/int(math.Pow(10, l/2))] += count
				stones[stone%int(math.Pow(10, l/2))] += count
			} else {
				stones[stone*2024] += count
			}
		}
		current = stones
	}
	sum := 0
	for count := range maps.Values(stones) {
		sum += count
	}
	return sum
}

func run(input string) (interface{}, interface{}) {
	numStr := strings.Split(input, " ")
	numbers := make(map[int]int, len(numStr))
	for _, n := range numStr {
		numbers[pkg.MustAtoi(n)] = 1
	}

	return countStones(numbers, 25), countStones(numbers, 75)
}

func main() {
	execute.Run(run, Tests, Puzzle, false)
}
