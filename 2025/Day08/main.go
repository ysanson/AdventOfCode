package main

import (
	"cmp"
	"slices"
	"strings"

	"github.com/ysanson/AdventOfCode/pkg"
	"github.com/ysanson/AdventOfCode/pkg/execute"
)

type Box struct {
	X, Y, Z int
}

// squaredDistance orders pairs the same way as the real distance, without
// floating point.
func (b Box) squaredDistance(other Box) int {
	dx, dy, dz := b.X-other.X, b.Y-other.Y, b.Z-other.Z
	return dx*dx + dy*dy + dz*dz
}

type BoxPair struct {
	box1, box2 int // indexes into the box list
	distance   int
}

// Circuits is a union-find over box indexes.
type Circuits struct {
	parent []int
	size   []int
}

func newCircuits(n int) *Circuits {
	c := &Circuits{parent: make([]int, n), size: make([]int, n)}
	for i := range n {
		c.parent[i] = i
		c.size[i] = 1
	}
	return c
}

func (c *Circuits) find(box int) int {
	for c.parent[box] != box {
		c.parent[box] = c.parent[c.parent[box]]
		box = c.parent[box]
	}
	return box
}

// connect joins two circuits, returning false if they were already one.
func (c *Circuits) connect(a, b int) bool {
	a, b = c.find(a), c.find(b)
	if a == b {
		return false
	}
	if c.size[a] < c.size[b] {
		a, b = b, a
	}
	c.parent[b] = a
	c.size[a] += c.size[b]
	return true
}

func (c *Circuits) largestThreeProduct() int {
	sizes := make([]int, 0, len(c.parent))
	for box, parent := range c.parent {
		if box == parent {
			sizes = append(sizes, c.size[box])
		}
	}
	slices.SortFunc(sizes, func(a, b int) int { return b - a })
	return sizes[0] * sizes[1] * sizes[2]
}

func parseInput(input string) []Box {
	lines := strings.Split(input, "\n")
	boxes := make([]Box, len(lines))
	for i, line := range lines {
		coords := strings.Split(line, ",")
		boxes[i] = Box{
			X: pkg.MustAtoi(coords[0]),
			Y: pkg.MustAtoi(coords[1]),
			Z: pkg.MustAtoi(coords[2]),
		}
	}
	return boxes
}

func sortedPairs(boxes []Box) []BoxPair {
	pairs := make([]BoxPair, 0, len(boxes)*(len(boxes)-1)/2)
	for i, box := range boxes {
		for j := i + 1; j < len(boxes); j++ {
			pairs = append(pairs, BoxPair{i, j, box.squaredDistance(boxes[j])})
		}
	}
	slices.SortFunc(pairs, func(a, b BoxPair) int { return cmp.Compare(a.distance, b.distance) })
	return pairs
}

func run(input string) (any, any) {
	boxes := parseInput(input)
	connections := 1000
	if len(boxes) == 20 {
		connections = 10
	}

	circuits := newCircuits(len(boxes))
	remaining := len(boxes)
	part1, part2 := 0, 0
	for i, pair := range sortedPairs(boxes) {
		if i == connections {
			part1 = circuits.largestThreeProduct()
		}
		if circuits.connect(pair.box1, pair.box2) {
			remaining--
			if remaining == 1 {
				part2 = boxes[pair.box1].X * boxes[pair.box2].X
			}
		}
		if i >= connections && remaining == 1 {
			break
		}
	}
	return part1, part2
}

func main() {
	execute.Run(run, Tests, Puzzle, true)
}
