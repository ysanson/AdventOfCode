package main

import (
	"fmt"
	"maps"
	"slices"

	"github.com/ysanson/AdventOfCode/pkg/execute"
	"github.com/ysanson/AdventOfCode/pkg/twod"
)

func computeFence(fromPos twod.Vector, plan twod.Map) (int, map[twod.Vector]bool) {
	flower := plan[fromPos]
	neighbors := plan.FindNeighbors(fromPos)
	area := make(map[twod.Vector]bool)
	area[fromPos] = true
	currentNeighbors := slices.DeleteFunc(neighbors, func(el twod.Vector) bool { return plan[el] != flower })
	perimeter := 4 - len(currentNeighbors)

	del := func(el twod.Vector) bool {
		if _, ok := area[el]; ok {
			return true
		}
		return false
	}
	for len(currentNeighbors) > 0 {
		newPoints := make([]twod.Vector, 0)
		for _, nb := range currentNeighbors {
			filtered := slices.DeleteFunc(plan.FindNeighbors(nb), func(el twod.Vector) bool { return plan[el] != flower })
			if len(filtered) != 3 {
				perimeter += 3 - len(filtered)
			}
			filtered = slices.DeleteFunc(filtered, del)
			newPoints = append(newPoints, filtered...)
			for _, val := range filtered {
				area[val] = true
			}

		}
		currentNeighbors = newPoints
	}
	fmt.Printf("Flowers %c, Perimeter is %d, area is size %d\n", flower, perimeter, len(area))

	return perimeter * len(area), area
}

func run(input string) (interface{}, interface{}) {
	plan := twod.NewMapFromInput(input)
	part1 := 0
	processed := make(map[twod.Vector]bool)
	for position := range plan {
		if _, ok := processed[position]; !ok {
			fences, area := computeFence(position, plan)
			part1 += fences
			maps.Insert(processed, maps.All(area))
		}
	}

	return part1, 0
}

func main() {
	execute.Run(run, Tests, Puzzle, false)
}
