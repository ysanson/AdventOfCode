package main

import (
	"regexp"
	"strings"

	"github.com/ysanson/AdventOfCode/2023/pkg"
	"github.com/ysanson/AdventOfCode/2023/pkg/execute"
)

func getChar(str string, index int) rune {
	return []rune(str)[index]
}

func sanitizeIndex(index int, maxIndex int) int {
	if index < 0 {
		return 0
	} else if index >= maxIndex {
		return maxIndex - 1
	} else {
		return index
	}
}

func getCharsAroundNumberRec(lineBefore string, currentLine string, lineAfter string, startIndex int, length int, surroundingChars string) string {
	maxLength := len(currentLine)
	if length == 0 {
		// Check right hand
		if lineBefore != "" {
			surroundingChars += string(getChar(lineBefore, sanitizeIndex(startIndex, maxLength)))
		}
		surroundingChars += string(getChar(currentLine, sanitizeIndex(startIndex, maxLength)))
		if lineAfter != "" {
			surroundingChars += string(getChar(lineAfter, sanitizeIndex(startIndex, maxLength)))
		}
		return surroundingChars
	}

	if surroundingChars == "" {
		// Check left hand
		if lineBefore != "" {
			surroundingChars += string(getChar(lineBefore, sanitizeIndex(startIndex-1, maxLength)))
		}
		surroundingChars += string(getChar(currentLine, sanitizeIndex(startIndex-1, maxLength)))
		if lineAfter != "" {
			surroundingChars += string(getChar(lineAfter, sanitizeIndex(startIndex-1, maxLength)))
		}
	}

	// Check up and down the number
	if lineBefore != "" {
		surroundingChars += string(getChar(lineBefore, startIndex))
	}
	if lineAfter != "" {
		surroundingChars += string(getChar(lineAfter, startIndex))
	}

	return getCharsAroundNumberRec(lineBefore, currentLine, lineAfter, startIndex+1, length-1, surroundingChars)
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
