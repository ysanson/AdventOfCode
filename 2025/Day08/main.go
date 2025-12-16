package main

import (
	"math"
	"slices"
	"strings"

	"github.com/ysanson/AdventOfCode/pkg"
	"github.com/ysanson/AdventOfCode/pkg/execute"
)

type Box struct {
	X int
	Y int
	Z int
}

func (b *Box) EuclidianDistance(dest Box) float64 {
	return math.Sqrt(math.Pow(float64(b.X-dest.X), 2) + math.Pow(float64(b.Y-dest.Y), 2) + math.Pow(float64(b.Z-dest.Z), 2))
}

func (b *Box) LowerThan(other Box) bool {
	return b.X < other.X && b.Y < other.Y && b.Z < other.Z
}

func (b *Box) Equals(other Box) bool {
	return b.X == other.X && b.Y == other.Y && b.Z == other.Z
}

type BoxPair struct {
	Box1 Box
	Box2 Box
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

func findShortestDistance(distances map[BoxPair]float64) (pair BoxPair) {
	shortest := math.MaxFloat64
	for p, distance := range distances {
		if distance < shortest {
			shortest = distance
			pair = p
		}
	}
	return
}

func findCircuitIndex(circuits [][]Box, box Box) int {
	compFunc := func(c Box) bool {
		return c.Equals(box)
	}
	return slices.IndexFunc(circuits, func(el []Box) bool {
		return slices.ContainsFunc(el, compFunc)
	})
}

func longestCircuitLength(circuits [][]Box) int {
	return len(slices.MaxFunc(circuits, func(a, b []Box) int {
		return len(a) - len(b)
	}))
}

func iterate(nbIterations int, boxes []Box, distances map[BoxPair]float64) int {
	circuits := make([][]Box, len(boxes))
	for i, box := range boxes {
		circuits[i] = []Box{box}
	}
	for range nbIterations {
		pair := findShortestDistance(distances)
		box1Idx, box2Idx := findCircuitIndex(circuits, pair.Box1), findCircuitIndex(circuits, pair.Box2)
		if box1Idx != box2Idx {
			circuits[box1Idx] = append(circuits[box1Idx], circuits[box2Idx]...)
			circuits[box2Idx] = []Box{}
		}
		delete(distances, pair)
	}
	slices.SortFunc(circuits, func(a, b []Box) int {
		return len(b) - len(a)
	})

	return len(circuits[0]) * len(circuits[1]) * len(circuits[2])
}

func findLastConnection(boxes []Box, distances map[BoxPair]float64) (pair BoxPair) {
	circuits := make([][]Box, len(boxes))
	for i, box := range boxes {
		circuits[i] = []Box{box}
	}
	for longestCircuitLength(circuits) < len(boxes) {
		pair = findShortestDistance(distances)
		box1Idx, box2Idx := findCircuitIndex(circuits, pair.Box1), findCircuitIndex(circuits, pair.Box2)
		if box1Idx != box2Idx {
			circuits[box1Idx] = append(circuits[box1Idx], circuits[box2Idx]...)
			circuits[box2Idx] = []Box{}
		}
		delete(distances, pair)
	}
	return
}

func run(input string) (any, any) {
	boxes := parseInput(input)
	distances := make(map[BoxPair]float64)
	for pair := range pkg.CombinationsChan(boxes, 2) {
		distances[BoxPair{Box1: pair[0], Box2: pair[1]}] = pair[0].EuclidianDistance(pair[1])
	}
	circuits := make([][]Box, len(boxes))
	for i, box := range boxes {
		circuits[i] = []Box{box}
	}
	part1 := 0
	if len(boxes) == 20 {
		part1 = iterate(10, boxes, distances)
	} else {
		part1 = iterate(1000, boxes, distances)
	}
	lastPair := findLastConnection(boxes, distances)
	return part1, lastPair.Box1.X * lastPair.Box2.X
}

func main() {
	execute.Run(run, Tests, Puzzle, true)
}
