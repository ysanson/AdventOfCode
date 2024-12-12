package main

import (
	"maps"
	"slices"

	"github.com/ysanson/AdventOfCode/pkg"
	"github.com/ysanson/AdventOfCode/pkg/execute"
	"github.com/ysanson/AdventOfCode/pkg/twod"
)

func countCorners(pos twod.Vector, neighbors map[string]twod.Vector, plan twod.Map) int {
	if len(neighbors) == 1 {
		return 2
	}
	corners := 0
	if len(neighbors) == 2 && (!pkg.SetContainsAll(neighbors, "up", "down") && !pkg.SetContainsAll(neighbors, "left", "right")) {
		corners = 1
	}
	if flower, ok := plan[pos+twod.UPLEFT]; ok && flower != plan[pos] && pkg.SetContainsAll(neighbors, "up", "left") {
		corners++
	}
	if flower, ok := plan[pos+twod.UPRIGHT]; ok && flower != plan[pos] && pkg.SetContainsAll(neighbors, "up", "right") {
		corners++
	}
	if flower, ok := plan[pos+twod.DOWNLEFT]; ok && flower != plan[pos] && pkg.SetContainsAll(neighbors, "down", "left") {
		corners++
	}
	if flower, ok := plan[pos+twod.DOWNRIGHT]; ok && flower != plan[pos] && pkg.SetContainsAll(neighbors, "down", "right") {
		corners++
	}
	return corners
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
	if len(ngb) == 0 {
		return 4, 4, area
	}
	perimeter := 4 - len(ngb)
	corners := countCorners(fromPos, ngb, plan)
	neighbors := slices.Collect(maps.Values(ngb))

	processNeighbors := func(neighbors []twod.Vector) []twod.Vector {
		newPoints := make([]twod.Vector, 0)
		for _, nb := range neighbors {
			area[nb] = true
			filtered := delUnrelatedNeighbor(plan.FindNeighbors(nb))
			perimeter += 4 - len(filtered)
			corners += countCorners(nb, filtered, plan)
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
	return perimeter, corners, area
}

func run(input string) (interface{}, interface{}) {
	plan := twod.NewMapFromInput(input)
	part1, part2 := 0, 0
	processed := make(map[twod.Vector]bool)
	for position := range plan {
		if _, ok := processed[position]; !ok {
			perimeter, corners, area := computeFence(position, plan)
			part1 += perimeter * len(area)
			part2 += corners * len(area)
			maps.Insert(processed, maps.All(area))
		}
	}

	return part1, part2
}

func main() {
	execute.Run(run, Tests, Puzzle, false)
}
