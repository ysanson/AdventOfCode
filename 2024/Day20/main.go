package main

import (
	"fmt"
	"runtime"

	"github.com/ysanson/AdventOfCode/pkg/execute"
	"github.com/ysanson/AdventOfCode/pkg/twod"
)

func findCheatableWalls(raceMap twod.Map, path []*twod.P) (canRemove map[twod.Vector]bool) {
	canRemove = make(map[twod.Vector]bool)
	findNeighbors := func(pos twod.Vector) twod.Vector {
		if raceMap[pos+twod.LEFT] == '#' && raceMap[pos+twod.LEFT+twod.LEFT] == '.' {
			return pos + twod.LEFT
		} else if raceMap[pos+twod.RIGHT] == '#' && raceMap[pos+twod.RIGHT+twod.RIGHT] == '.' {
			return pos + twod.RIGHT
		} else if raceMap[pos+twod.UP] == '#' && raceMap[pos+twod.UP+twod.UP] == '.' {
			return pos + twod.UP
		} else if raceMap[pos+twod.DOWN] == '#' && raceMap[pos+twod.DOWN+twod.DOWN] == '.' {
			return pos + twod.DOWN
		} else {
			return 0
		}
	}
	for _, pos := range path {
		if pos := findNeighbors(pos.Pos); pos != 0 {
			canRemove[pos] = true
		}
	}
	return
}

func runCheat(emptySpaces twod.Map, start twod.P, end twod.Vector, removedWalls <-chan twod.Vector, results chan<- float64) {
	cloneMap := func(toAdd twod.Vector) twod.Map {
		newMap := emptySpaces.Clone()
		newMap[toAdd] = '.'
		return newMap
	}
	for pos := range removedWalls {
		_, dist, _ := start.ShortestPathToPos(end, cloneMap(pos))
		results <- dist - 1
	}
}

func run(input string) (interface{}, interface{}) {
	part1, part2 := 0, 0
	raceMap := twod.NewMapFromInput(input)
	start, end := twod.P{Pos: raceMap.Find('S')[0]}, raceMap.Find('E')[0]
	emptySpace := raceMap.FilterOut('#')
	baseRun, dist, found := start.ShortestPathToPos(end, emptySpace)
	dist -= 1
	fmt.Printf("Base run is %d, dist is %f, found %v\n", len(baseRun), dist, found)
	removableWalls := findCheatableWalls(raceMap, baseRun)
	cheatWalls := make(chan twod.Vector, len(removableWalls))
	results := make(chan float64, len(removableWalls))
	for range runtime.NumCPU() {
		go runCheat(emptySpace, start, end, cheatWalls, results)
	}
	for pos := range removableWalls {
		cheatWalls <- pos
	}
	close(cheatWalls)
	for range len(removableWalls) {
		if d := <-results; dist-d >= 100 {
			part1++
		}
	}
	return part1, part2
}

func main() {
	execute.Run(run, nil, Puzzle, false)
}
