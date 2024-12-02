package main

import (
	"strings"

	"github.com/ysanson/AdventOfCode/pkg"
	"github.com/ysanson/AdventOfCode/pkg/execute"
)

type Conditions struct {
	line   string
	groups []int
}

func parseLine(line string, grow bool) Conditions {
	test, nums, _ := strings.Cut(line, " ")
	if grow {
		var grownLine strings.Builder
		grownLine.WriteString(test)

		var grownNums strings.Builder
		grownNums.WriteString(nums)
		for i := 0; i < 4; i++ {
			grownLine.WriteString("?" + test)
			grownNums.WriteString("," + nums)
		}
		return Conditions{
			line:   grownLine.String(),
			groups: pkg.ParseIntList(grownNums.String(), ","),
		}
	} else {
		return Conditions{
			line:   test,
			groups: pkg.ParseIntList(nums, ","),
		}
	}

}

func findPossibilities(cond Conditions) int {
	poss := 0
	currentStates := map[[4]int]int{{0, 0, 0, 0}: 1}
	newStates := map[[4]int]int{}
	for len(currentStates) > 0 {
		for state, num := range currentStates {
			strIdx, groupIdx, currentGroup, expdot := state[0], state[1], state[2], state[3]
			if strIdx == len(cond.line) {
				if groupIdx == len(cond.groups) {
					poss += num
				}
				continue
			}
			switch {
			case (cond.line[strIdx] == '#' || cond.line[strIdx] == '?') && groupIdx < len(cond.groups) && expdot == 0:
				// we are still looking for broken springs
				if cond.line[strIdx] == '?' && currentGroup == 0 {
					// we are not in a run of broken springs, so ? can be working
					newStates[[4]int{strIdx + 1, groupIdx, currentGroup, expdot}] += num
				}
				currentGroup++
				if currentGroup == cond.groups[groupIdx] {
					// we've found the full next contiguous section of broken springs
					groupIdx++
					currentGroup = 0
					expdot = 1 // we only want a working spring next
				}
				newStates[[4]int{strIdx + 1, groupIdx, currentGroup, expdot}] += num
			case (cond.line[strIdx] == '.' || cond.line[strIdx] == '?') && currentGroup == 0:
				// we are not in a contiguous run of broken springs
				expdot = 0
				newStates[[4]int{strIdx + 1, groupIdx, currentGroup, expdot}] += num
			}
		}
		currentStates, newStates = newStates, currentStates
		for k := range newStates {
			delete(newStates, k)
		}
	}
	return poss
}

func run(input string) (any, any) {
	lines := strings.Split(input, "\n")
	part1, part2 := 0, 0
	for _, line := range lines {
		cond := parseLine(line, false)
		part1 += findPossibilities(cond)
		cond = parseLine(line, true)
		part2 += findPossibilities(cond)
	}

	return part1, part2
}

func main() {
	execute.Run(run, Tests, Puzzle, true)
}
