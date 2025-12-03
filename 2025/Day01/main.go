package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ysanson/AdventOfCode/pkg/execute"
)

func reduceRotation(rotation int) (int, clicks int) {
	for rotation > 100 {
		rotation -= 100
		clicks++
	}
	return rotation, clicks
}
func run(input string) (any, any) {
	lines := strings.Split(input, "\n")
	zeros, passOver := 0, 0
	current := 50
	for _, line := range lines {
		rot, _ := strconv.ParseInt(line[1:], 10, 0)
		rotation, rounds := reduceRotation(int(rot))
		passOver += rounds
		if line[0] == 'L' {
			if current == 0 {
				passOver--
			}
			current -= rotation
			if current < 0 {
				passOver++
				current = 100 + current
			}
		} else {
			if current+rotation > 100 {
				passOver++
			}
			current = (current + rotation) % 100
		}
		if current == 0 {
			zeros++
		}
		fmt.Printf("The dial is rotated %s to point at %d\n", line, current)
	}
	part1, part2 := zeros, zeros+passOver
	return part1, part2
}

func main() {
	execute.Run(run, Tests, Puzzle, true)
}
