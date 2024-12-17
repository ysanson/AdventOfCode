package main

import (
	"github.com/ysanson/AdventOfCode/pkg/execute"
	"github.com/ysanson/AdventOfCode/pkg/twod"
)

func computeScore(prevDir, nextDir twod.Vector) int {
	if prevDir == nextDir {
		return 1
	}
	return 1001
}

func run(input string) (interface{}, interface{}) {
	plan := twod.NewMapFromInput(input)
	start, end := plan.Find('S')[0], plan.Find('E')[0]

	score, path := plan.Djikstra(start, end, computeScore)
	uniqueTiles := plan.GetUniqueTilesCount(start, end, path, computeScore)

	return score, uniqueTiles
}

func main() {
	execute.Run(run, Tests, Puzzle, false)
}
