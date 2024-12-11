package main

import (
	"runtime"

	"github.com/ysanson/AdventOfCode/pkg/execute"
	"github.com/ysanson/AdventOfCode/pkg/twod"
)

func findTrail(current twod.P, topoMap twod.Map, foundSummits map[twod.Vector]bool) int {
	curHeight := topoMap[current.Pos]
	up, down := current.GetPositionAtDestination(twod.UP, 1), current.GetPositionAtDestination(twod.DOWN, 1)
	left, right := current.GetPositionAtDestination(twod.LEFT, 1), current.GetPositionAtDestination(twod.RIGHT, 1)
	summit := 0
	if curHeight == '8' {
		if _, ok := foundSummits[up]; topoMap[up] == '9' && !ok {
			foundSummits[up] = true
			summit++
		}
		if _, ok := foundSummits[down]; topoMap[down] == '9' && !ok {
			foundSummits[down] = true
			summit++
		}
		if _, ok := foundSummits[left]; topoMap[left] == '9' && !ok {
			foundSummits[left] = true
			summit++
		}
		if _, ok := foundSummits[right]; topoMap[right] == '9' && !ok {
			foundSummits[right] = true
			summit++
		}
	} else {
		if pos, ok := topoMap[up]; ok && pos.(int32)-curHeight.(int32) == 1 {
			summit += findTrail(twod.P{Pos: up, Speed: twod.DOWN}, topoMap, foundSummits)
		}
		if pos, ok := topoMap[down]; ok && pos.(int32)-curHeight.(int32) == 1 {
			summit += findTrail(twod.P{Pos: down, Speed: twod.DOWN}, topoMap, foundSummits)
		}
		if pos, ok := topoMap[left]; ok && pos.(int32)-curHeight.(int32) == 1 {
			summit += findTrail(twod.P{Pos: left, Speed: twod.DOWN}, topoMap, foundSummits)
		}
		if pos, ok := topoMap[right]; ok && pos.(int32)-curHeight.(int32) == 1 {
			summit += findTrail(twod.P{Pos: right, Speed: twod.DOWN}, topoMap, foundSummits)
		}
	}
	return summit
}

func findUniqueTrails(current twod.P, topoMap twod.Map) int {
	curHeight := topoMap[current.Pos]
	up, down := current.GetPositionAtDestination(twod.UP, 1), current.GetPositionAtDestination(twod.DOWN, 1)
	left, right := current.GetPositionAtDestination(twod.LEFT, 1), current.GetPositionAtDestination(twod.RIGHT, 1)
	summit := 0
	if curHeight == '8' {
		if topoMap[up] == '9' {
			summit++
		}
		if topoMap[down] == '9' {
			summit++
		}
		if topoMap[left] == '9' {
			summit++
		}
		if topoMap[right] == '9' {
			summit++
		}
	} else {
		if pos, ok := topoMap[up]; ok && pos.(int32)-curHeight.(int32) == 1 {
			summit += findUniqueTrails(twod.P{Pos: up, Speed: twod.DOWN}, topoMap)
		}
		if pos, ok := topoMap[down]; ok && pos.(int32)-curHeight.(int32) == 1 {
			summit += findUniqueTrails(twod.P{Pos: down, Speed: twod.DOWN}, topoMap)
		}
		if pos, ok := topoMap[left]; ok && pos.(int32)-curHeight.(int32) == 1 {
			summit += findUniqueTrails(twod.P{Pos: left, Speed: twod.DOWN}, topoMap)
		}
		if pos, ok := topoMap[right]; ok && pos.(int32)-curHeight.(int32) == 1 {
			summit += findUniqueTrails(twod.P{Pos: right, Speed: twod.DOWN}, topoMap)
		}
	}
	return summit
}
func trailWorker(topoMap twod.Map, jobs <-chan twod.P, result chan<- int) {
	for pos := range jobs {
		result <- findTrail(pos, topoMap, make(map[twod.Vector]bool))
	}
}

func uniqueTrailWorker(topoMap twod.Map, jobs <-chan twod.P, result chan<- int) {
	for pos := range jobs {
		result <- findUniqueTrails(pos, topoMap)
	}
}

func run(input string) (interface{}, interface{}) {
	topoMap := twod.NewMapFromInput(input)
	trailheads := topoMap.Find('0')
	jobs := make(chan twod.P, len(trailheads))
	results := make(chan int, len(trailheads))
	jobsP2 := make(chan twod.P, len(trailheads))
	resultsP2 := make(chan int, len(trailheads))
	for range runtime.NumCPU() / 2 {
		go trailWorker(topoMap, jobs, results)
		go uniqueTrailWorker(topoMap, jobsP2, resultsP2)
	}
	for _, trailhead := range trailheads {
		jobs <- twod.P{Pos: trailhead, Speed: twod.DOWN}
		jobsP2 <- twod.P{Pos: trailhead, Speed: twod.DOWN}
	}
	close(jobs)
	close(jobsP2)

	part1, part2 := 0, 0
	for range len(trailheads) {
		part1 += <-results
		part2 += <-resultsP2
	}
	return part1, part2
}

func main() {
	execute.Run(run, Tests, Puzzle, false)
}
