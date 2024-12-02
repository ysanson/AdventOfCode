package main

import (
	"strings"

	"github.com/juliangruber/go-intersect"
	"github.com/ysanson/AdventOfCode/pkg"
	"github.com/ysanson/AdventOfCode/pkg/execute"
)

type Coord struct {
	x int
	y int
}

func parseInput(input string) [][]rune {
	lines := strings.Split(input, "\n")
	result := make([][]rune, len(lines))
	for idx, line := range lines {
		result[idx] = []rune(line)
	}
	return result
}

func findEmptyLines(field [][]rune) []int {
	indices := make([]int, 0, 10)

	for idx, line := range field {
		if pkg.IsSameElement(line) {
			indices = append(indices, idx)
		}
	}
	return indices
}

func findEmptyCols(field [][]rune) []int {
	indices := make([]int, 0, 10)
	for i := 0; i < len(field[0]); i++ {
		line := getCol(field, i)
		if pkg.IsSameElement(line) {
			indices = append(indices, i)
		}
	}
	return indices
}

func getCol(field [][]rune, x int) []rune {
	col := make([]rune, len(field))
	for idx, line := range field {
		col[idx] = line[x]
	}
	return col
}

func findGalaxies(universe [][]rune, nb int) []Coord {
	galaxies := make([]Coord, 0, nb)
	for x, line := range universe {
		for y, char := range line {
			if char == '#' {
				galaxies = append(galaxies, Coord{x: x, y: y})
			}
		}
	}
	return galaxies
}

func generateRange(min int, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func calculate(galaxies []Coord, emptyLines []int, emptyCols []int, expansionFactor int) int {
	sum := 0
	for idx, first := range galaxies {
		for i := idx + 1; i < len(galaxies); i++ {
			second := galaxies[i]
			x1, x2 := pkg.Min(first.x, second.x), pkg.Max(first.x, second.x)
			y1, y2 := pkg.Min(first.y, second.y), pkg.Max(first.y, second.y)
			sum += (x2 - x1) + (len(intersect.Hash(emptyLines, generateRange(x1, x2))) * expansionFactor)
			sum += (y2 - y1) + (len(intersect.Hash(emptyCols, generateRange(y1, y2))) * expansionFactor)
		}
	}
	return sum
}

func run(input string) (any, any) {
	field := parseInput(input)
	galaxies := findGalaxies(field, strings.Count(input, "#"))
	emptyLines := findEmptyLines(field)
	emptyCols := findEmptyCols(field)
	part1 := calculate(galaxies, emptyLines, emptyCols, 1)
	part2 := calculate(galaxies, emptyLines, emptyCols, 1000000-1)

	return part1, part2
}

func main() {
	execute.Run(run, Tests, Puzzle, true)
}
