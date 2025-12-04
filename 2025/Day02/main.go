package main

import (
	"strconv"
	"strings"

	"github.com/ysanson/AdventOfCode/pkg/execute"
)

func checkRepeatingPrefix(input string, prefixLength int) bool {
	if prefixLength%len(input) == 0 {
		return false
	}
	return strings.Repeat(input[:prefixLength], len(input)/prefixLength) == input
}

func run(input string) (any, any) {
	invalidP1, invalidP2 := 0, 0
	for _, r := range strings.Split(input, ",") {
		limits := strings.Split(r, "-")
		low, _ := strconv.Atoi(limits[0])
		high, _ := strconv.Atoi(limits[1])
		for i := low; i <= high; i++ {
			number := strconv.Itoa(i)
			if len(number)%2 == 0 {
				if number[:len(number)/2] == number[len(number)/2:] {
					invalidP1 += i
				}
			}
			for j := 1; j <= len(number)/2; j++ {
				if checkRepeatingPrefix(number, j) {
					invalidP2 += i
					break
				}
			}
		}
	}

	part1, part2 := invalidP1, invalidP2
	return part1, part2
}

func main() {
	execute.Run(run, Tests, Puzzle, true)
}
