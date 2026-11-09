package execute

import (
	"fmt"
	"time"
)

type result struct {
	part1, part2 any
	elapsed      time.Duration
}

func Run(run func(string) (any, any), test TestCases, puzzle string, verbose bool) {
	// Inputs are only solved once, so a puzzle that is also listed in the test
	// table is not computed a second time below.
	results := make(map[string]result)
	solve := func(input string) (any, any) {
		r, seen := results[input]
		if !seen {
			start := time.Now()
			r.part1, r.part2 = run(input)
			r.elapsed = time.Since(start)
			results[input] = r
		}
		return r.part1, r.part2
	}

	if test != nil {
		test.Run(solve, !verbose)
	}
	if puzzle != "" {
		part1, part2 := solve(puzzle)
		fmt.Printf("Part 1: %v\nPart 2: %v\n", part1, part2)
		fmt.Printf("Execution took %s", results[puzzle].elapsed)
	}
}
