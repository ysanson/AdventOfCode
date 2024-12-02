package main

import (
	"fmt"
	"image/color"

	"github.com/ysanson/AdventOfCode/pkg/execute"
	"github.com/ysanson/AdventOfCode/pkg/twod"
	"golang.org/x/image/colornames"
)

func run(input string) (any, any) {
	part1, part2 := 0, 0

	m := twod.NewMapFromInput(input)
	var start twod.Vector
	for vector, tile := range m {
		if tile == 'S' {
			start = vector
		}
	}

	distance := make(twod.Map)
	point := twod.NewPoint(start, twod.LEFT)
	startTileCounter := 0
loop:
	for {
		dist, exists := distance[point.Pos]
		if !exists || dist.(int) > point.Steps {
			distance[point.Pos] = point.Steps
		}
		tile, exists := m[point.Pos]
		if !exists {
			panic(fmt.Errorf("%v", point))
		}
		switch tile {
		case '|':
		case '-':
		case '7':
			if point.Speed == twod.RIGHT {
				point.TurnRight()
			} else {
				point.TurnLeft()
			}
		case 'J':
			if point.Speed == twod.DOWN {
				point.TurnRight()
			} else {
				point.TurnLeft()
			}
		case 'F':
			if point.Speed == twod.UP {
				point.TurnRight()
			} else {
				point.TurnLeft()
			}
		case 'L':
			if point.Speed == twod.LEFT {
				point.TurnRight()
			} else {
				point.TurnLeft()
			}
		case '.':
			panic(fmt.Errorf("%v", point))
		case 'S':
			if startTileCounter == 1 {
				point = twod.NewPoint(start, twod.DOWN)
			} else if startTileCounter == 2 {
				break loop
			}
			startTileCounter++
		}
		point.Move(1)
	}

	for v := range m {
		dist, exist := distance[v]
		if exist && dist.(int) > part1 {
			part1 = dist.(int)
		}
	}

	insidePoints := make(twod.Map)
	for pos := range m {
		if _, exist := distance[pos]; exist {
			continue
		}
		crosses := 0

		for y := pos.Y(); y >= 0; y-- {
			newPos := twod.NewVector(pos.X(), y)
			_, exist := distance[newPos]

			if exist && (m[newPos] == '-' || m[newPos] == 'L' || m[newPos] == 'F') {
				crosses++
			}
		}
		if crosses%2 == 1 {
			insidePoints[pos] = '.'
			part2++
		}
	}

	twod.RenderingMap = map[interface{}]color.Color{
		'.': colornames.Blue,
	}
	for i := 0; i <= part1; i++ {
		twod.RenderingMap[i] = colornames.Green
	}

	for vector, i := range insidePoints {
		distance[vector] = i
	}

	distance.Render()
	//time.Sleep(time.Minute)

	return part1, part2
}

func main() {
	execute.RunWithPixel(run, nil, Puzzle, true)
}
