package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/ysanson/AdventOfCode/pkg"
	"github.com/ysanson/AdventOfCode/pkg/execute"
)

type Position struct {
	x, y int
}

var DIRECTIONS = []Position{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

func inBounds(pos Position, xMax, yMax int) bool {
	return pos.x >= 0 && pos.x <= xMax && pos.y >= 0 && pos.y <= yMax
}

// never marks a cell that no byte falls on.
const never = math.MaxInt

// fallOrder returns, for each cell, the index of the byte landing on it.
func fallOrder(bytes []Position, size int) []int {
	order := make([]int, (size+1)*(size+1))
	for i := range order {
		order[i] = never
	}
	for i, b := range bytes {
		cell := b.y*(size+1) + b.x
		order[cell] = min(order[cell], i)
	}
	return order
}

// findPath returns the shortest distance across the grid once the first
// nbFallen bytes have landed, or -1 if the exit cannot be reached.
func findPath(order []int, nbFallen, size int) int {
	dist := make([]int, len(order)) // doubles as the visited set: -1 is unseen
	for i := range dist {
		dist[i] = -1
	}
	queue := make([]Position, 0, len(order))
	queue = append(queue, Position{0, 0})
	dist[0] = 0
	end := len(order) - 1

	for head := 0; head < len(queue); head++ {
		pos := queue[head]
		current := dist[pos.y*(size+1)+pos.x]
		for _, d := range DIRECTIONS {
			np := Position{pos.x + d.x, pos.y + d.y}
			if !inBounds(np, size, size) {
				continue
			}
			cell := np.y*(size+1) + np.x
			if dist[cell] != -1 || order[cell] < nbFallen {
				continue
			}
			dist[cell] = current + 1
			if cell == end {
				return dist[cell]
			}
			queue = append(queue, np)
		}
	}
	return -1
}

// findFirstBrokenPath binary searches for the first byte blocking the exit.
// Bytes never leave, so once the exit is blocked it stays blocked.
func findFirstBrokenPath(startFrom, maxSize int, bytes []Position, order []int) (pos Position) {
	remaining := len(bytes) - startFrom
	offset := sort.Search(remaining, func(i int) bool {
		return findPath(order, startFrom+i+1, maxSize) == -1
	})
	if offset < remaining {
		pos = bytes[startFrom+offset]
	}
	return
}

func run(input string) (part1 any, part2 any) {
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
	order := fallOrder(bytes, maxSize)
	part1 = findPath(order, cap, maxSize)
	firstPos := findFirstBrokenPath(cap, maxSize, bytes, order)
	part2 = fmt.Sprintf("%d,%d", firstPos.x, firstPos.y)
	return
}

func main() {
	execute.Run(run, Tests, Puzzle, false)
}
