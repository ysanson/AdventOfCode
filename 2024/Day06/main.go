package main

import (
	"runtime"
	"slices"

	"github.com/ysanson/AdventOfCode/pkg/execute"
	"github.com/ysanson/AdventOfCode/pkg/twod"
)

type VisitedPoint struct {
	Pos       twod.Vector
	Direction twod.Vector
}

func getNextDirection(curDir twod.Vector) twod.Vector {
	switch curDir {
	case twod.UP:
		return twod.RIGHT
	case twod.RIGHT:
		return twod.DOWN
	case twod.DOWN:
		return twod.LEFT
	case twod.LEFT:
		return twod.UP
	default:
		return twod.UP
	}
}

func runPart1(plan twod.Map, startingPos *twod.P, width, height *int) (twod.Map, []VisitedPoint) {
	position := startingPos
	visited := make([]VisitedPoint, 0)
	for !position.IsOutOfMapLimits(&plan, width, height) {
		visited = append(visited, VisitedPoint{Pos: position.Pos, Direction: position.Speed})
		plan.UpdatePosition(position.Pos, 'X')
		dest := position.GetPositionAtDestination(position.Speed, 1)
		if plan[dest] == '#' {
			position.Speed = getNextDirection(position.Speed)
		}
		position.Move(1)
	}
	return plan, visited
}

func runWithObstacle(plan *twod.Map, position twod.P, upd twod.Vector, path *[]VisitedPoint, width, height *int) int {
	wallHit := make(map[VisitedPoint]bool)
	var current VisitedPoint
	if skip := slices.IndexFunc(*path, func(e VisitedPoint) bool { return e.Pos == upd }); skip > 0 {
		pos := (*path)[skip-1]
		position.Pos = pos.Pos
		position.Speed = pos.Direction
	}
	for !position.IsOutOfMapLimits(plan, width, height) {
		for dest := position.GetPositionAtDestination(position.Speed, 1); (*plan)[dest] == '#' || dest == upd; dest = position.GetPositionAtDestination(position.Speed, 1) {
			current = VisitedPoint{Pos: position.Pos, Direction: position.Speed}
			if _, ok := wallHit[current]; ok {
				return 1
			}
			wallHit[current] = true
			position.Speed = getNextDirection(position.Speed)
		}
		position.Move(1)
	}
	return 0
}

func worker(baseMap *twod.Map, startingPoint *twod.P, path *[]VisitedPoint, width, height *int, jobs <-chan twod.Vector, results chan<- int) {
	for position := range jobs {
		results <- runWithObstacle(baseMap, *startingPoint.Clone(), position, path, width, height)
	}
}

func computeWorkerGroup(positions []twod.Vector, baseMap *twod.Map, startingPoint *twod.P, path *[]VisitedPoint, width, height *int) int {
	n := len(positions)
	jobs := make(chan twod.Vector, n)
	resChannel := make(chan int, n)
	for range runtime.NumCPU() {
		go worker(baseMap, startingPoint, path, width, height, jobs, resChannel)
	}
	for _, pos := range positions {
		jobs <- pos
	}
	close(jobs)
	result := 0
	for range n {
		result += <-resChannel
	}
	close(resChannel)
	return result
}

func run(input string) (interface{}, interface{}) {
	baseMap := twod.NewMapFromInput(input)
	width, height := baseMap.Width(), baseMap.Height()
	startingPoint := twod.NewPoint(baseMap.Find('^')[0], twod.UP)
	planPart1, path := runPart1(baseMap.Clone(), startingPoint.Clone(), &width, &height)
	discoveredPlaces := planPart1.Find('X')

	return len(discoveredPlaces), computeWorkerGroup(discoveredPlaces, &baseMap, startingPoint, &path, &width, &height)
}

func main() {
	execute.Run(run, nil, Puzzle, false)
}
