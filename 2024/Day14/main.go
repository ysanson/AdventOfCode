package main

import (
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/ysanson/AdventOfCode/pkg"
	"github.com/ysanson/AdventOfCode/pkg/execute"
)

type Robot struct {
	x    int
	y    int
	velX int
	velY int
}

func (r *Robot) UpdateCoords(x, y int) {
	r.x = x
	r.y = y
}

func isUniquePos(robots []Robot) bool {
	uniquePositions := make(map[Robot]bool)
	for _, robot := range robots {
		uniquePositions[Robot{x: robot.x, y: robot.y}] = true
	}
	return len(uniquePositions) == len(robots)
}

// hasRow reports whether at least length robots stand side by side in a row.
func hasRow(robots []Robot, length int) bool {
	occupied := make(map[[2]int]bool, len(robots))
	for _, robot := range robots {
		occupied[[2]int{robot.x, robot.y}] = true
	}
	for pos := range occupied {
		run := 0
		for occupied[[2]int{pos[0] + run, pos[1]}] {
			run++
		}
		if run >= length {
			return true
		}
	}
	return false
}

func printGrid(robots []Robot, width, height int) {
	var sb strings.Builder
	for y := range height {
		for x := range width {
			if slices.ContainsFunc(robots, func(r Robot) bool { return r.x == x && r.y == y }) {
				sb.WriteRune('#')
			} else {
				sb.WriteRune('.')
			}
		}
		sb.WriteRune('\n')
	}
	fmt.Println(sb.String())
}

func run(input string) (interface{}, interface{}) {
	lines := strings.Split(input, "\n")
	re := regexp.MustCompile(`[-]?\d+`)
	robots := make([]Robot, len(lines))
	for i, line := range lines {
		nums := re.FindAllString(line, -1)
		robots[i] = Robot{x: pkg.MustAtoi(nums[0]), y: pkg.MustAtoi(nums[1]), velX: pkg.MustAtoi(nums[2]), velY: pkg.MustAtoi(nums[3])}
	}
	maxWidth, maxHeight := 101, 103
	if len(robots) == 12 {
		maxWidth, maxHeight = 11, 7
	}

	for i, robot := range robots {
		newX := (robot.x + robot.velX*100) % maxWidth
		newY := (robot.y + robot.velY*100) % maxHeight
		for newX < 0 {
			newX += maxWidth
		}
		for newY < 0 {
			newY += maxHeight
		}
		robots[i].UpdateCoords(pkg.Abs(newX), pkg.Abs(newY))
	}
	midWidth, midHeight := maxWidth/2, maxHeight/2

	quad1, quad2, quad3, quad4 := 0, 0, 0, 0
	for _, robot := range robots {
		if robot.x < midWidth && robot.y < midHeight {
			quad1++
		} else if robot.x < midWidth && robot.y > midHeight {
			quad2++
		} else if robot.x > midWidth && robot.y < midHeight {
			quad3++
		} else if robot.x > midWidth && robot.y > midHeight {
			quad4++
		}
	}

	part1 := pkg.Multiply(quad1, quad2, quad3, quad4)
	if len(robots) == 12 {
		return part1, 0 // the example has no tree to find
	}

	// Positions repeat after width*height seconds.
	for count := 101; count <= 100+maxWidth*maxHeight; count++ {
		for i, robot := range robots {
			newX := (robot.x + robot.velX) % maxWidth
			newY := (robot.y + robot.velY) % maxHeight
			for newX < 0 {
				newX += maxWidth
			}
			for newY < 0 {
				newY += maxHeight
			}
			robots[i].UpdateCoords(pkg.Abs(newX), pkg.Abs(newY))
		}
		if isUniquePos(robots) && hasRow(robots, 20) {
			return part1, count
		}
	}

	return part1, 0
}

func main() {
	execute.Run(run, nil, Puzzle, false)
}
