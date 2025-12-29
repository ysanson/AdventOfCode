package main

import (
	"slices"
	"strings"

	"github.com/ysanson/AdventOfCode/pkg"
	"github.com/ysanson/AdventOfCode/pkg/execute"
)

type Point struct {
	X int
	Y int
}

type Rect struct {
	X1, Y1 int
	X2, Y2 int
}

func (r *Rect) Area() int {
	return (r.X2 - r.X1 + 1) * (r.Y2 - r.Y1 + 1)
}

func rectFromPoints(a, b Point) Rect {
	return Rect{
		X1: pkg.Min(a.X, b.X),
		Y1: pkg.Min(a.Y, b.Y),
		X2: pkg.Max(a.X, b.X),
		Y2: pkg.Max(a.Y, b.Y),
	}
}

func parse(input string) []Point {
	lines := strings.Split(input, "\n")
	points := make([]Point, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, ",")
		points[i] = Point{X: pkg.MustAtoi(parts[0]), Y: pkg.MustAtoi(parts[1])}
	}
	return points
}

func lineInside(pair, line Rect) bool {
	return line.X1 < pair.X2 &&
		line.Y1 < pair.Y2 &&
		line.X2 > pair.X1 &&
		line.Y2 > pair.Y1
}

func sortDescending(a, b Rect) int {
	return b.Area() - a.Area()
}

func run(input string) (any, any) {
	points := parse(input)
	area, areaWithin := 0, 0
	pairs := make([]Rect, 0)
	for pair := range pkg.CombinationsChan(points, 2) {
		pairs = append(pairs, rectFromPoints(pair[0], pair[1]))
	}
	slices.SortFunc(pairs, sortDescending)

	lines := make([]Rect, 0)
	for _, line := range pkg.AdjacentPairs(points) {
		lines = append(lines, rectFromPoints(line[0], line[1]))
	}
	slices.SortFunc(lines, sortDescending)

	area = pairs[0].Area()
	for _, p := range pairs {
		valid := true
		for _, l := range lines {
			if lineInside(p, l) {
				valid = false
				break
			}
		}
		if valid {
			areaWithin = p.Area()
			break
		}
	}

	return area, areaWithin
}

func main() {
	execute.Run(run, nil, Puzzle, true)
}
