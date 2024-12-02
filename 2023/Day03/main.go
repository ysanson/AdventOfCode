package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ysanson/AdventOfCode/pkg"
	"github.com/ysanson/AdventOfCode/pkg/execute"
)

type Point struct {
	row int
	col int
}

var defaultPoint = Point{row: -1, col: -1}

func appendChar(line string, index int, lineNumber int, buildString string) (string, Point) {
	maxLength := len(line)
	idx := pkg.SanitizeIndex(index, maxLength)
	char := pkg.GetChar(line, idx)
	if char == '*' {
		return buildString + string(char), Point{row: lineNumber, col: index}
	} else {
		return buildString + string(char), defaultPoint
	}
}

func getCharsAroundNumberRec(lineBefore string, currentLine string, lineAfter string, index int, length int, lineNumber int, surroundingChars string, gear Point) (string, Point) {
	var newGear Point
	if length == 0 {
		// Check right hand
		if lineBefore != "" {
			surroundingChars, newGear = appendChar(lineBefore, index, lineNumber-1, surroundingChars)
			if newGear.col != -1 && newGear.row != -1 {
				gear = newGear
			}
		}
		surroundingChars, newGear = appendChar(currentLine, index, lineNumber, surroundingChars)
		if newGear.col != -1 && newGear.row != -1 {
			gear = newGear
		}
		if lineAfter != "" {
			surroundingChars, newGear = appendChar(lineAfter, index, lineNumber+1, surroundingChars)
			if newGear.col != -1 && newGear.row != -1 {
				gear = newGear
			}
		}
		return surroundingChars, gear
	}

	if surroundingChars == "" {
		// Check left hand
		if lineBefore != "" {
			surroundingChars, newGear = appendChar(lineBefore, index-1, lineNumber-1, surroundingChars)
			if newGear.col != -1 && newGear.row != -1 {
				gear = newGear
			}
		}
		surroundingChars, newGear = appendChar(currentLine, index-1, lineNumber, surroundingChars)
		if newGear.col != -1 && newGear.row != -1 {
			gear = newGear
		}
		if lineAfter != "" {
			surroundingChars, newGear = appendChar(lineAfter, index-1, lineNumber+1, surroundingChars)
			if newGear.col != -1 && newGear.row != -1 {
				gear = newGear
			}
		}
	}

	// Check up and down the number
	if lineBefore != "" {
		surroundingChars, newGear = appendChar(lineBefore, index, lineNumber-1, surroundingChars)
		if newGear.col != -1 && newGear.row != -1 {
			gear = newGear
		}
	}
	if lineAfter != "" {
		surroundingChars, newGear = appendChar(lineAfter, index, lineNumber+1, surroundingChars)
		if newGear.col != -1 && newGear.row != -1 {
			gear = newGear
		}
	}

	return getCharsAroundNumberRec(lineBefore, currentLine, lineAfter, index+1, length-1, lineNumber, surroundingChars, gear)
}

func analyzeLine(lineBefore string, currentLine string, lineAfter string, lineNumber int, gears map[string][]int) int {
	result := 0
	re := regexp.MustCompile("\\d+")
	specialCharsRe := regexp.MustCompile("[^\\d.]")
	var charsAround string
	var gear Point
	var extractedNumber int
	for _, numbers := range re.FindAllStringIndex(currentLine, -1) {
		charsAround, gear = getCharsAroundNumberRec(lineBefore, currentLine, lineAfter, numbers[0], numbers[1]-numbers[0], lineNumber, "", defaultPoint)
		if specialCharsRe.MatchString(charsAround) {
			extractedNumber = pkg.MustAtoi(currentLine[numbers[0]:numbers[1]])
			result += extractedNumber
			if gear.col != -1 && gear.row != -1 {
				key := fmt.Sprint(gear.row, gear.col)
				gears[key] = append(gears[key], extractedNumber)
			}
		}
	}

	return result
}

func run(input string) (interface{}, interface{}) {
	lines := strings.Split(input, "\n")
	var gears = map[string][]int{}

	part1 := analyzeLine("", lines[0], lines[1], 0, gears) + analyzeLine(lines[len(lines)-2], lines[len(lines)-1], "", len(lines)-1, gears)
	for i := 1; i < len(lines)-1; i++ {
		part1 += analyzeLine(lines[i-1], lines[i], lines[i+1], i, gears)
	}

	part2 := 0
	for _, numbers := range gears {
		if len(numbers) == 2 {
			part2 += numbers[0] * numbers[1]
		}
	}

	return part1, part2
}

func main() {
	execute.Run(run, Tests, Puzzle, true)
}
