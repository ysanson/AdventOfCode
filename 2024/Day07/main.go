package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"

	"github.com/ysanson/AdventOfCode/pkg/execute"
)

type WorkerData struct {
	numbers []int64
	sum     int64
	result  int64
}

func parseEquation(line string) (int64, []int64) {
	line = strings.Replace(line, ":", "", 1)
	operators := strings.Split(line, " ")
	numbers := make([]int64, len(operators))
	for i, operator := range operators {
		if n, err := strconv.ParseInt(operator, 10, 64); err == nil {
			numbers[i] = n
		}

	}
	return numbers[0], numbers[1:]
}

func processPart1(numbers []int64, sum, result int64) bool {
	if len(numbers) == 0 {
		return result == sum
	}
	if result > sum {
		return false
	}
	return processPart1(numbers[1:], sum, result+numbers[0]) || processPart1(numbers[1:], sum, result*numbers[0])
}

func processPart2(numbers []int64, sum, result int64) bool {
	if len(numbers) == 0 {
		return result == sum
	}
	if result > sum {
		return false
	}
	concat, err := strconv.ParseInt(fmt.Sprintf("%d%d", result, numbers[0]), 10, 64)
	if err != nil {
		panic(err)
	}
	return processPart2(numbers[1:], sum, result+numbers[0]) || processPart2(numbers[1:], sum, result*numbers[0]) || processPart2(numbers[1:], sum, concat)
}

func workerP1(jobs <-chan WorkerData, result chan<- int64) {
	for job := range jobs {
		if processPart1(job.numbers, job.sum, job.result) {
			result <- job.sum
		} else {
			result <- 0
		}
	}
}

func workerP2(jobs <-chan WorkerData, result chan<- int64) {
	for job := range jobs {
		if processPart2(job.numbers, job.sum, job.result) {
			result <- job.sum
		} else {
			result <- 0
		}
	}
}

func run(input string) (interface{}, interface{}) {
	lines := strings.Split(input, "\n")
	n := len(lines)
	jobsP1 := make(chan WorkerData, n)
	resChannelP1 := make(chan int64, n)
	jobsP2 := make(chan WorkerData, n)
	resChannelP2 := make(chan int64, n)
	for range runtime.NumCPU() / 2 {
		go workerP1(jobsP1, resChannelP1)
		go workerP2(jobsP2, resChannelP2)
	}
	for _, line := range lines {
		sum, operands := parseEquation(line)
		jobsP1 <- WorkerData{numbers: operands[1:], sum: sum, result: operands[0]}
		jobsP2 <- WorkerData{numbers: operands[1:], sum: sum, result: operands[0]}
	}
	close(jobsP1)
	close(jobsP2)
	var part1, part2 int64
	for range n {
		part1 += <-resChannelP1
		part2 += <-resChannelP2
	}
	return part1, part2
}

func main() {
	execute.Run(run, Tests, Puzzle, false)
}
