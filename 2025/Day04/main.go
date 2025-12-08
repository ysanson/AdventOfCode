package main

import (
	"github.com/ysanson/AdventOfCode/pkg/execute"
	"github.com/ysanson/AdventOfCode/pkg/twod"
)

func verifyPosition(rollsMap twod.Map, position twod.Vector) bool {
	count := 0
	for _, direction := range twod.AllDirections {
		if adjacent := position + direction; rollsMap[adjacent] == '@' {
			count++
		}
	}
	return count < 4
}

func removeRolls(rollsMap twod.Map, rolls []twod.Vector) twod.Map {
	for _, roll := range rolls {
		rollsMap[roll] = '.'
	}
	return rollsMap
}

func findAccessibleRolls(rollsMap twod.Map) (rolls int, removedRolls []twod.Vector) {
	for _, point := range rollsMap.Find('@') {
		if verifyPosition(rollsMap, point) {
			rolls++
			removedRolls = append(removedRolls, point)
		}
	}
	return
}

func run(input string) (any, any) {
	part1, part2, count := 0, 0, 0
	rollsMap := twod.NewMapFromInput(input)
	part1, markedRolls := findAccessibleRolls(rollsMap)
	part2 = part1
	for len(markedRolls) > 0 {
		rollsMap = removeRolls(rollsMap, markedRolls)
		count, markedRolls = findAccessibleRolls(rollsMap)
		part2 += count
	}
	return part1, part2
}

func main() {
	execute.Run(run, Tests, Puzzle, true)
}
