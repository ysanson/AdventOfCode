package main

import (
	"github.com/ysanson/AdventOfCode/pkg/execute"
	"github.com/ysanson/AdventOfCode/pkg/twod"
)

func iterate(points map[twod.Vector]int, beamMap twod.Map, splits *int) map[twod.Vector]int {
	newPoints := make(map[twod.Vector]int)
	for point, beams := range points {
		if beamMap[point+twod.DOWN] == '^' {
			newPoints[point+twod.DOWNLEFT] += beams
			newPoints[point+twod.DOWNRIGHT] += beams
			*splits++
		} else {
			newPoints[point+twod.DOWN] += beams
		}
	}
	return newPoints
}

func run(input string) (any, any) {
	beamMap := twod.NewMapFromInput(input)
	startPoint := beamMap.Find('S')[0]
	beamPoints := make(map[twod.Vector]int)
	beamPoints[startPoint] = 1
	splits := 0
	for range beamMap.Height() {
		beamPoints = iterate(beamPoints, beamMap, &splits)
	}
	part2 := 0
	for _, beams := range beamPoints {
		part2 += beams
	}
	return splits, part2
}

func main() {
	execute.Run(run, Tests, Puzzle, true)
}
