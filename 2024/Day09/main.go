package main

import (
	"fmt"
	"slices"

	"github.com/ysanson/AdventOfCode/pkg/execute"
)

type Byte struct {
	from   int
	length int
	value  int
}

func createArray(input []int) []int {
	length := 0
	for _, l := range input {
		length += l
	}
	diskMap := make([]int, 0, length)
	for idx, l := range input {
		for range l {
			if idx%2 != 0 {
				// If idx is pair => empty space
				diskMap = append(diskMap, -1)
			} else {
				diskMap = append(diskMap, idx/2)
			}
		}
	}
	return diskMap
}

func getLastFileIdx(diskMap []int) (int, int) {
	for i := len(diskMap) - 1; i >= 0; i-- {
		if diskMap[i] != -1 {
			firstIdx := slices.Index(diskMap, diskMap[i])
			return firstIdx, i - firstIdx + 1
		}
	}
	fmt.Println("Could not find last file")
	return 0, 0
}

func processFileP1(diskMap []int) []int {
	lastFileIdx, fileLength := getLastFileIdx(diskMap)
	for range fileLength {
		emptySpaceIdx := slices.Index(diskMap, -1)
		if emptySpaceIdx == -1 {
			panic("No more empty space!")
		}
		diskMap[emptySpaceIdx] = diskMap[lastFileIdx]
		diskMap[lastFileIdx] = -1
		lastFileIdx++
	}
	return diskMap
}

func processWholeFile(diskMap []int, upperLimit int) ([]int, int) {
	lastFileIdx, fileLength := getLastFileIdx(diskMap[:upperLimit])
	if start := hasSpaceForFile(diskMap[:upperLimit], fileLength, lastFileIdx+1); start != -1 {
		for i := range fileLength {
			diskMap[start+i] = diskMap[lastFileIdx]
			diskMap[lastFileIdx] = -1
			lastFileIdx++
		}
	}
	return diskMap, lastFileIdx
}

func hasSpaceBetweenFiles(diskMap []int) bool {
	emptySpaceIdx := slices.Index(diskMap, -1)
	if emptySpaceIdx == -1 {
		panic("No more empty space!")
	}
	return slices.ContainsFunc(diskMap[emptySpaceIdx:], func(chunk int) bool { return chunk != -1 })
}

func hasSpaceForFile(diskMap []int, fileLength, until int) int {
	searchFrom := 0
	startIdx := slices.Index(diskMap[searchFrom:until], -1) + searchFrom
	for searchFrom <= until-1 && startIdx-searchFrom != -1 {
		endIdx := slices.IndexFunc(diskMap[startIdx:until], func(e int) bool { return e != -1 }) + startIdx
		if endIdx-startIdx >= fileLength {
			return startIdx
		} else if startIdx >= endIdx {
			return -1
		}
		searchFrom = endIdx
		startIdx = slices.Index(diskMap[searchFrom:until], -1) + searchFrom
	}
	return -1
}

func computeChecksum(diskMap []int) int {
	checksum := 0
	for pos, chunk := range diskMap {
		if chunk != -1 {
			checksum += pos * chunk
		}
	}
	return checksum
}

func run(input string) (interface{}, interface{}) {
	in := make([]int, len(input))
	for i, char := range input {
		in[i] = int(char - '0')
	}
	fragmentedMap := createArray(in)
	continuousMap := slices.Clone(fragmentedMap)
	for hasSpaceBetweenFiles(fragmentedMap) {
		fragmentedMap = processFileP1(fragmentedMap)
	}

	upperLimit := len(continuousMap)
	for upperLimit > 0 {
		if hasSpaceBetweenFiles(continuousMap[:upperLimit]) {
			continuousMap, upperLimit = processWholeFile(continuousMap, upperLimit)
		} else {
			upperLimit = 0
		}

	}

	return computeChecksum(fragmentedMap), computeChecksum(continuousMap)
}

func main() {
	execute.Run(run, Tests, Puzzle, false)
}
