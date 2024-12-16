package main

import (
	"slices"
	"strings"

	"github.com/ysanson/AdventOfCode/pkg/execute"
	"github.com/ysanson/AdventOfCode/pkg/twod"
)

const (
	EMPTY    = '.'
	ROBOT    = '@'
	BOX      = 'O'
	LEFTBOX  = '['
	RIGHTBOX = ']'
	WALL     = '#'
)

func getSpacePosition(position twod.P, plan twod.Map) (twod.Vector, bool) {
	hasBoxes := false
	for {
		position.Move(1)
		if plan[position.Pos] == EMPTY {
			return position.Pos, hasBoxes
		}
		if plan[position.Pos] == WALL {
			return -1, false
		}
		if plan[position.Pos] == BOX || plan[position.Pos] == LEFTBOX || plan[position.Pos] == RIGHTBOX {
			hasBoxes = true
		}
	}
}

func moveBoxes(fromPosition *twod.P, plan *twod.Map) {
	spacePos, hasBoxes := getSpacePosition(*fromPosition, *plan)
	if spacePos == -1 {
		return
	}
	if hasBoxes {
		(*plan)[spacePos] = BOX
	}
	(*plan)[fromPosition.Pos] = EMPTY
	fromPosition.Move(1)
	(*plan)[fromPosition.Pos] = ROBOT
}

func moveWideBoxes(fromPosition *twod.P, plan *twod.Map) {
	direction := fromPosition.Speed
	isVertical := direction.Y() != 0
	boxesMoved := make([]twod.Vector, 0)
	locationsToCheck := []twod.Vector{fromPosition.GetPositionAtDestination(direction, 1)}
	for len(locationsToCheck) != 0 {
		nextLocations := make([]twod.Vector, 0)
		for _, location := range locationsToCheck {
			switch (*plan)[location] {
			case LEFTBOX:
				boxesMoved = append(boxesMoved, location)
				if isVertical {
					nextLocations = append(nextLocations, location+direction)
				}
				nextLocations = append(nextLocations, location+twod.RIGHT+direction)
			case RIGHTBOX:
				boxLeft := location + twod.LEFT
				boxesMoved = append(boxesMoved, boxLeft)
				nextLocations = append(nextLocations, boxLeft+direction)
				if isVertical {
					nextLocations = append(nextLocations, location+direction)
				}
			case WALL:
				return
			default:
				break
			}
		}
		locationsToCheck = nextLocations
	}
	slices.Reverse(boxesMoved)
	for _, boxLocation := range boxesMoved {
		(*plan)[boxLocation] = EMPTY
		(*plan)[boxLocation+twod.RIGHT] = EMPTY
		newLocation := boxLocation + fromPosition.Speed
		(*plan)[newLocation] = LEFTBOX
		(*plan)[newLocation+twod.RIGHT] = RIGHTBOX
	}
	(*plan)[fromPosition.Pos] = EMPTY
	fromPosition.Move(1)
	(*plan)[fromPosition.Pos] = ROBOT
}

func widenMap(originalMap string) twod.Map {
	var sb strings.Builder
	sb.Grow(len(originalMap) * 2)
	for _, tile := range originalMap {
		if tile == BOX {
			sb.WriteString("[]")
		} else if tile == '\n' {
			sb.WriteRune(tile)
		} else if tile == ROBOT {
			sb.WriteString("@.")
		} else {
			sb.WriteRune(tile)
			sb.WriteRune(tile)
		}
	}
	return twod.NewMapFromInput(sb.String())
}

func computeGPS(plan twod.Map, H int) int {
	sum := 0
	for _, box := range plan.Find(BOX, LEFTBOX) {
		sum += box.X() + (H-box.Y()-1)*100
	}
	return sum
}

func part1(plan *twod.Map, moves string) {
	robotPos := twod.P{Pos: plan.Find(ROBOT)[0]}
	for _, move := range moves {
		if move == '\n' {
			continue
		}
		switch move {
		case '^':
			robotPos.Speed = twod.UP
		case '>':
			robotPos.Speed = twod.RIGHT
		case 'v':
			robotPos.Speed = twod.DOWN
		case '<':
			robotPos.Speed = twod.LEFT
		}
		moveBoxes(&robotPos, plan)
	}
}

func part2(input string, moves string) twod.Map {
	plan := widenMap(input)
	robotPos := twod.P{Pos: plan.Find(ROBOT)[0]}
	for _, move := range moves {
		if move == '\n' {
			continue
		}
		switch move {
		case '^':
			robotPos.Speed = twod.UP
		case '>':
			robotPos.Speed = twod.RIGHT
		case 'v':
			robotPos.Speed = twod.DOWN
		case '<':
			robotPos.Speed = twod.LEFT
		}
		moveWideBoxes(&robotPos, &plan)
	}
	return plan
}

func run(input string) (interface{}, interface{}) {
	parsed := strings.Split(input, "\n\n")
	plan := twod.NewMapFromInput(parsed[0])
	H := plan.Height()
	part1(&plan, parsed[1])
	widePlan := part2(parsed[0], parsed[1])

	return computeGPS(plan, H), computeGPS(widePlan, H)
}

func main() {
	execute.Run(run, Tests, Puzzle, false)
}
