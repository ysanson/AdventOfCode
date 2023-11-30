package execute

import (
	"fmt"
	"time"
)

func Run(run func(string) (interface{}, interface{}), test TestCases, puzzle string, verbose bool) {
	if test != nil {
		test.Run(run, !verbose)
	}

	start := time.Now()
	part1, part2 := run(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %v\nPart 2: %v", part1, part2)
	fmt.Printf("Execution took %s", elapsed)
}
