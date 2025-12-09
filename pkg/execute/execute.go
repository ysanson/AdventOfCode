package execute

import (
	"fmt"
	"os"
	"time"

	"github.com/faiface/pixel/pixelgl"
	"github.com/ysanson/AdventOfCode/pkg/twod"
)

func Run(run func(string) (any, any), test TestCases, puzzle string, verbose bool) {
	if test != nil {
		test.Run(run, !verbose)
	}
	if puzzle != "" {
		start := time.Now()
		part1, part2 := run(puzzle)
		elapsed := time.Since(start)

		fmt.Printf("Part 1: %v\nPart 2: %v\n", part1, part2)
		fmt.Printf("Execution took %s", elapsed)
	}

}

func RunWithPixel(run func(string) (interface{}, interface{}), tests TestCases, puzzle string, verbose bool) {
	if os.Getenv("CI") == "true" {
		Run(run, tests, puzzle, verbose)
		return
	}
	twod.RenderingEnabled = true
	pixelgl.Run(func() {
		Run(run, tests, puzzle, verbose)
	})
}
