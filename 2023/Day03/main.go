package main

import (
	"regexp"
	"strings"
	"unicode"

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
		return maxIndex
	} else {
		return index
	}
}

func getCharsAroundNumber(lineBefore string, currentLine string, lineAfter string, startIndex int) string {
	surroundingChars := ""
	index := startIndex
	maxLength := len(currentLine)

	// First digit check
	if lineBefore != "" {
		surroundingChars += lineBefore[sanitizeIndex(index-1, maxLength) : index+1]
	}
	surroundingChars += string(getChar(currentLine, sanitizeIndex(index-1, maxLength)))
	if lineAfter != "" {
		surroundingChars += lineAfter[sanitizeIndex(index-1, maxLength) : index+1]
	}

	// Digits between check
	for index < maxLength && unicode.IsDigit(getChar(currentLine, index)) {
		if lineBefore != "" {
			surroundingChars += string(getChar(lineBefore, index))
		}
		if lineAfter != "" {
			surroundingChars += string(getChar(lineAfter, index))
		}
		index++
	}

	if index < maxLength {
		// Last digit check
		if lineBefore != "" {
			surroundingChars += lineBefore[index-1 : sanitizeIndex(index+1, maxLength)]
		}
		surroundingChars += string(getChar(currentLine, sanitizeIndex(index, maxLength)))
		if lineAfter != "" {
			surroundingChars += lineAfter[index-1 : sanitizeIndex(index+1, maxLength)]
		}
	}

	return surroundingChars
}

func analyzeLine(lineBefore string, currentLine string, lineAfter string) int {
	result := 0
	re := regexp.MustCompile("[1-9]+")
	specialCharsRe := regexp.MustCompile("[^0-9.]")

	for _, numbers := range re.FindAllStringIndex(currentLine, -1) {
		charsAround := getCharsAroundNumber(lineBefore, currentLine, lineAfter, numbers[0])
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
