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
		switch plan[position.Pos] {
		case EMPTY:
			return position.Pos, hasBoxes
		case WALL:
			return -1, false
		case BOX:
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

func processMoves(plan *twod.Map, moves string, moveFunc func(*twod.P, *twod.Map)) {
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
		moveFunc(&robotPos, plan)
	}
}

func run(input string) (interface{}, interface{}) {
	parsed := strings.Split(input, "\n\n")
	plan := twod.NewMapFromInput(parsed[0])
	H := plan.Height()
	processMoves(&plan, parsed[1], moveBoxes)

	widePlan := widenMap(parsed[0])
	processMoves(&widePlan, parsed[1], moveWideBoxes)

	return computeGPS(plan, H), computeGPS(widePlan, H)
}

func main() {
	execute.Run(run, Tests, Puzzle, false)
}
