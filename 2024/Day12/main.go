package main

import (
	"maps"
	"slices"

	"github.com/ysanson/AdventOfCode/pkg"
	"github.com/ysanson/AdventOfCode/pkg/execute"
	"github.com/ysanson/AdventOfCode/pkg/twod"
)

func countCorners(neighbors map[string]twod.Vector) int {
	switch len(neighbors) {
	case 1:
		return 2
	case 2:
		if pkg.SetContainsAll(neighbors, "up", "down") || pkg.SetContainsAll(neighbors, "left", "right") {
			return 0
		} else {
			return 1
		}
	default:
		return 0
	}
}

func computeFence(fromPos twod.Vector, plan twod.Map) (int, int, map[twod.Vector]bool) {
	flower := plan[fromPos]
	area := map[twod.Vector]bool{fromPos: true}

	delUnrelatedNeighbor := func(ngb map[string]twod.Vector) map[string]twod.Vector {
		for dir, pos := range ngb {
			if plan[pos] != flower {
				delete(ngb, dir)
			}
		}
		return ngb
	}
	ngb := delUnrelatedNeighbor(plan.FindNeighbors(fromPos))
	perimeter := 4 - len(ngb)
	corners := countCorners(ngb)
	neighbors := slices.Collect(maps.Values(ngb))

	processNeighbors := func(neighbors []twod.Vector) []twod.Vector {
		newPoints := make([]twod.Vector, 0)
		for _, nb := range neighbors {
			area[nb] = true
			filtered := delUnrelatedNeighbor(plan.FindNeighbors(nb))
			perimeter += 4 - len(filtered)
			corners += countCorners(filtered)
			for dir, pos := range filtered {
				if _, ok := area[pos]; ok {
					delete(filtered, dir)
				} else {
					newPoints = append(newPoints, pos)
				}
			}
			for _, val := range filtered {
				area[val] = true
			}
		}
		return newPoints
	}

	for len(neighbors) > 0 {
		neighbors = processNeighbors(neighbors)
	}
	return perimeter * len(area), corners, area
}

func run(input string) (interface{}, interface{}) {
	plan := twod.NewMapFromInput(input)
	part1, part2 := 0, 0
	processed := make(map[twod.Vector]bool)
	for position := range plan {
		if _, ok := processed[position]; !ok {
			fences, corners, area := computeFence(position, plan)
			part1 += fences
			part2 += corners
			maps.Insert(processed, maps.All(area))
		}
	}

	return part1, part2
}

func main() {
	execute.Run(run, Tests, Puzzle, false)
}
