package main

import "github.com/ysanson/AdventOfCode/pkg/execute"

var Tests = execute.TestCases{
	{
		Input:         Test1,
		ExpectedPart1: 288,
		ExpectedPart2: 71503,
	},
	// {
	// 	Input:         Puzzle,
	// 	ExpectedPart1: 993500720,
	// 	ExpectedPart2: 6050769,
	// },
}

const Test1 = `Time:      7  15   30
Distance:  9  40  200`

const Puzzle = `Time:        44     70     70     80
Distance:   283   1134   1134   1491`
