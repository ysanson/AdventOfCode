package main

import (
	"slices"

	"github.com/ysanson/AdventOfCode/pkg/execute"
	"github.com/ysanson/AdventOfCode/pkg/twod"
)

func checkWalls(raceMap twod.Map, path []*twod.P) (cheats int) {
	for i, pos := range path {
		for _, dir := range twod.FourDirections {
			if raceMap[pos.Pos+dir] == '#' {
				if slices.IndexFunc(path[i:], func(e *twod.P) bool { return e.Pos == pos.Pos+dir+dir }) > 100 {
					cheats++
				}
			}
		}
	}
	return
}

func checkLongCheat(path []*twod.P) (cheats int) {
	for i, p1 := range path[:len(path)-100] {
		for j, p2 := range path[100+i:] {
			if d := p2.Pos.ManhatanDistance(p1.Pos); d <= 20 && d <= j {
				cheats++
			}
		}
	}
	return
}

func run(input string) (interface{}, interface{}) {
	raceMap := twod.NewMapFromInput(input)
	start, end := twod.P{Pos: raceMap.Find('S')[0]}, raceMap.Find('E')[0]
	emptySpace := raceMap.FilterOut('#')
	baseRun, _, _ := start.ShortestPathToPos(end, emptySpace)
	return checkWalls(raceMap, baseRun), checkLongCheat(baseRun)
}

func main() {
	execute.Run(run, Tests, Puzzle, false)
}
