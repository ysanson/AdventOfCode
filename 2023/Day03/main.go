package main

import (
	"regexp"
	"strings"

	"github.com/ysanson/AdventOfCode/2023/pkg"
	"github.com/ysanson/AdventOfCode/2023/pkg/execute"
)

func getCharsAroundNumberRec(lineBefore string, currentLine string, lineAfter string, index int, length int, surroundingChars string) string {
	maxLength := len(currentLine)
	if length == 0 {
		// Check right hand
		if lineBefore != "" {
			surroundingChars += string(pkg.GetChar(lineBefore, pkg.SanitizeIndex(index, maxLength)))
		}
		surroundingChars += string(pkg.GetChar(currentLine, pkg.SanitizeIndex(index, maxLength)))
		if lineAfter != "" {
			surroundingChars += string(pkg.GetChar(lineAfter, pkg.SanitizeIndex(index, maxLength)))
		}
		return surroundingChars
	}

	if surroundingChars == "" {
		// Check left hand
		if lineBefore != "" {
			surroundingChars += string(pkg.GetChar(lineBefore, pkg.SanitizeIndex(index-1, maxLength)))
		}
		surroundingChars += string(pkg.GetChar(currentLine, pkg.SanitizeIndex(index-1, maxLength)))
		if lineAfter != "" {
			surroundingChars += string(pkg.GetChar(lineAfter, pkg.SanitizeIndex(index-1, maxLength)))
		}
	}

	// Check up and down the number
	if lineBefore != "" {
		surroundingChars += string(pkg.GetChar(lineBefore, index))
	}
	if lineAfter != "" {
		surroundingChars += string(pkg.GetChar(lineAfter, index))
	}

	return getCharsAroundNumberRec(lineBefore, currentLine, lineAfter, index+1, length-1, surroundingChars)
}

func analyzeLine(lineBefore string, currentLine string, lineAfter string) int {
	result := 0
	re := regexp.MustCompile("\\d+")
	specialCharsRe := regexp.MustCompile("[^\\d.]")

	for _, numbers := range re.FindAllStringIndex(currentLine, -1) {
		charsAround := getCharsAroundNumberRec(lineBefore, currentLine, lineAfter, numbers[0], numbers[1]-numbers[0], "")
		if specialCharsRe.MatchString(charsAround) {
			result += pkg.MustAtoi(currentLine[numbers[0]:numbers[1]])
		}
	}

	return result
}

func run(input string) (interface{}, interface{}) {
	lines := strings.Split(input, "\n")

	part1 := analyzeLine("", lines[0], lines[1]) + analyzeLine(lines[len(lines)-2], lines[len(lines)-1], "")
	for i := 1; i < len(lines)-1; i++ {
		part1 += analyzeLine(lines[i-1], lines[i], lines[i+1])
	}
	return part1, 0
}

func main() {
	execute.Run(run, Tests, Puzzle, true)
}
