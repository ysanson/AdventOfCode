package main

import (
	"github.com/ysanson/AdventOfCode/pkg/execute"
	"github.com/ysanson/AdventOfCode/pkg/twod"
)

func run(input string) (interface{}, interface{}) {
	basePlan := twod.NewMapFromInput(input)
	width, height := basePlan.Width(), basePlan.Height()
	antennas := basePlan.FilterOut('.')
	freq := make(map[int32][]twod.Vector)
	for antenna, val := range antennas {
		if _, ok := freq[val.(int32)]; !ok {
			freq[val.(int32)] = make([]twod.Vector, 0, 10)
		}
		freq[val.(int32)] = append(freq[val.(int32)], antenna)
	}
	antinodesP1 := make(map[twod.Vector]bool)
	antinodesP2 := make(map[twod.Vector]bool)

	for _, position := range freq {
		for i := range len(position) {
			for j := range len(position) {
				if i == j {
					continue
				}
				vec := position[i] - position[j]
				point := position[i] - vec.ScalarMult(2)
				antinodesP2[position[i]] = true
				if point.X() >= 0 && point.X() < width && point.Y() >= 0 && point.Y() < height {
					antinodesP1[point] = true
					antinodesP2[point] = true
				}
				scalar := 3
				point = position[i] - vec.ScalarMult(scalar)
				for point.X() >= 0 && point.X() < width && point.Y() >= 0 && point.Y() < height {
					antinodesP2[point] = true
					scalar++
					point = position[i] - vec.ScalarMult(scalar)
				}
			}
		}
	}

	return len(antinodesP1), len(antinodesP2)
}

func main() {
	execute.Run(run, Tests, Puzzle, false)
}
