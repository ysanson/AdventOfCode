package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
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

func runPart1(plan twod.Map, startingPos *twod.P) (twod.Map, []VisitedPoint) {
	position := startingPos
	visited := make([]VisitedPoint, 0)
	for !position.IsOutOfMap(plan) {
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

// if same point and same direction => cycle detected
func hasCycle(visited map[VisitedPoint]bool, current VisitedPoint) bool {
	_, ok := visited[current]
	return ok
}

func runWithObstacle(plan *twod.Map, position *twod.P, upd twod.Vector, path *[]VisitedPoint) int {
	visited := make(map[VisitedPoint]bool)
	var current VisitedPoint
	if skip := slices.IndexFunc(*path, func(e VisitedPoint) bool { return e.Pos.X() == upd.X() && e.Pos.Y() == upd.Y() }); skip > 0 {
		current = (*path)[skip-1]
	}
	for !position.IsOutOfMap(*plan) {
		current = VisitedPoint{Pos: position.Pos, Direction: position.Speed}
		if hasCycle(visited, current) {
			fmt.Printf("Found cycle at X=%d, Y=%d\n", upd.X(), upd.Y())
			return 1
		}
		visited[current] = true
		dest := position.GetPositionAtDestination(position.Speed, 1)
		for (*plan)[dest] == '#' || dest == upd {
			position.Speed = getNextDirection(position.Speed)
			dest = position.GetPositionAtDestination(position.Speed, 1)
		}
		position.Move(1)
	}
	return 0
}

func worker(baseMap *twod.Map, startingPoint *twod.P, path *[]VisitedPoint, jobs <-chan twod.Vector, results chan<- int) {
	for position := range jobs {
		results <- runWithObstacle(baseMap, startingPoint.Clone(), position, path)
	}
}

func computeWorkerGroup(positions []twod.Vector, baseMap *twod.Map, startingPoint *twod.P, path *[]VisitedPoint) int {
	n := len(positions)
	jobs := make(chan twod.Vector, n)
	resChannel := make(chan int, n)
	for range runtime.NumCPU() {
		go worker(baseMap, startingPoint, path, jobs, resChannel)
	}
	for _, pos := range positions {
		jobs <- pos
	}
	close(jobs)
	result := 0
	for a := 1; a <= n; a++ {
		result += <-resChannel
	}
	return result
}

func run(input string) (interface{}, interface{}) {
	baseMap := twod.NewMapFromInput(input)
	startingPoint := twod.NewPoint(baseMap.Find('^')[0], twod.UP)
	planPart1, path := runPart1(baseMap.Clone(), startingPoint.Clone())
	fmt.Println("Part 1 is done")
	discoveredPlaces := planPart1.Find('X')

	return len(discoveredPlaces), computeWorkerGroup(discoveredPlaces, &baseMap, startingPoint, &path)
}

func main() {
	f, _ := os.Create("trace.out")
	trace.Start(f)
	defer trace.Stop()
	execute.Run(run, nil, Puzzle, false)
}
