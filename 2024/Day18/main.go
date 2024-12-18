package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/ysanson/AdventOfCode/pkg"
	"github.com/ysanson/AdventOfCode/pkg/execute"
)

type Position struct {
	x, y int
}

type Visit struct {
	pos  Position
	path []Position
}

var DIRECTIONS = []Position{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

func inBounds(pos Position, xMax, yMax int) bool {
	return pos.x >= 0 && pos.x <= xMax && pos.y >= 0 && pos.y <= yMax
}

func findPath(bytes []Position, size int) (dist int) {
	start, end := Position{0, 0}, Position{size, size}
	positions := []Visit{{start, []Position{}}}
	visited := make(map[Position]bool)
	fallenBytes := make(map[Position]bool)
	for _, b := range bytes {
		fallenBytes[b] = true
	}
	var visit Visit
	for len(positions) > 0 {
		visit, positions = positions[0], positions[1:]
		if visit.pos == end {
			dist = len(visit.path)
			return
		}
		if visited[visit.pos] {
			continue
		}
		visited[visit.pos] = true
		for _, d := range DIRECTIONS {
			np := Position{visit.pos.x + d.x, visit.pos.y + d.y}
			if inBounds(np, size, size) && !fallenBytes[np] {
				positions = append(positions, Visit{np, append(slices.Clone(visit.path), visit.pos)})
			}
		}

	}
	return
}

func findFirstBrokenPath(startFrom, maxSize int, bytes []Position) (pos Position) {
	for startFrom < len(bytes) {
		if findPath(bytes[:startFrom+1], maxSize) == 0 {
			pos = bytes[startFrom]
			break
		}
		startFrom++
	}
	return
}

func run(input string) (part1 interface{}, part2 interface{}) {
	b := strings.Split(input, "\n")
	bytes := make([]Position, len(b))
	for i, byte := range b {
		pos := strings.Split(byte, ",")
		bytes[i] = Position{pkg.MustAtoi(pos[0]), pkg.MustAtoi(pos[1])}
	}
	maxSize, cap := 6, 12
	if len(bytes) > 1000 {
		maxSize, cap = 70, 1024
	}
	part1 = findPath(bytes[:cap], maxSize)
	firstPos := findFirstBrokenPath(cap, maxSize, bytes)
	part2 = fmt.Sprintf("%d,%d", firstPos.x, firstPos.y)
	return
}

func main() {
	execute.Run(run, Tests, Puzzle, false)
}
