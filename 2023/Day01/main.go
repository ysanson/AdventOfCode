package main

import (
	"regexp"
	"strings"

	"github.com/ysanson/AdventOfCode/pkg"
	"github.com/ysanson/AdventOfCode/pkg/execute"
)

func getGameId(input string) int {
	game := input[strings.IndexRune(input, ' ')+1 : strings.IndexRune(input, ':')]
	return pkg.MustAtoi(game)
}

func replaceSpelledDigits(s string) string {
	re := regexp.MustCompile("(one|two|three|four|five|six|seven|eight|nine)")
	if re.MatchString(s) {
		digitMap := map[string]string{"one": "o1e", "two": "t2o", "three": "t3e", "four": "f4r", "five": "f5e", "six": "s6x", "seven": "s7n", "eight": "e8t", "nine": "n9e"}
		for key, value := range digitMap {
			s = strings.ReplaceAll(s, key, value)
		}
	}
	return s
}

func extractDigits(input string, includeSpelled bool) int {
	re := regexp.MustCompile("[1-9]+")
	formattedInput := input
	if includeSpelled {
		formattedInput = replaceSpelledDigits(input)
	}
	digits := strings.Join(re.FindAllString(formattedInput, -1), "")
	length := len(digits)
	if length == 0 {
		return 0
	} else if length == 1 {
		return pkg.MustAtoi(digits + digits)
	} else {
		return pkg.MustAtoi(digits[0:1] + digits[length-1:])
	}
}

func run(input string) (interface{}, interface{}) {
	lines := strings.Split(input, "\n")
	var part1, part2 = 0, 0
	for _, line := range lines {
		part1 += extractDigits(line, false)
		part2 += extractDigits(line, true)
	}

	return part1, part2
}

func main() {
	execute.Run(run, nil, Puzzle, true)
}
